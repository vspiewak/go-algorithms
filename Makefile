.PHONY=test bench

all : test bench

test:
	go test ./... -v -count=1

bench:
	go test ./... -v -bench=. -benchmem -count=1 -run=^#
