.PHONY: generate
generate:
	kitex \
	-service "openiot-device" \
	-module "github.com/OpenIoT-Hub/openiot-server" \
	-type protobuf \
	--protobuf Muser.proto= github.com/OpenIoT-Hub/openiot-server/kitex_gen/openiot/device \
	-I ./../../idl/ device.proto\
	&& go mod tidy

.PHONY: run_client
run_client:
	go run ./client/main.go

.PHONY: run_server
run_server:
	go run main.go handler.go

# Service: user
# Path: internal/user
# Optional: -use "github.com/OpenIoT-Hub/openiot-server/kitex_gen"
