#!/usr/bin/env bash

# create an ibc connection

set -euo pipefail

function usage() {
    >&2 echo "Usage: ${0} <relayer_endpoint> <polymer_l2_1_json_path> <polymer_l2_2_json_path>"
    exit 1
}

if (( $# != 3 )); then
    usage
fi

RELAYER_ENDPOINT="${1}" # example localhost:4001
POLYMER_L2_1_JSON="${2}"
POLYMER_L2_2_JSON="${3}"
ORDERING=0

receiverA="$( jq -r .polymer_receiver_address ${POLYMER_L2_1_JSON} )"
receiverB="$( jq -r .polymer_receiver_address ${POLYMER_L2_2_JSON} )"

data=$( jq -n \
   --arg receiver "${receiverA}"  \
   --arg portid "polyibc.op2.${receiverB:2}" \
   --arg ordering ${ORDERING} \
   '{
          "receiverAddress": $receiver,
          "version": "1.0",
          "ordering": $ordering|tonumber,
          "feeEnabled": false,
          "connectionHops": ["connection-2", "connection-1"],
          "counterparty": {
              "portID": $portid
          }
    }')

curl "${RELAYER_ENDPOINT}/paths/op-polymer-1/createChannel" -d "${data}"
