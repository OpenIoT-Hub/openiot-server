.PHONY: update
update:
	hz update -proto_path ../idl/api

.PHONY: build
build:
	go build -o hertz_test

.PHONY: run
run:
	go run ./main.go ./router_gen.go
