BIN_DIR=$(PWD)/bin
PICO_DIR=$(PWD)/cmd/pico
CC=gcc
CXX=g++
GOFILES=`go list ./...`
GOFILESNOTEST=`go list ./... | grep -v test`
VERSION=$(shell git describe --abbrev=0 --tags 2>/dev/null || echo "v0.0.0")
BUILD=$(shell git rev-parse HEAD)
LDFLAGS=-ldflags="-s -w -X github.com/PicoTools/pico/internal/version.gitCommit=${BUILD} -X github.com/PicoTools/pico/internal/version.gitVersion=${VERSION}"
TAGS=sqlite_foreign_keys

run-local: darwin-arm64
	@bin/pico.darwin.arm64 --config config/config.yml -d run

build-all: darwin-arm64 darwin-amd64 linux-arm64 linux-amd64

darwin-arm64: proto-gen ent-gen go-lint
	@mkdir -p ${BIN_DIR}
	@echo "Building server darwin/arm64 ${VERSION}"
	@GOOS=darwin GOARCH=arm64 CGO_ENABLED=0 CC=${CC} CXX=${CXX} go build -trimpath ${LDFLAGS} -tags="${TAGS}" -o ${BIN_DIR}/pico.darwin.arm64 ${PICO_DIR}

darwin-amd64: proto-gen ent-gen go-lint
	@mkdir -p ${BIN_DIR}
	@echo "Building server darwin/amd64 ${VERSION}"
	@GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 CC=${CC} CXX=${CXX} go build -trimpath ${LDFLAGS} -tags="${TAGS}" -o ${BIN_DIR}/pico.darwin.amd64 ${PICO_DIR}

linux-arm64: proto-gen ent-gen go-lint
	@mkdir -p ${BIN_DIR}
	@echo "Building server linux/arm64 ${VERSION}" 
	@GOOS=linux GOARCH=arm64 CGO_ENABLED=0 CC=${CC} CXX=${CXX} go build -trimpath ${LDFLAGS} -tags="${TAGS}" -o ${BIN_DIR}/pico.linux.arm64 ${PICO_DIR}

linux-amd64: proto-gen ent-gen go-lint
	@mkdir -p ${BIN_DIR}
	@echo "Building server linux/amd64 ${VERSION}"
	@GOOS=linux GOARCH=amd64 CGO_ENABLED=0 CC=${CC} CXX=${CXX} go build -trimpath ${LDFLAGS} -tags="${TAGS}" -o ${BIN_DIR}/pico.linux.amd64 ${PICO_DIR}

darwin-arm64-race: proto-gen ent-gen go-lint
	@mkdir -p ${BIN_DIR}
	@echo "Building race server darwin/arm64 ${VERSION}"
	@GOOS=darwin GOARCH=arm64 CC=${CC} CXX=${CXX} go build -race ${LDFLAGS} -o ${BIN_DIR}/pico.darwin.arm64.race ${PICO_DIR}

go-lint:
	@echo "Linting Golang code..."
	@go fmt ${GOFILES}
	@go vet ${GOFILESNOTEST}

ent-gen:
	@echo "Generating ent models..."
	@go generate ./internal/ent

proto-gen:
	@echo "Generating proto schemes..."
	@protoc --proto_path=$(PWD)/proto \
		--go-grpc_out=paths=source_relative:$(PWD)/pkg/proto \
		--go_out=paths=source_relative:$(PWD)/pkg/proto \
		--go_opt=Mcommon/v1/common.proto=github.com/PicoTools/pico/pkg/proto/common/v1 \
		--go-grpc_opt=Mcommon/v1/common.proto=github.com/PicoTools/pico/pkg/proto/common/v1 \
		--go_opt=Mlistener/v1/listener.proto=github.com/PicoTools/pico/pkg/proto/listener/v1 \
		--go-grpc_opt=Mlistener/v1/listener.proto=github.com/PicoTools/pico/pkg/proto/listener/v1 \
		--go_opt=Moperator/v1/operator.proto=github.com/PicoTools/pico/pkg/proto/operator/v1 \
		--go-grpc_opt=Moperator/v1/operator.proto=github.com/PicoTools/pico/pkg/proto/operator/v1 \
		--go_opt=Mmanagement/v1/management.proto=github.com/PicoTools/pico/pkg/proto/management/v1 \
		--go-grpc_opt=Mmanagement/v1/management.proto=github.com/PicoTools/pico/pkg/proto/management/v1 \
		common/v1/common.proto \
		listener/v1/listener.proto \
		operator/v1/operator.proto \
		management/v1/management.proto

go-sync:
	@go mod tidy && go mod vendor

atlas-sqlite:
	@atlas schema inspect -u "ent://internal/ent/schema" --format '{{ sql . "  " }}' --dev-url "sqlite://file?mode=memory&_fk=1"

atlas-erd:
	@atlas schema inspect -u "ent://internal/ent/schema" --dev-url "sqlite://file?mode=memory&_fk=1" -w

clean:
	@rm -rf ${BINARY_DIR}
