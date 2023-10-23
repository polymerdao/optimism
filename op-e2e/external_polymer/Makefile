build: shim polymer-peptide
.PHONY: build

POLYMER_CHAIN_DIR := ../../../polymerase/chain

polymer-peptide:
	@test -d $(POLYMER_CHAIN_DIR) || { echo "$(POLYMER_CHAIN_DIR) does not exist" ; exit 1; }
	cd $(POLYMER_CHAIN_DIR) && go build -o $(realpath .)/polymer-peptide ./cmd/peptide
.PHONY: polymer-peptide

shim: main.go
	go build -o shim .
.PHONY: shim

test-shim: build
	go test -v ./...
.PHONY: test-shim

clean:
	rm -f shim polymer-peptide
.PHONY: clean

.DEFAULT_GOAL := build
