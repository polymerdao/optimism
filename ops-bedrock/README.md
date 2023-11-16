Devnet
======

Init Contracts
--------------
`make build-ts`


Setup
-----
`make devnet-up`

Reset
-----
`scripts/clean-devnet.sh`

IBC Setup
=========

Create IBC connections
----------------------
`./scripts/create-ibc-connection.sh localhost:4001 901 902`

Create IBC channel
------------------
`./scripts/create-ibc-channel.sh localhost:4001 ../.devnet/polymer-l2-1.json ../.devnet/polymer-l2-2.json`

Send IBC Packet
---------------
`./scripts/send-ibc-packet.sh http://127.0.0.1:9545/ ../.devnet/polymer-l2-1.json channel-0 hello`

Query VIBC Channel
------------------
`./scripts/query-vibc-channel.sh http://127.0.0.1:9545 ../.devnet/polymer-l2-1.json channel-0`

