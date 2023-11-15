Devnet
======

Setup
-----
`make devnet-up`

Reset
-----
`docker kill $(docker ps -a --filter name=^/ops-bedrock -q)`

`docker volume rm $(docker volume ls --filter='name=ops-bedrock*' -q)`

`rm -r .devnet`

IBC Setup
=========

Create IBC connections
----------------------
`./create-ibc-connection.sh localhost:4001 901`
`./create-ibc-connection.sh localhost:4001 902`

Create IBC channel
------------------
`./create-ibc-channel.sh localhost:4001 ../.devnet/polymer-l2-1.json ../.devnet/polymer-l2-2.json`

Send IBC Packet
---------------
`./send-ibc-packet.sh http://127.0.0.1:9545/ ../.devnet/polymer-l2-1.json channel-0 hello`

Query VIBC Channel
------------------
`./query-vibc-channel.sh http://127.0.0.1:9545 ../.devnet/polymer-l2-1.json channel-0`

