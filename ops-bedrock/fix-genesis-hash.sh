#!/bin/sh

genesis_hash=$(docker logs ops-bedrock-op-peptide-1 | grep -o -E 'hash=(.*)' | cut -d = -f 2)
echo "Found Peptide genesis hash: ${genesis_hash}"

jq ".genesis.l2.hash = \"${genesis_hash}\"" .devnet/rollup.json > .devnet/peptide-rollup.json
echo "Updated rollup for peptide"
