#!/bin/sh

# create ibc connection

CHAIN_ID="${1}"
data="{\"chainID\":\"${CHAIN_ID}\"}\"

curl localhost:4001/paths/op-polymer-1/createConnection -d ${data}
