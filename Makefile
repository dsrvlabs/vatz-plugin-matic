build:
	@go build

run:
	@go run main.go

reflect:
	@grpcurl -plaintext localhost:9091 list pilot.plugin.ManagerPlugin
