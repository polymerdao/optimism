Devnet
======

Requirements
------------
Install Foundry: <https://book.getfoundry.sh/getting-started/installation>

```sh
curl -L https://foundry.paradigm.xyz | bash
```

Init Contracts
--------------
```sh
make build-ts
```

Setup
-----
Start two OP stack chains and Polymer OP stack chain w/Peptide along with an op-relayer.

```sh
make devnet-up
```

Reset
-----
Removes currently running devnet containers, volumes, and initialization state.

```sh
scripts/clean-devnet.sh
```

IBC Setup
=========

Create IBC connections
----------------------
Create IBC connections for each OP stack chain.

```sh
./scripts/create-ibc-connection.sh localhost:4001 901 902
```

Create IBC channel
------------------
Establish an IBC channel connection the two OP stack chains via Polymer.
```sh
./scripts/create-ibc-channel.sh localhost:4001 ../.devnet/polymer-l2-1.json ../.devnet/polymer-l2-2.json
```


Send IBC Packet
---------------
Send a test IBC packet.

```sh
./scripts/send-ibc-packet.sh http://127.0.0.1:9545/ ../.devnet/polymer-l2-1.json channel-0 hello
```

Query VIBC Channel
------------------
Query IBC channel state.

```sh
./scripts/query-vibc-channel.sh http://127.0.0.1:9545 ../.devnet/polymer-l2-1.json channel-0
./scripts/query-vibc-channel.sh http://127.0.0.1:9546 ../.devnet/polymer-l2-2.json channel-1
```
