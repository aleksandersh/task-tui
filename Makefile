SHELL = /bin/sh

binary = ./task-tui

.PHONY: build
build:
	go build

.PHONY: build_and_run_sample
build_and_run_sample: build
	$(binary) -t ./sample

$(binary):
	go build

run_sample: task-tui
	$(binary) -t ./sample
