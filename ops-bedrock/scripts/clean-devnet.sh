#!/usr/bin/env bash

# tear down the running devnet

set -uo pipefail

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

docker kill $(docker ps -a --filter name=^/ops-bedrock -q)
docker volume rm $(docker volume ls --filter='name=ops-bedrock*' -q)
rm -rf ${SCRIPT_DIR}/../../.devnet
