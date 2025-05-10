SHELL = /bin/sh

binary = ./task-tui

.PHONY: build
build:
	go build

.PHONY: build_and_run_example
build_and_run_example: build
	$(binary) -t ./examples

$(binary):
	go build

run_example: task-tui
	$(binary) -t ./examples

.PHONY: goreleaser_build
goreleaser_build:
	goreleaser build --clean

.PHONY: goreleaser_release
goreleaser_release:
	goreleaser release --clean
