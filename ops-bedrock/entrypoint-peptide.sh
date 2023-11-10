#!/bin/sh
set -exu

PEPTIDE_GENESIS_FILE='/root/.peptide/config/genesis.json'
L1_HASH=$(jq -r .genesis.l1.hash /rollup.json)

if [ ! -r "${PEPTIDE_GENESIS_FILE}" ]; then
	echo "${PEPTIDE_GENESIS_FILE} missing, running 'init' and 'seal'"
	echo "Initializing genesis."
    peptide init --l1-hash="${L1_HASH}" --l1-height=0
    echo "Sealing genesis block"
    peptide seal
else
	echo "Peptide already initialized."
fi
