#!/bin/bash

set -euxo pipefail

WORKSPACE=$(mktemp -d)
function cleanup() {
    rm -rf "$WORKSPACE"
}
trap cleanup EXIT

function get_build_info() {
    key="$1"
    fname="$2"
    if rpm -q --quiet weldr-client; then
        key=".body${key}"
    fi
    jq -r "${key}" "${fname}"
}

# Get the compose log.
get_compose_log () {
    COMPOSE_ID=$1
    LOG_FILE=/root/logs/osbuild-upgrade8to9.log

    # Download the logs.
    composer-cli compose log "$COMPOSE_ID" | tee "$LOG_FILE" > /dev/null
}

# Get the compose metadata.
get_compose_metadata () {
    COMPOSE_ID=$1
    METADATA_FILE=/root/logs/osbuild-upgrade8to9.json

    # Download the metadata.
    composer-cli compose metadata "$COMPOSE_ID" > /dev/null

    # Find the tarball and extract it.
    TARBALL=$(basename "$(find . -maxdepth 1 -type f -name "*-metadata.tar")")
    tar -xf "$TARBALL"
    rm -f "$TARBALL"

    # Move the JSON file into place.
    jq -M '.' "${COMPOSE_ID}".json | tee "$METADATA_FILE" > /dev/null
}

IMAGE_KEY=osbuild-composer-upgrade-test
COMPOSE_START=${WORKSPACE}/compose-start-${IMAGE_KEY}.json
COMPOSE_INFO=${WORKSPACE}/compose-info-${IMAGE_KEY}.json

# check installed osbuild-composer version
rpm -qi osbuild-composer

# install jq in case it's not present
dnf install -y jq

# Prepare repository override
mkdir -p /etc/osbuild-composer/repositories
tee /etc/osbuild-composer/repositories/rhel-90.json > /dev/null << EOF
{
    "x86_64": [
        {
            "name": "baseos",
            "baseurl": "http://download.devel.redhat.com/rhel-9/nightly/RHEL-9/latest-RHEL-9.0.0/compose/BaseOS/x86_64/os/",
            "gpgkey": "-----BEGIN PGP PUBLIC KEY BLOCK-----\n\nmQINBErgSTsBEACh2A4b0O9t+vzC9VrVtL1AKvUWi9OPCjkvR7Xd8DtJxeeMZ5eF\n0HtzIG58qDRybwUe89FZprB1ffuUKzdE+HcL3FbNWSSOXVjZIersdXyH3NvnLLLF\n0DNRB2ix3bXG9Rh/RXpFsNxDp2CEMdUvbYCzE79K1EnUTVh1L0Of023FtPSZXX0c\nu7Pb5DI5lX5YeoXO6RoodrIGYJsVBQWnrWw4xNTconUfNPk0EGZtEnzvH2zyPoJh\nXGF+Ncu9XwbalnYde10OCvSWAZ5zTCpoLMTvQjWpbCdWXJzCm6G+/hx9upke546H\n5IjtYm4dTIVTnc3wvDiODgBKRzOl9rEOCIgOuGtDxRxcQkjrC+xvg5Vkqn7vBUyW\n9pHedOU+PoF3DGOM+dqv+eNKBvh9YF9ugFAQBkcG7viZgvGEMGGUpzNgN7XnS1gj\n/DPo9mZESOYnKceve2tIC87p2hqjrxOHuI7fkZYeNIcAoa83rBltFXaBDYhWAKS1\nPcXS1/7JzP0ky7d0L6Xbu/If5kqWQpKwUInXtySRkuraVfuK3Bpa+X1XecWi24JY\nHVtlNX025xx1ewVzGNCTlWn1skQN2OOoQTV4C8/qFpTW6DTWYurd4+fE0OJFJZQF\nbuhfXYwmRlVOgN5i77NTIJZJQfYFj38c/Iv5vZBPokO6mffrOTv3MHWVgQARAQAB\ntDNSZWQgSGF0LCBJbmMuIChyZWxlYXNlIGtleSAyKSA8c2VjdXJpdHlAcmVkaGF0\nLmNvbT6JAjYEEwECACAFAkrgSTsCGwMGCwkIBwMCBBUCCAMEFgIDAQIeAQIXgAAK\nCRAZni+R/UMdUWzpD/9s5SFR/ZF3yjY5VLUFLMXIKUztNN3oc45fyLdTI3+UClKC\n2tEruzYjqNHhqAEXa2sN1fMrsuKec61Ll2NfvJjkLKDvgVIh7kM7aslNYVOP6BTf\nC/JJ7/ufz3UZmyViH/WDl+AYdgk3JqCIO5w5ryrC9IyBzYv2m0HqYbWfphY3uHw5\nun3ndLJcu8+BGP5F+ONQEGl+DRH58Il9Jp3HwbRa7dvkPgEhfFR+1hI+Btta2C7E\n0/2NKzCxZw7Lx3PBRcU92YKyaEihfy/aQKZCAuyfKiMvsmzs+4poIX7I9NQCJpyE\nIGfINoZ7VxqHwRn/d5mw2MZTJjbzSf+Um9YJyA0iEEyD6qjriWQRbuxpQXmlAJbh\n8okZ4gbVFv1F8MzK+4R8VvWJ0XxgtikSo72fHjwha7MAjqFnOq6eo6fEC/75g3NL\nGht5VdpGuHk0vbdENHMC8wS99e5qXGNDued3hlTavDMlEAHl34q2H9nakTGRF5Ki\nJUfNh3DVRGhg8cMIti21njiRh7gyFI2OccATY7bBSr79JhuNwelHuxLrCFpY7V25\nOFktl15jZJaMxuQBqYdBgSay2G0U6D1+7VsWufpzd/Abx1/c3oi9ZaJvW22kAggq\ndzdA27UUYjWvx42w9menJwh/0jeQcTecIUd0d0rFcw/c1pvgMMl/Q73yzKgKYw==\n=zbHE\n-----END PGP PUBLIC KEY BLOCK-----\n-----BEGIN PGP PUBLIC KEY BLOCK-----\n\n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\nvt61kpkXoRGxjTlc6h+69aidSg==\n=ls8J\n-----END PGP PUBLIC KEY BLOCK-----\n"
        },
        {
            "name": "appstream",
            "baseurl": "http://download.devel.redhat.com/rhel-9/nightly/RHEL-9/latest-RHEL-9.0.0/compose/AppStream/x86_64/os/",
            "gpgkey": "-----BEGIN PGP PUBLIC KEY BLOCK-----\n\nmQINBErgSTsBEACh2A4b0O9t+vzC9VrVtL1AKvUWi9OPCjkvR7Xd8DtJxeeMZ5eF\n0HtzIG58qDRybwUe89FZprB1ffuUKzdE+HcL3FbNWSSOXVjZIersdXyH3NvnLLLF\n0DNRB2ix3bXG9Rh/RXpFsNxDp2CEMdUvbYCzE79K1EnUTVh1L0Of023FtPSZXX0c\nu7Pb5DI5lX5YeoXO6RoodrIGYJsVBQWnrWw4xNTconUfNPk0EGZtEnzvH2zyPoJh\nXGF+Ncu9XwbalnYde10OCvSWAZ5zTCpoLMTvQjWpbCdWXJzCm6G+/hx9upke546H\n5IjtYm4dTIVTnc3wvDiODgBKRzOl9rEOCIgOuGtDxRxcQkjrC+xvg5Vkqn7vBUyW\n9pHedOU+PoF3DGOM+dqv+eNKBvh9YF9ugFAQBkcG7viZgvGEMGGUpzNgN7XnS1gj\n/DPo9mZESOYnKceve2tIC87p2hqjrxOHuI7fkZYeNIcAoa83rBltFXaBDYhWAKS1\nPcXS1/7JzP0ky7d0L6Xbu/If5kqWQpKwUInXtySRkuraVfuK3Bpa+X1XecWi24JY\nHVtlNX025xx1ewVzGNCTlWn1skQN2OOoQTV4C8/qFpTW6DTWYurd4+fE0OJFJZQF\nbuhfXYwmRlVOgN5i77NTIJZJQfYFj38c/Iv5vZBPokO6mffrOTv3MHWVgQARAQAB\ntDNSZWQgSGF0LCBJbmMuIChyZWxlYXNlIGtleSAyKSA8c2VjdXJpdHlAcmVkaGF0\nLmNvbT6JAjYEEwECACAFAkrgSTsCGwMGCwkIBwMCBBUCCAMEFgIDAQIeAQIXgAAK\nCRAZni+R/UMdUWzpD/9s5SFR/ZF3yjY5VLUFLMXIKUztNN3oc45fyLdTI3+UClKC\n2tEruzYjqNHhqAEXa2sN1fMrsuKec61Ll2NfvJjkLKDvgVIh7kM7aslNYVOP6BTf\nC/JJ7/ufz3UZmyViH/WDl+AYdgk3JqCIO5w5ryrC9IyBzYv2m0HqYbWfphY3uHw5\nun3ndLJcu8+BGP5F+ONQEGl+DRH58Il9Jp3HwbRa7dvkPgEhfFR+1hI+Btta2C7E\n0/2NKzCxZw7Lx3PBRcU92YKyaEihfy/aQKZCAuyfKiMvsmzs+4poIX7I9NQCJpyE\nIGfINoZ7VxqHwRn/d5mw2MZTJjbzSf+Um9YJyA0iEEyD6qjriWQRbuxpQXmlAJbh\n8okZ4gbVFv1F8MzK+4R8VvWJ0XxgtikSo72fHjwha7MAjqFnOq6eo6fEC/75g3NL\nGht5VdpGuHk0vbdENHMC8wS99e5qXGNDued3hlTavDMlEAHl34q2H9nakTGRF5Ki\nJUfNh3DVRGhg8cMIti21njiRh7gyFI2OccATY7bBSr79JhuNwelHuxLrCFpY7V25\nOFktl15jZJaMxuQBqYdBgSay2G0U6D1+7VsWufpzd/Abx1/c3oi9ZaJvW22kAggq\ndzdA27UUYjWvx42w9menJwh/0jeQcTecIUd0d0rFcw/c1pvgMMl/Q73yzKgKYw==\n=zbHE\n-----END PGP PUBLIC KEY BLOCK-----\n-----BEGIN PGP PUBLIC KEY BLOCK-----\n\n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\nvt61kpkXoRGxjTlc6h+69aidSg==\n=ls8J\n-----END PGP PUBLIC KEY BLOCK-----\n"
        }
        ,{
            "name": "rt",
            "baseurl": "http://download.devel.redhat.com/rhel-9/nightly/RHEL-9/latest-RHEL-9.0.0/compose/RT/x86_64/os/",
            "gpgkey": "-----BEGIN PGP PUBLIC KEY BLOCK-----\n\nmQINBErgSTsBEACh2A4b0O9t+vzC9VrVtL1AKvUWi9OPCjkvR7Xd8DtJxeeMZ5eF\n0HtzIG58qDRybwUe89FZprB1ffuUKzdE+HcL3FbNWSSOXVjZIersdXyH3NvnLLLF\n0DNRB2ix3bXG9Rh/RXpFsNxDp2CEMdUvbYCzE79K1EnUTVh1L0Of023FtPSZXX0c\nu7Pb5DI5lX5YeoXO6RoodrIGYJsVBQWnrWw4xNTconUfNPk0EGZtEnzvH2zyPoJh\nXGF+Ncu9XwbalnYde10OCvSWAZ5zTCpoLMTvQjWpbCdWXJzCm6G+/hx9upke546H\n5IjtYm4dTIVTnc3wvDiODgBKRzOl9rEOCIgOuGtDxRxcQkjrC+xvg5Vkqn7vBUyW\n9pHedOU+PoF3DGOM+dqv+eNKBvh9YF9ugFAQBkcG7viZgvGEMGGUpzNgN7XnS1gj\n/DPo9mZESOYnKceve2tIC87p2hqjrxOHuI7fkZYeNIcAoa83rBltFXaBDYhWAKS1\nPcXS1/7JzP0ky7d0L6Xbu/If5kqWQpKwUInXtySRkuraVfuK3Bpa+X1XecWi24JY\nHVtlNX025xx1ewVzGNCTlWn1skQN2OOoQTV4C8/qFpTW6DTWYurd4+fE0OJFJZQF\nbuhfXYwmRlVOgN5i77NTIJZJQfYFj38c/Iv5vZBPokO6mffrOTv3MHWVgQARAQAB\ntDNSZWQgSGF0LCBJbmMuIChyZWxlYXNlIGtleSAyKSA8c2VjdXJpdHlAcmVkaGF0\nLmNvbT6JAjYEEwECACAFAkrgSTsCGwMGCwkIBwMCBBUCCAMEFgIDAQIeAQIXgAAK\nCRAZni+R/UMdUWzpD/9s5SFR/ZF3yjY5VLUFLMXIKUztNN3oc45fyLdTI3+UClKC\n2tEruzYjqNHhqAEXa2sN1fMrsuKec61Ll2NfvJjkLKDvgVIh7kM7aslNYVOP6BTf\nC/JJ7/ufz3UZmyViH/WDl+AYdgk3JqCIO5w5ryrC9IyBzYv2m0HqYbWfphY3uHw5\nun3ndLJcu8+BGP5F+ONQEGl+DRH58Il9Jp3HwbRa7dvkPgEhfFR+1hI+Btta2C7E\n0/2NKzCxZw7Lx3PBRcU92YKyaEihfy/aQKZCAuyfKiMvsmzs+4poIX7I9NQCJpyE\nIGfINoZ7VxqHwRn/d5mw2MZTJjbzSf+Um9YJyA0iEEyD6qjriWQRbuxpQXmlAJbh\n8okZ4gbVFv1F8MzK+4R8VvWJ0XxgtikSo72fHjwha7MAjqFnOq6eo6fEC/75g3NL\nGht5VdpGuHk0vbdENHMC8wS99e5qXGNDued3hlTavDMlEAHl34q2H9nakTGRF5Ki\nJUfNh3DVRGhg8cMIti21njiRh7gyFI2OccATY7bBSr79JhuNwelHuxLrCFpY7V25\nOFktl15jZJaMxuQBqYdBgSay2G0U6D1+7VsWufpzd/Abx1/c3oi9ZaJvW22kAggq\ndzdA27UUYjWvx42w9menJwh/0jeQcTecIUd0d0rFcw/c1pvgMMl/Q73yzKgKYw==\n=zbHE\n-----END PGP PUBLIC KEY BLOCK-----\n-----BEGIN PGP PUBLIC KEY BLOCK-----\n\n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\nvt61kpkXoRGxjTlc6h+69aidSg==\n=ls8J\n-----END PGP PUBLIC KEY BLOCK-----\n"
        }
    ]
}
EOF

# start osbuild-compsoer socket
systemctl start osbuild-composer.socket

# prepare a simple blueprint
tee blueprint.toml > /dev/null << EOF
name = "bash"
description = "A base system with bash"
version = "0.0.1"

[[packages]]
name = "bash"
EOF

# push and depsolve the blueprint
composer-cli blueprints push blueprint.toml
composer-cli blueprints despsolve bash

# build a qcow image to verify functionality
composer-cli --json compose start bash qcow2 | tee "$COMPOSE_START"
COMPOSE_ID=$(get_build_info ".build_id" "$COMPOSE_START")

# Wait for the compose to finish.
while true; do
    composer-cli --json compose info "${COMPOSE_ID}" | tee "$COMPOSE_INFO" > /dev/null
    COMPOSE_STATUS=$(get_build_info ".queue_status" "$COMPOSE_INFO")

    # Is the compose finished?
    if [[ $COMPOSE_STATUS != RUNNING ]] && [[ $COMPOSE_STATUS != WAITING ]]; then
        break
    fi

    # Wait 30 seconds and try again.
    sleep 30
done

# Capture the compose logs from osbuild.
mkdir /root/logs
get_compose_log "$COMPOSE_ID"
get_compose_metadata "$COMPOSE_ID"

# if the compose succeds consider the test pass
if [[ $COMPOSE_STATUS != FINISHED ]]; then
    echo "Something went wrong with the compose. 😢"
    exit 1
else
    echo "Image has been built successfully!"
    exit 0
fi