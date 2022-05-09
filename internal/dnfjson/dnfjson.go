// Package dnfjson is an interface to the dnf-json Python script that is
// packaged with the osbuild-composer project. The core component of this
// package is the Solver type. The Solver can be configured with
// distribution-specific values (platform ID, architecture, and version
// information) and provides methods for dependency resolution (Depsolve) and
// retrieving a full list of repository package metadata (FetchMetadata).
//
// Alternatively, a BaseSolver can be created which represents an un-configured
// Solver. This type can't be used for depsolving, but can be used to create
// configured Solver instances sharing the same cache directory and
// subscription credentials.
//
// This package relies on the types defined in rpmmd to describe RPM package
// metadata.
package dnfjson

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"sort"

	"github.com/osbuild/osbuild-composer/internal/rhsm"
	"github.com/osbuild/osbuild-composer/internal/rpmmd"
)

type BaseSolver struct {
	subscriptions *rhsm.Subscriptions

	// Cache directory for the DNF metadata
	cacheDir string

	// Path to the dnf-json binary and optional args (default: "/usr/libexec/osbuild-composer/dnf-json")
	dnfJsonCmd []string
}

// Create a new unconfigured BaseSolver (without platform information). It can
// be used to create configured Solver instances with the NewWithConfig()
// method. Creating a BaseSolver also loads system subscription information.
func NewBaseSolver(cacheDir string) *BaseSolver {
	subscriptions, _ := rhsm.LoadSystemSubscriptions()
	return &BaseSolver{
		cacheDir:      cacheDir,
		subscriptions: subscriptions,
		dnfJsonCmd:    []string{"/usr/libexec/osbuild-composer/dnf-json"},
	}
}

// SetDNFJSONPath sets the path to the dnf-json binary and optionally any command line arguments.
func (s *BaseSolver) SetDNFJSONPath(cmd string, args ...string) {
	s.dnfJsonCmd = append([]string{cmd}, args...)
}

// NewWithConfig initialises a Solver with the platform information and the
// BaseSolver's subscription info, cache directory, and dnf-json path.
func (bs *BaseSolver) NewWithConfig(modulePlatformID string, releaseVer string, arch string) *Solver {
	s := new(Solver)
	s.BaseSolver = *bs
	s.modulePlatformID = modulePlatformID
	s.arch = arch
	s.releaseVer = releaseVer
	return s
}

// Solver is configured with system information in order to resolve
// dependencies for RPM packages using DNF.
type Solver struct {
	BaseSolver

	// Platform ID, e.g., "platform:el8"
	modulePlatformID string

	// System architecture
	arch string

	// Release version of the distro. This is used in repo files on the host
	// system and required for subscription support.
	releaseVer string
}

// Create a new Solver with the given configuration. Initialising a Solver also loads system subscription information.
func NewSolver(modulePlatformID string, releaseVer string, arch string, cacheDir string) *Solver {
	s := NewBaseSolver(cacheDir)
	return s.NewWithConfig(modulePlatformID, releaseVer, arch)
}

// Depsolve the given packages with explicit excludes using the given configuration and repos
func Depsolve(pkgSets []rpmmd.PackageSet, modulePlatformID string, releaseVer string, arch string, cacheDir string) (*DepsolveResult, error) {
	return NewSolver(modulePlatformID, releaseVer, arch, cacheDir).Depsolve(pkgSets)
}

// Depsolve the list of required package sets with explicit excludes using
// their associated repositories.  Each package set is depsolved as a separate
// transactions in a chain.  It returns a list of all packages (with solved
// dependencies) that will be installed into the system.
func (s *Solver) Depsolve(pkgSets []rpmmd.PackageSet) (*DepsolveResult, error) {
	req, repoMap, err := s.makeDepsolveRequest(pkgSets)
	if err != nil {
		return nil, err
	}

	output, err := run(s.dnfJsonCmd, req)
	if err != nil {
		return nil, err
	}
	var result *depsolveResult
	if err := json.Unmarshal(output, &result); err != nil {
		return nil, err
	}

	return resultToPublic(result, repoMap), nil
}

func FetchMetadata(repos []rpmmd.RepoConfig, modulePlatformID string, releaseVer string, arch string, cacheDir string) (*FetchMetadataResult, error) {
	return NewSolver(modulePlatformID, releaseVer, arch, cacheDir).FetchMetadata(repos)
}

func (s *Solver) FetchMetadata(repos []rpmmd.RepoConfig) (*FetchMetadataResult, error) {
	req, err := s.makeDumpRequest(repos)
	if err != nil {
		return nil, err
	}
	result, err := run(s.dnfJsonCmd, req)
	if err != nil {
		return nil, err
	}

	metadata := new(FetchMetadataResult)
	if err := json.Unmarshal(result, metadata); err != nil {
		return nil, err
	}

	sortID := func(pkg rpmmd.Package) string {
		return fmt.Sprintf("%s-%s-%s", pkg.Name, pkg.Version, pkg.Release)
	}
	pkgs := metadata.Packages
	sort.Slice(pkgs, func(i, j int) bool {
		return sortID(pkgs[i]) < sortID(pkgs[j])
	})
	metadata.Packages = pkgs
	namedChecksums := make(map[string]string)
	for _, repo := range repos {
		namedChecksums[repo.Name] = metadata.Checksums[repo.Hash()]
	}
	metadata.Checksums = namedChecksums
	return metadata, nil
}

func (s *Solver) reposFromRPMMD(rpmRepos []rpmmd.RepoConfig) ([]repoConfig, error) {
	dnfRepos := make([]repoConfig, len(rpmRepos))
	for idx, rr := range rpmRepos {
		dr := repoConfig{
			ID:             rr.Hash(),
			Name:           rr.Name,
			BaseURL:        rr.BaseURL,
			Metalink:       rr.Metalink,
			MirrorList:     rr.MirrorList,
			GPGKey:         rr.GPGKey,
			IgnoreSSL:      rr.IgnoreSSL,
			MetadataExpire: rr.MetadataExpire,
		}
		if rr.RHSM {
			if s.subscriptions == nil {
				return nil, fmt.Errorf("This system does not have any valid subscriptions. Subscribe it before specifying rhsm: true in sources.")
			}
			secrets, err := s.subscriptions.GetSecretsForBaseurl(rr.BaseURL, s.arch, s.releaseVer)
			if err != nil {
				return nil, fmt.Errorf("RHSM secrets not found on the host for this baseurl: %s", rr.BaseURL)
			}
			dr.SSLCACert = secrets.SSLCACert
			dr.SSLClientKey = secrets.SSLClientKey
			dr.SSLClientCert = secrets.SSLClientCert
		}
		dnfRepos[idx] = dr
	}
	return dnfRepos, nil
}

// Repository configuration for resolving dependencies for a set of packages. A
// Solver needs at least one RPM repository configured to be able to depsolve.
type repoConfig struct {
	ID             string `json:"id"`
	Name           string `json:"name,omitempty"`
	BaseURL        string `json:"baseurl,omitempty"`
	Metalink       string `json:"metalink,omitempty"`
	MirrorList     string `json:"mirrorlist,omitempty"`
	GPGKey         string `json:"gpgkey,omitempty"`
	IgnoreSSL      bool   `json:"ignoressl"`
	SSLCACert      string `json:"sslcacert,omitempty"`
	SSLClientKey   string `json:"sslclientkey,omitempty"`
	SSLClientCert  string `json:"sslclientcert,omitempty"`
	MetadataExpire string `json:"metadata_expire,omitempty"`
}

// Helper function for creating a depsolve request payload.
// The request defines a sequence of transactions, each depsolving one of the
// elements of `pkgSets` in the order they appear.  The `repoConfigs` are used
// as the base repositories for all transactions.  The extra repository configs
// in `pkgsetsRepos` are used for each of the `pkgSets` with matching index.
// The length of `pkgsetsRepos` must match the length of `pkgSets` or be empty
// (nil or empty slice).
//
// NOTE: Due to implementation limitations of DNF and dnf-json, each package set
// in the chain must use all of the repositories used by its predecessor.
// An error is returned if this requirement is not met.
func (s *Solver) makeDepsolveRequest(pkgSets []rpmmd.PackageSet) (*Request, map[string]rpmmd.RepoConfig, error) {

	// dedupe repository configurations but maintain order
	// the order in which repositories are added to the request affects the
	// order of the dependencies in the result
	repos := make([]rpmmd.RepoConfig, 0)
	rpmRepoMap := make(map[string]rpmmd.RepoConfig)

	for _, ps := range pkgSets {
		for _, repo := range ps.Repositories {
			id := repo.Hash()
			if _, ok := rpmRepoMap[id]; !ok {
				rpmRepoMap[id] = repo
				repos = append(repos, repo)
			}
		}
	}

	transactions := make([]transactionArgs, len(pkgSets))
	for dsIdx, pkgSet := range pkgSets {
		transactions[dsIdx] = transactionArgs{
			PackageSpecs: pkgSet.Include,
			ExcludeSpecs: pkgSet.Exclude,
		}

		for _, jobRepo := range pkgSet.Repositories {
			transactions[dsIdx].RepoIDs = append(transactions[dsIdx].RepoIDs, jobRepo.Hash())
		}

		// If more than one transaction, ensure that the transaction uses
		// all of the repos from its predecessor
		if dsIdx > 0 {
			prevRepoIDs := transactions[dsIdx-1].RepoIDs
			if len(transactions[dsIdx].RepoIDs) < len(prevRepoIDs) {
				return nil, nil, fmt.Errorf("chained packageSet %d does not use all of the repos used by its predecessor", dsIdx)
			}

			for idx, repoID := range prevRepoIDs {
				if repoID != transactions[dsIdx].RepoIDs[idx] {
					return nil, nil, fmt.Errorf("chained packageSet %d does not use all of the repos used by its predecessor", dsIdx)
				}
			}
		}
	}

	dnfRepoMap, err := s.reposFromRPMMD(repos)
	if err != nil {
		return nil, nil, err
	}
	args := arguments{
		Repos:        dnfRepoMap,
		Transactions: transactions,
	}

	req := Request{
		Command:          "depsolve",
		ModulePlatformID: s.modulePlatformID,
		Arch:             s.arch,
		CacheDir:         s.cacheDir,
		Arguments:        args,
	}

	return &req, rpmRepoMap, nil
}

// Helper function for creating a dump request payload
func (s *Solver) makeDumpRequest(repos []rpmmd.RepoConfig) (*Request, error) {
	dnfRepos, err := s.reposFromRPMMD(repos)
	if err != nil {
		return nil, err
	}
	req := Request{
		Command:          "dump",
		ModulePlatformID: s.modulePlatformID,
		Arch:             s.arch,
		CacheDir:         s.cacheDir,
		Arguments: arguments{
			Repos: dnfRepos,
		},
	}
	return &req, nil
}

// convert an internal depsolveResult to a public DepsolveResult.
func resultToPublic(result *depsolveResult, repos map[string]rpmmd.RepoConfig) *DepsolveResult {
	return &DepsolveResult{
		Checksums:    result.Checksums,
		Dependencies: depsToRPMMD(result.Dependencies, repos),
	}
}

// convert internal a list of PackageSpecs to the rpmmd equivalent and attach
// key and subscription information based on the repository configs.
func depsToRPMMD(dependencies []PackageSpec, repos map[string]rpmmd.RepoConfig) []rpmmd.PackageSpec {
	rpmDependencies := make([]rpmmd.PackageSpec, len(dependencies))
	for i, dep := range dependencies {
		repo, ok := repos[dep.RepoID]
		if !ok {
			panic("dependency repo ID not found in repositories")
		}
		dep := dependencies[i]
		rpmDependencies[i].Name = dep.Name
		rpmDependencies[i].Epoch = dep.Epoch
		rpmDependencies[i].Version = dep.Version
		rpmDependencies[i].Release = dep.Release
		rpmDependencies[i].Arch = dep.Arch
		rpmDependencies[i].RemoteLocation = dep.RemoteLocation
		rpmDependencies[i].Checksum = dep.Checksum
		rpmDependencies[i].CheckGPG = repo.CheckGPG
		if repo.RHSM {
			rpmDependencies[i].Secrets = "org.osbuild.rhsm"
		}
	}
	return rpmDependencies
}

// Request command and arguments for dnf-json
type Request struct {
	// Command should be either "depsolve" or "dump"
	Command string `json:"command"`

	// Platform ID, e.g., "platform:el8"
	ModulePlatformID string `json:"module_platform_id"`

	// System architecture
	Arch string `json:"arch"`

	// Cache directory for the DNF metadata
	CacheDir string `json:"cachedir"`

	// Arguments for the action defined by Command
	Arguments arguments `json:"arguments"`
}

// arguments for a dnf-json request
type arguments struct {
	// Repositories to use for depsolving
	Repos []repoConfig `json:"repos"`

	// Depsolve package sets and repository mappings for this request
	Transactions []transactionArgs `json:"transactions"`
}

type transactionArgs struct {
	// Packages to depsolve
	PackageSpecs []string `json:"package-specs"`

	// Packages to exclude from results
	ExcludeSpecs []string `json:"exclude-specs"`

	// IDs of repositories to use for this depsolve
	RepoIDs []string `json:"repo-ids"`
}

// Private version of the depsolve result.  Uses a slightly different
// PackageSpec than the public one that uses the rpmmd type.
type depsolveResult struct {
	// Repository checksums
	Checksums map[string]string `json:"checksums"`

	// Resolved package dependencies
	Dependencies []PackageSpec `json:"dependencies"`
}

// DepsolveResult is the result returned from a Depsolve call.
type DepsolveResult struct {
	// Repository checksums
	Checksums map[string]string

	// Resolved package dependencies
	Dependencies []rpmmd.PackageSpec
}

// FetchMetadataResult is the result returned from a FetchMetadata call.
type FetchMetadataResult struct {
	Checksums map[string]string `json:"checksums"`
	Packages  rpmmd.PackageList `json:"packages"`
}

// Package specification
type PackageSpec struct {
	Name           string `json:"name"`
	Epoch          uint   `json:"epoch"`
	Version        string `json:"version,omitempty"`
	Release        string `json:"release,omitempty"`
	Arch           string `json:"arch,omitempty"`
	RepoID         string `json:"repo_id,omitempty"`
	Path           string `json:"path,omitempty"`
	RemoteLocation string `json:"remote_location,omitempty"`
	Checksum       string `json:"checksum,omitempty"`
	Secrets        string `json:"secrets,omitempty"`
}

// dnf-json error structure
type Error struct {
	Kind   string `json:"kind"`
	Reason string `json:"reason"`
}

func (err Error) Error() string {
	return fmt.Sprintf("DNF error occurred: %s: %s", err.Kind, err.Reason)
}

// parseError parses the response from dnf-json into the Error type.
func parseError(data []byte) Error {
	var e Error
	if err := json.Unmarshal(data, &e); err != nil {
		// dumping the error into the Reason can get noisy, but it's good for troubleshooting
		return Error{
			Kind:   "InternalError",
			Reason: fmt.Sprintf("Failed to unmarshal dnf-json error output %q: %s", string(data), err.Error()),
		}
	}
	return e
}

func run(dnfJsonCmd []string, req *Request) ([]byte, error) {
	if len(dnfJsonCmd) == 0 {
		return nil, fmt.Errorf("dnf-json command undefined")
	}
	ex := dnfJsonCmd[0]
	args := make([]string, len(dnfJsonCmd)-1)
	if len(dnfJsonCmd) > 1 {
		args = dnfJsonCmd[1:]
	}
	cmd := exec.Command(ex, args...)
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}

	cmd.Stderr = os.Stderr
	stdout := new(bytes.Buffer)
	cmd.Stdout = stdout

	err = cmd.Start()
	if err != nil {
		return nil, err
	}

	err = json.NewEncoder(stdin).Encode(req)
	if err != nil {
		return nil, err
	}
	stdin.Close()

	err = cmd.Wait()
	output := stdout.Bytes()
	if runError, ok := err.(*exec.ExitError); ok && runError.ExitCode() != 0 {
		return nil, parseError(output)
	}

	return output, nil
}
