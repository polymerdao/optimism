#!/usr/bin/env bash

# Helper script to deploy Polymer smart contracts and update the rollup.json as part of private testnet setup.

set -xeuo pipefail

PRIVATE_KEY='bf7604d9d3a1c7748642b1b7b05c2bd219c9faa91458b370f85e5a40f3b03af7'

RPC_URL="${1}"        # example: http://127.0.0.1:9545
PORT_PREFIX="${2}"    # example: polyibc.op1
POLYMER_JSON="${3}"   # example: polymer-l2-1.json
CONTRACTS_DIR="${4}"  # packages/contracts-bedrock/lib/vibc-core-smart-contracts

if [ ! -e ${POLYMER_JSON} ]; then
  # deploy contracts
  pushd ${CONTRACTS_DIR}
  ESCROW_CONTRACT_ADDRESS="$(
    forge create --json --rpc-url "${RPC_URL}" --private-key "${PRIVATE_KEY}" contracts/Escrow.sol:Escrow | \
      jq -r .deployedTo
  )"
  VERIFIER_CONTRACT_ADDRESS="$(
    forge create --json --rpc-url "${RPC_URL}" --private-key "${PRIVATE_KEY}" contracts/DummyVerifier.sol:DummyVerifier | \
      jq -r .deployedTo
  )"
  OP_CONSENSUS_STATE_MANAGER_ADDRESS="$(
    forge create --json --rpc-url "${RPC_URL}" --private-key "${PRIVATE_KEY}" \
      contracts/DummyConsensusStateManager.sol:DummyConsensusStateManager | \
    jq -r .deployedTo
  )"
  DISPATCHER_CONTRACT_ADDRESS="$(
    forge create --json --rpc-url "${RPC_URL}" --private-key "${PRIVATE_KEY}" \
      contracts/Dispatcher.sol:Dispatcher --constructor-args "${VERIFIER_CONTRACT_ADDRESS}" \
      "${ESCROW_CONTRACT_ADDRESS}" "${PORT_PREFIX}" "${OP_CONSENSUS_STATE_MANAGER_ADDRESS}" | \
    jq -r .deployedTo
  )"
  RECEIVER_CONTRACT_ADDRESS="$(
    forge create --json --rpc-url "${RPC_URL}" --private-key "${PRIVATE_KEY}" contracts/Mars.sol:Mars | \
      jq -r .deployedTo
  )"
  popd

  # create json to store addresses
  jq -n \
    --arg escrow "$ESCROW_CONTRACT_ADDRESS" \
    --arg verifier "$VERIFIER_CONTRACT_ADDRESS" \
    --arg op "$OP_CONSENSUS_STATE_MANAGER_ADDRESS" \
    --arg dispatcher "$DISPATCHER_CONTRACT_ADDRESS" \
    --arg receiver "$RECEIVER_CONTRACT_ADDRESS" \
  '{
    polymer_escrow_address: $escrow,
    polymer_verifier_address: $verifier,
    op_consensus_state_manager_address: $op,
    polymer_dispatcher_address: $dispatcher,
    polymer_receiver_address: $receiver
  }' | tee "${POLYMER_JSON}"
fi