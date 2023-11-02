/*
Pulp 3 API

Fetch, Upload, Organize, and Distribute Software Packages

API version: v3
Contact: pulp-list@redhat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pulpclient

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"reflect"
)


// PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPIService PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPI service
type PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPIService service

type PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPIPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexDeleteRequest struct {
	ctx context.Context
	ApiService *PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPIService
	distroBasePath string
	name string
	namespace string
	path string
}

func (r PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPIPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexDeleteRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexDeleteExecute(r)
}

/*
PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexDelete Method for PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexDelete

Trigger an asynchronous delete task

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param distroBasePath
 @param name
 @param namespace
 @param path
 @return PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPIPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexDeleteRequest
*/
func (a *PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPIService) PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexDelete(ctx context.Context, distroBasePath string, name string, namespace string, path string) PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPIPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexDeleteRequest {
	return PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPIPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexDeleteRequest{
		ApiService: a,
		ctx: ctx,
		distroBasePath: distroBasePath,
		name: name,
		namespace: namespace,
		path: path,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPIService) PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexDeleteExecute(r PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPIPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexDeleteRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPIService.PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp_ansible/galaxy/{path}/api/v3/plugin/ansible/content/{distro_base_path}/collections/index/{namespace}/{name}/"
	localVarPath = strings.Replace(localVarPath, "{"+"distro_base_path"+"}", parameterValueToString(r.distroBasePath, "distroBasePath"), -1)  // NOTE: paths aren't escaped because Pulp uses hrefs as path parameters
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", parameterValueToString(r.name, "name"), -1)  // NOTE: paths aren't escaped because Pulp uses hrefs as path parameters
	localVarPath = strings.Replace(localVarPath, "{"+"namespace"+"}", parameterValueToString(r.namespace, "namespace"), -1)  // NOTE: paths aren't escaped because Pulp uses hrefs as path parameters
	localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", parameterValueToString(r.path, "path"), -1)  // NOTE: paths aren't escaped because Pulp uses hrefs as path parameters

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPIPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexListRequest struct {
	ctx context.Context
	ApiService *PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPIService
	distroBasePath string
	path string
	deprecated *bool
	limit *int32
	name *string
	namespace *string
	offset *int32
	ordering *[]string
	pulpHrefIn *[]string
	pulpIdIn *[]string
	fields *[]string
	excludeFields *[]string
}

func (r PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPIPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexListRequest) Deprecated(deprecated bool) PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPIPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexListRequest {
	r.deprecated = &deprecated
	return r
}

// Number of results to return per page.
func (r PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPIPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexListRequest) Limit(limit int32) PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPIPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexListRequest {
	r.limit = &limit
	return r
}

func (r PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPIPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexListRequest) Name(name string) PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPIPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexListRequest {
	r.name = &name
	return r
}

func (r PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPIPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexListRequest) Namespace(namespace string) PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPIPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexListRequest {
	r.namespace = &namespace
	return r
}

// The initial index from which to return the results.
func (r PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPIPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexListRequest) Offset(offset int32) PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPIPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexListRequest {
	r.offset = &offset
	return r
}

// Ordering  * &#x60;pulp_id&#x60; - Pulp id * &#x60;-pulp_id&#x60; - Pulp id (descending) * &#x60;pulp_created&#x60; - Pulp created * &#x60;-pulp_created&#x60; - Pulp created (descending) * &#x60;pulp_last_updated&#x60; - Pulp last updated * &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending) * &#x60;namespace&#x60; - Namespace * &#x60;-namespace&#x60; - Namespace (descending) * &#x60;name&#x60; - Name * &#x60;-name&#x60; - Name (descending) * &#x60;pk&#x60; - Pk * &#x60;-pk&#x60; - Pk (descending)
func (r PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPIPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexListRequest) Ordering(ordering []string) PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPIPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexListRequest {
	r.ordering = &ordering
	return r
}

// Multiple values may be separated by commas.
func (r PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPIPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexListRequest) PulpHrefIn(pulpHrefIn []string) PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPIPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexListRequest {
	r.pulpHrefIn = &pulpHrefIn
	return r
}

// Multiple values may be separated by commas.
func (r PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPIPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexListRequest) PulpIdIn(pulpIdIn []string) PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPIPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexListRequest {
	r.pulpIdIn = &pulpIdIn
	return r
}

// A list of fields to include in the response.
func (r PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPIPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexListRequest) Fields(fields []string) PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPIPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPIPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexListRequest) ExcludeFields(excludeFields []string) PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPIPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPIPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexListRequest) Execute() (*PaginatedCollectionResponseList, *http.Response, error) {
	return r.ApiService.PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexListExecute(r)
}

/*
PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexList Method for PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexList

ViewSet for Collections.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param distroBasePath
 @param path
 @return PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPIPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexListRequest
*/
func (a *PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPIService) PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexList(ctx context.Context, distroBasePath string, path string) PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPIPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexListRequest {
	return PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPIPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexListRequest{
		ApiService: a,
		ctx: ctx,
		distroBasePath: distroBasePath,
		path: path,
	}
}

// Execute executes the request
//  @return PaginatedCollectionResponseList
func (a *PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPIService) PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexListExecute(r PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPIPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexListRequest) (*PaginatedCollectionResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedCollectionResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPIService.PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp_ansible/galaxy/{path}/api/v3/plugin/ansible/content/{distro_base_path}/collections/index/"
	localVarPath = strings.Replace(localVarPath, "{"+"distro_base_path"+"}", parameterValueToString(r.distroBasePath, "distroBasePath"), -1)  // NOTE: paths aren't escaped because Pulp uses hrefs as path parameters
	localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", parameterValueToString(r.path, "path"), -1)  // NOTE: paths aren't escaped because Pulp uses hrefs as path parameters

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.deprecated != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "deprecated", r.deprecated, "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.name != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name", r.name, "")
	}
	if r.namespace != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "namespace", r.namespace, "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
	}
	if r.ordering != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ordering", r.ordering, "csv")
	}
	if r.pulpHrefIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_href__in", r.pulpHrefIn, "csv")
	}
	if r.pulpIdIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_id__in", r.pulpIdIn, "csv")
	}
	if r.fields != nil {
		t := *r.fields
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "fields", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "fields", t, "multi")
		}
	}
	if r.excludeFields != nil {
		t := *r.excludeFields
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "exclude_fields", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "exclude_fields", t, "multi")
		}
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPIPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexReadRequest struct {
	ctx context.Context
	ApiService *PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPIService
	distroBasePath string
	name string
	namespace string
	path string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPIPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexReadRequest) Fields(fields []string) PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPIPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPIPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexReadRequest) ExcludeFields(excludeFields []string) PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPIPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPIPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexReadRequest) Execute() (*CollectionResponse, *http.Response, error) {
	return r.ApiService.PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexReadExecute(r)
}

/*
PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexRead Method for PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexRead

ViewSet for Collections.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param distroBasePath
 @param name
 @param namespace
 @param path
 @return PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPIPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexReadRequest
*/
func (a *PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPIService) PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexRead(ctx context.Context, distroBasePath string, name string, namespace string, path string) PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPIPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexReadRequest {
	return PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPIPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexReadRequest{
		ApiService: a,
		ctx: ctx,
		distroBasePath: distroBasePath,
		name: name,
		namespace: namespace,
		path: path,
	}
}

// Execute executes the request
//  @return CollectionResponse
func (a *PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPIService) PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexReadExecute(r PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPIPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexReadRequest) (*CollectionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CollectionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPIService.PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp_ansible/galaxy/{path}/api/v3/plugin/ansible/content/{distro_base_path}/collections/index/{namespace}/{name}/"
	localVarPath = strings.Replace(localVarPath, "{"+"distro_base_path"+"}", parameterValueToString(r.distroBasePath, "distroBasePath"), -1)  // NOTE: paths aren't escaped because Pulp uses hrefs as path parameters
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", parameterValueToString(r.name, "name"), -1)  // NOTE: paths aren't escaped because Pulp uses hrefs as path parameters
	localVarPath = strings.Replace(localVarPath, "{"+"namespace"+"}", parameterValueToString(r.namespace, "namespace"), -1)  // NOTE: paths aren't escaped because Pulp uses hrefs as path parameters
	localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", parameterValueToString(r.path, "path"), -1)  // NOTE: paths aren't escaped because Pulp uses hrefs as path parameters

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fields != nil {
		t := *r.fields
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "fields", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "fields", t, "multi")
		}
	}
	if r.excludeFields != nil {
		t := *r.excludeFields
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "exclude_fields", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "exclude_fields", t, "multi")
		}
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPIPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexUpdateRequest struct {
	ctx context.Context
	ApiService *PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPIService
	distroBasePath string
	name string
	namespace string
	path string
	patchedCollection *PatchedCollection
}

func (r PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPIPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexUpdateRequest) PatchedCollection(patchedCollection PatchedCollection) PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPIPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexUpdateRequest {
	r.patchedCollection = &patchedCollection
	return r
}

func (r PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPIPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexUpdateRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexUpdateExecute(r)
}

/*
PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexUpdate Method for PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexUpdate

Trigger an asynchronous update task

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param distroBasePath
 @param name
 @param namespace
 @param path
 @return PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPIPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexUpdateRequest
*/
func (a *PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPIService) PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexUpdate(ctx context.Context, distroBasePath string, name string, namespace string, path string) PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPIPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexUpdateRequest {
	return PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPIPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexUpdateRequest{
		ApiService: a,
		ctx: ctx,
		distroBasePath: distroBasePath,
		name: name,
		namespace: namespace,
		path: path,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPIService) PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexUpdateExecute(r PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPIPulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexUpdateRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexAPIService.PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp_ansible/galaxy/{path}/api/v3/plugin/ansible/content/{distro_base_path}/collections/index/{namespace}/{name}/"
	localVarPath = strings.Replace(localVarPath, "{"+"distro_base_path"+"}", parameterValueToString(r.distroBasePath, "distroBasePath"), -1)  // NOTE: paths aren't escaped because Pulp uses hrefs as path parameters
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", parameterValueToString(r.name, "name"), -1)  // NOTE: paths aren't escaped because Pulp uses hrefs as path parameters
	localVarPath = strings.Replace(localVarPath, "{"+"namespace"+"}", parameterValueToString(r.namespace, "namespace"), -1)  // NOTE: paths aren't escaped because Pulp uses hrefs as path parameters
	localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", parameterValueToString(r.path, "path"), -1)  // NOTE: paths aren't escaped because Pulp uses hrefs as path parameters

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.patchedCollection == nil {
		return localVarReturnValue, nil, reportError("patchedCollection is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "application/x-www-form-urlencoded", "multipart/form-data"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.patchedCollection
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}