#!/usr/bin/env bash

# send a test ibc packet

set -euo pipefail

function usage() {
    >&2 echo "Usage: ${0} <rpc_endpoint> <polymer_l2_json_path> <channel_id> <message>"
    exit 1
}

if (( $# != 4 )); then
    usage
fi


RPC_ENDPOINT="${1}" # example: http://127.0.0.1:9545
POLYMER_L2_JSON="${2}"
CHANNEL_ID="${3}"
MESSAGE="${4}"
PRIVATE_KEY='0xbf7604d9d3a1c7748642b1b7b05c2bd219c9faa91458b370f85e5a40f3b03af7'
dispatcher="$( jq -r .polymer_dispatcher_address ${POLYMER_L2_JSON} )"
receiver="$( jq -r .polymer_receiver_address ${POLYMER_L2_JSON} )"
channel_bytes=$(cast format-bytes32-string "${CHANNEL_ID}")

# send a greeting packet
cast send --rpc-url "${RPC_ENDPOINT}" \
     --value 0.01ether \
     --private-key "${PRIVATE_KEY}" \
     --gas-limit 1000000 \
     "${receiver}" \
     "greet(address,string,bytes32,uint64,(uint256,uint256,uint256))" \
     "${dispatcher}" "${MESSAGE}" "${channel_bytes}" 1699392755000000000 "(0,0,0)"

echo "Verify  transaction: cast tx ${tx_hash}"
