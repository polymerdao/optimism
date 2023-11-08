#!/bin/sh

set -e -x

VERSION='v0.0.14'
CONTRACTS_DIR="./vibc-core-smart-contracts"
PRIVATE_KEY='bf7604d9d3a1c7748642b1b7b05c2bd219c9faa91458b370f85e5a40f3b03af7'

RPC_URL="${1}"     # http://127.0.0.1:9545
PORT_PREFIX="${2}" # polyibc.op1
ROLLUP_JSON="${3}" # rollup-1.json

# download contracts if needed
if [ ! -e "${CONTRACTS_DIR}" ]; then
    git clone --depth 1 --branch "${VERSION}" https://github.com/open-ibc/vibc-core-smart-contracts \
    && cd vibc-core-smart-contracts \
    && git submodule update --init --recursive \
    && npm install \
    && cd ..
fi

# deploy contracts on OP chain 1
cd vibc-core-smart-contracts
ESCROW_CONTRACT_ADDRESS=$(forge create --json --rpc-url "${RPC_URL}" --private-key "${PRIVATE_KEY}" contracts/Escrow.sol:Escrow | jq -r .deployedTo)
VERIFIER_CONTRACT_ADDRESS=$(forge create --json --rpc-url "${RPC_URL}" --private-key "${PRIVATE_KEY}" contracts/DummyVerifier.sol:DummyVerifier | jq -r .deployedTo)
OP_CONSENSUS_STATE_MANAGER_ADDRESS=$(forge create --json --rpc-url "${RPC_URL}" --private-key "${PRIVATE_KEY}" contracts/OpConsensusStateManager.sol:OptimisticConsensusStateManager --constructor-args "100" "${ESCROW_CONTRACT_ADDRESS}" | jq -r .deployedTo)
DISPATCHER_CONTRACT_ADDRESS=$(forge create --json --rpc-url "${RPC_URL}" --private-key "${PRIVATE_KEY}" contracts/Dispatcher.sol:Dispatcher --constructor-args "${VERIFIER_CONTRACT_ADDRESS}" "${ESCROW_CONTRACT_ADDRESS}" "${PORT_PREFIX}" "${OP_CONSENSUS_STATE_MANAGER_ADDRESS}" | jq -r .deployedTo)
cd ..

jq ".polymer_escrow_address=\"${ESCROW_CONTRACT_ADDRESS}\" |
    .polymer_verifier_address=\"${VERIFIER_CONTRACT_ADDRESS}\" |
    .op_consensus_state_manager_address=\"${OP_CONSENSUS_STATE_MANAGER_ADDRESS}\" |
    .polymer_dispatcher_address=\"${DISPATCHER_CONTRACT_ADDRESS}\"" ../.devnet/rollup.json | tee "${ROLLUP_JSON}"