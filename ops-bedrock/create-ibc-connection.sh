#!/usr/bin/env bash

set -euo pipefail

# create ibc connection

function usage() {
    >&2 echo "Usage: ${0} <relayer_endpoint> <chain_id>"
    exit 1
}

if (( $# != 2 )); then
    usage
fi

RELAYER_ENDPOINT="${1}" # example localhost:4001
CHAIN_ID="${2}"
data="{\"chainID\":\"${CHAIN_ID}\"}\"

curl "${RELAYER_ENDPOINT}/paths/op-polymer-1/createConnection" -d ${data}
