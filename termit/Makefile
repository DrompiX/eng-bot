LOCAL_BIN:=$(CURDIR)/bin

.PHONY: generate-mocks
generate-mocks:
	$(LOCAL_BIN)/mockgen --version || GOBIN=$(LOCAL_BIN) go install github.com/golang/mock/mockgen@v1.6.0
	GOBIN=$(LOCAL_BIN) go generate -run "mockgen" ./...

.PHONY: test
test:
	go test -v ./...

.PHONY: run
run:
	go run cmd/termit.go

.PHONY: gen-proto
gen-proto:
	mkdir -p pb && protoc \
	--proto_path=proto proto/*.proto \
	--go_out=:pb --go-grpc_out=:pb \
	--grpc-gateway_out=:pb \
	--grpc-gateway_opt=logtostderr=true \
    --grpc-gateway_opt=paths=source_relative \
    --grpc-gateway_opt=generate_unbound_methods=true \
	--openapiv2_out=:swagger