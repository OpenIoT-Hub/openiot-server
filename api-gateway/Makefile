.PHONY: update
update:
	hz update -I idl -idl idl/device/device.proto
	hz update -I idl -idl idl/user/user.proto

.PHONY: build
build:
	go build -o hertz_test

.PHONY: run
run:
	go run ./main.go ./router_gen.go
