.PHONY: generate
generate:
	kitex \
	-service "openiot-user" \
	-module "github.com/OpenIoT-Hub/openiot-server" \
	-type protobuf \
	--protobuf Muser.proto= github.com/OpenIoT-Hub/openiot-server/kitex_gen/openiot/user \
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
# Optional: -use "github.com/OpenIoT-Hub/openiot-server/kitex_gen"
# kitex -service "openiot-user" -module "github.com/OpenIoT-Hub/openiot-server" -use "github.com/OpenIoT-Hub/openiot-server/kitex_gen" -type protobuf -I ./../../idl/ server/user.proto
