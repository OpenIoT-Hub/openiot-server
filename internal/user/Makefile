.PHONY: generate
generate:
	kitex \
	-service "openiot-user" \
	-module "github.com/anxiu0101/openiot-hub" \
	-type protobuf \
	-I ./../../idl/ user.proto \
	&& go mod tidy

.PHONY: run_client
run_client:
	go run ./client/main.go

.PHONY: run_server
run_server:
	go run main.go handler.go

# Service: user
# Path: internal/user
# Optional: -use "github.com/anxiu0101/openiot-hub/kitex_gen"
