#!/usr/bin/env bash

# tear down the running devnet

set -euo pipefail

docker kill $(docker ps -a --filter name=^/ops-bedrock -q)
docker volume rm $(docker volume ls --filter='name=ops-bedrock*' -q)
rm -r .devnet
