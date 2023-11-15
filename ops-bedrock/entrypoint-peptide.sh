#!/bin/sh
set -exu

PEPTIDE_GENESIS_JSON='/root/.peptide/config/genesis.json'
L1_HASH=$(jq -r .genesis.l1.hash /devnet/rollup.json)

if [ ! -r "${PEPTIDE_GENESIS_JSON}" ]; then
	echo "${PEPTIDE_GENESIS_JSON} missing, running 'init' and 'seal'"
	echo "Initializing genesis."
    peptide init --l1-hash="${L1_HASH}" --l1-height=0
    echo "Sealing genesis block"
    peptide seal

    cp ${PEPTIDE_GENESIS_JSON} /devnet/genesis-peptide.json

    genesis_hash=$(jq -r .genesis_block.hash ${PEPTIDE_GENESIS_JSON})
    genesis_block=$(jq -r .genesis_block.number ${PEPTIDE_GENESIS_JSON})

    echo "Found Peptide genesis hash: ${genesis_hash} for block: ${genesis_block}"

    jq \
        --arg genesis_hash "${genesis_hash}" \
        --arg genesis_block "${genesis_block}" \
        '.genesis.l2.hash = $genesis_hash | .genesis.l2.number=($genesis_block|tonumber)' /devnet/rollup.json > /devnet/peptide-rollup.json

    echo "Peptide genesis initialized."
else
	echo "Peptide already initialized."
fi
