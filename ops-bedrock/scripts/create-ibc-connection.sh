#!/usr/bin/env bash

set -euo pipefail

# create ibc connection

function usage() {
    >&2 echo "Usage: ${0} <relayer_endpoint> <chain_id_1> <chain_id_2>"
    exit 1
}

if (( $# != 3 )); then
    usage
fi

RELAYER_ENDPOINT="${1}" # example localhost:4001
CHAIN_ID1="${2}"
CHAIN_ID2="${3}"
data_1="{\"chainID\":\"${CHAIN_ID1}\"}"
data_2="{\"chainID\":\"${CHAIN_ID2}\"}"

curl "${RELAYER_ENDPOINT}/paths/op-polymer-1/createConnection" -d ${data_1}
curl "${RELAYER_ENDPOINT}/paths/op-polymer-2/createConnection" -d ${data_2}
