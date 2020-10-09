.PHONY: all

all: build run

build:
	@go build -o bs

run:
	@echo "run..." && ./bs