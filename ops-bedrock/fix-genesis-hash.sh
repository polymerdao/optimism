#!/bin/sh

genesis_hash=$(jq -r .genesis_block.hash .devnet/config/genesis.json)
echo "Found Peptide genesis hash: ${genesis_hash}"

jq ".genesis.l2.hash = \"${genesis_hash}\" | .genesis.l2.number=1" .devnet/rollup.json > .devnet/peptide-rollup.json
echo "Updated rollup for peptide"
