SHELL := /bin/bash
BIN = bin

.PHONY: build
build: $(BIN)/
	go build -o $(BIN)/solar-controller .


.PHONY: run_example
run_example: build
	$(BIN)/solar-controller server -c examples/config.yaml

$(BIN)/:
	mkdir -p $@