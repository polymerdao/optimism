#!/bin/sh

# create an ibc connection

POLYMER_L2_1_JSON="${1}"
POLYMER_L2_2_JSON="${2}"

receiverA="$( jq -r .polymer_receiver_address ${POLYMER_L2_1_JSON} )"
receiverB="$( jq -r .polymer_receiver_address ${POLYMER_L2_2_JSON} )"

data=$( jq -n \
   --arg receiver "$receiverA"  \
   --arg portid "polyibc.op2.${receiverB:2}" \
   '{
          "receiverAddress": $receiver,
          "version": "1.0",
          "ordering": 0,
          "feeEnabled": false,
          "connectionHops": ["connection-2", "connection-1"],
          "counterparty": {
              "portID": $portid
          }
    }')

curl localhost:4001/paths/op-polymer-1/createChannel -d "${data}"
