include .env
LOCAL_BIN:=$(CURDIR)/bin

.PHONY: help
help: ## List all available targets with help
	@grep -E '^[0-9a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) \
		| awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

.PHONY: install-golangci-lint
install-golangci-lint: ## install linter
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.54

.PHONY: install-deps
install-deps: ## install dependencies
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.14.0
	GOBIN=$(LOCAL_BIN) go install github.com/envoyproxy/protoc-gen-validate@v0.10.1
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.15.2
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.15.2
	GOBIN=$(LOCAL_BIN) go install github.com/rakyll/statik@v0.1.7
	GOBIN=$(LOCAL_BIN) go install go.uber.org/mock/mockgen@latest
.PHONY: get-deps
get-deps: ## get dependencies
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

.PHONY: generate
generate: ## generate handlers
	make generate-access-api

.PHONY: generate-access-api
generate-access-api:  ## generate handlers for /api/access
	mkdir -p pkg/access_v1
	protoc --proto_path api/access_v1 --proto_path vendor.protogen \
	--go_out=pkg/access_v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/access_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	--validate_out lang=go:pkg/access_v1 --validate_opt=paths=source_relative \
	--plugin=protoc-gen-validate=bin/protoc-gen-validate \
	--grpc-gateway_out=pkg/access_v1 --grpc-gateway_opt=paths=source_relative \
	--plugin=protoc-gen-grpc-gateway=bin/protoc-gen-grpc-gateway \
	api/access_v1/access.proto

.PHONY: local-migration-status
local-migration-status: ## show status migrations
	goose -dir ${MIGRATION_DIR} postgres ${GOOSE_DSN} status -v

.PHONY: local-migration-up
local-migration-up: ## apply migration
	goose -dir ${MIGRATION_DIR} postgres ${GOOSE_DSN} up -v

.PHONY: local-migration-down
local-migration-down: ## cancel migration
	goose -dir ${MIGRATION_DIR} postgres ${GOOSE_DSN} down -v

.PHONY: create-migration
create-migration: ## create new migration
	if [ -z "$(name)" ]; then \
			echo "Укажите название миграции в формате 'make create-migration name=create_auth'"; \
		else \
			goose -dir ./migrations create $(name) sql; \
		fi

.PHONY: up
up: ## up server
	docker-compose up -d

.PHONY: down
down: ## down server
	docker-compose down

vendor-proto:
		@if [ ! -d vendor.protogen/validate ]; then \
			mkdir -p vendor.protogen/validate &&\
			git clone https://github.com/envoyproxy/protoc-gen-validate vendor.protogen/protoc-gen-validate &&\
			mv vendor.protogen/protoc-gen-validate/validate/*.proto vendor.protogen/validate &&\
			rm -rf vendor.protogen/protoc-gen-validate ;\
		fi
		@if [ ! -d vendor.protogen/google ]; then \
			git clone https://github.com/googleapis/googleapis vendor.protogen/googleapis &&\
			mkdir -p  vendor.protogen/google/ &&\
			mv vendor.protogen/googleapis/google/api vendor.protogen/google &&\
			rm -rf vendor.protogen/googleapis ;\
		fi
		@if [ ! -d vendor.protogen/protoc-gen-openapiv2 ]; then \
			mkdir -p vendor.protogen/protoc-gen-openapiv2/options &&\
			git clone https://github.com/grpc-ecosystem/grpc-gateway vendor.protogen/openapiv2 &&\
			mv vendor.protogen/openapiv2/protoc-gen-openapiv2/options/*.proto vendor.protogen/protoc-gen-openapiv2/options &&\
			rm -rf vendor.protogen/openapiv2 ;\
		fi

.PHONY: mock
mock: ### Generate mocks
	go generate ./...

.PHONY: cover
cover: ### Get test covering
	go test -coverprofile=coverage.out -cover ./internal/service/... && go tool cover -html=coverage.out -o coverage.html

.PHONY: lint
lint: ## run linter
	GOBIN=$(LOCAL_BIN) golangci-lint run ./... --config .golangci.yaml

.PHONY: test
test: ## run tests
	go test ./... -race -count 100

.PHONY: integration-test
integration-test: ## run integration tests
	go test ./... -tags=integration

.PHONY: docker-build
docker-build: ## build for CI/CD
	docker buildx build --no-cache --platform linux/amd64 -t <REGISTRY>/access-server:v0.0.1 .
	docker login -u <REGISTRY_USERNAME> -p <REGISTRY_PASSWORD> <REGISTRY>
	docker push  <REGISTRY>/access-server:v0.0.1


.PHONY: build
build: ## Build the project using docker-compose
	docker login -u ${REGISTRY_USERNAME} -p ${REGISTRY_PASSWORD} ${REGISTRY}
	sh update_tag.sh
	docker-compose build

.PHONY: run
run: ## Run the project using docker-compose
	docker-compose up -d