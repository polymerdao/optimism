#!/usr/bin/env bash

# query vibc channel state

set -euo pipefail

function usage() {
    >&2 echo "Usage: ${0} <rpc_endpoint> <polymer_l2_json_path> <channel_id>"
    exit 1
}

if (( $# != 3 )); then
    usage
fi

RPC_ENDPOINT="${1}" # example: http://127.0.0.1:9545
POLYMER_L2_JSON="${2}"
CHANNEL_ID="${3}"
dispatcher="$( jq -r .polymer_dispatcher_address ${POLYMER_L2_JSON} )"
receiver="$( jq -r .polymer_receiver_address ${POLYMER_L2_JSON} )"
channel_bytes=$(cast format-bytes32-string "${CHANNEL_ID}")

cast call --rpc-url "${RPC_ENDPOINT}" "${dispatcher}" "getChannel(address,bytes32)((string,uint8,bool,string[],string,bytes32))" "${dispatcher}" "${channel_bytes}"
