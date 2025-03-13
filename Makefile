BIN_DIR=$(PWD)/bin
PICO_DIR=$(PWD)/cmd/pico
PICOCTL_DIR=$(PWD)/cmd/pico-ctl
CC=gcc
CXX=g++
GOFILES=`go list ./...`
GOFILESNOTEST=`go list ./... | grep -v test`
VERSION=$(shell git describe --abbrev=0 --tags 2>/dev/null || echo "v0.0.0")
BUILD=$(shell git rev-parse HEAD)
GOARCH=`go env GOARCH`
GOOS=`go env GOOS`
LDFLAGS=-ldflags="-s -w -X github.com/PicoTools/pico/internal/version.gitCommit=${BUILD} -X github.com/PicoTools/pico/internal/version.gitVersion=${VERSION}"
TAGS=sqlite_foreign_keys

run-local: build
	@bin/pico.${GOOS}.${GOARCH} --config config/config.yml -d run

build: proto-gen ent-gen go-lint
	@mkdir -p ${BIN_DIR}
	@echo "Building server ${GOOS}/${GOARCH} ${VERSION}"
	@GOOS=${GOOS} GOARCH=${GOARCH} CGO_ENABLED=0 CC=${CC} CXX=${CXX} go build -trimpath ${LDFLAGS} -tags="${TAGS}" -o ${BIN_DIR}/pico.${GOOS}.${GOARCH} ${PICO_DIR}

build-ctl: go-lint
	@mkdir -p ${BIN_DIR}
	@echo "Building operator management cli for ${GOOS}/${GOARCH} ${VERSION}"
	@GOOS=${GOOS} GOARCH=${GOARCH} CGO_ENABLED=0 CC=${CC} CXX=${CXX} go build -trimpath ${LDFLAGS} -o ${BIN_DIR}/pico-ctl.${GOOS}.${GOARCH} ${PICOCTL_DIR}

build-all: darwin-arm64 darwin-amd64 linux-arm64 linux-amd64 ctl-darwin-arm64 ctl-darwin-amd64 ctl-linux-arm64 ctl-linux-amd64

darwin-arm64: go-lint
	@mkdir -p ${BIN_DIR}
	@echo "Building server darwin/arm64 ${VERSION}"
	@GOOS=darwin GOARCH=arm64 CGO_ENABLED=0 CC=${CC} CXX=${CXX} go build -trimpath ${LDFLAGS} -tags="${TAGS}" -o ${BIN_DIR}/pico.darwin.arm64 ${PICO_DIR}
	@strip ${BIN_DIR}/pico.darwin.arm64 2>/dev/null || true

darwin-amd64: go-lint
	@mkdir -p ${BIN_DIR}
	@echo "Building server darwin/amd64 ${VERSION}"
	@GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 CC=${CC} CXX=${CXX} go build -trimpath ${LDFLAGS} -tags="${TAGS}" -o ${BIN_DIR}/pico.darwin.amd64 ${PICO_DIR}
	@strip ${BIN_DIR}/pico.darwin.amd64 2>/dev/null || true

linux-arm64: go-lint
	@mkdir -p ${BIN_DIR}
	@echo "Building server linux/arm64 ${VERSION}" 
	@GOOS=linux GOARCH=arm64 CGO_ENABLED=0 CC=${CC} CXX=${CXX} go build -trimpath ${LDFLAGS} -tags="${TAGS}" -o ${BIN_DIR}/pico.linux.arm64 ${PICO_DIR}
	@strip ${BIN_DIR}/pico.linux.arm64 2>/dev/null || true

linux-amd64: go-lint
	@mkdir -p ${BIN_DIR}
	@echo "Building server linux/amd64 ${VERSION}"
	@GOOS=linux GOARCH=amd64 CGO_ENABLED=0 CC=${CC} CXX=${CXX} go build -trimpath ${LDFLAGS} -tags="${TAGS}" -o ${BIN_DIR}/pico.linux.amd64 ${PICO_DIR}
	@strip ${BIN_DIR}/pico.linux.amd64 2>/dev/null || true

ctl-darwin-arm64: go-lint
	@mkdir -p ${BIN_DIR}
	@echo "Building management cli darwin/arm64 ${VERSION}"
	@GOOS=darwin GOARCH=arm64 CGO_ENABLED=0 CC=${CC} CXX=${CXX} go build -trimpath ${LDFLAGS} -o ${BIN_DIR}/pico-ctl.darwin.arm64 ${PICOCTL_DIR}
	@strip ${BIN_DIR}/pico-ctl.darwin.arm64 2>/dev/null || true

ctl-darwin-amd64: go-lint
	@mkdir -p ${BIN_DIR}
	@echo "Building management cli darwin/amd64 ${VERSION}"
	@GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 CC=${CC} CXX=${CXX} go build -trimpath ${LDFLAGS} -o ${BIN_DIR}/pico-ctl.darwin.amd64 ${PICOCTL_DIR}
	@strip ${BIN_DIR}/pico-ctl.darwin.amd64 2>/dev/null || true

ctl-linux-arm64: go-lint
	@mkdir -p ${BIN_DIR}
	@echo "Building management cli linux/arm64 ${VERSION}"
	@GOOS=linux GOARCH=arm64 CGO_ENABLED=0 CC=${CC} CXX=${CXX} go build -trimpath ${LDFLAGS} -o ${BIN_DIR}/pico-ctl.linux.arm64 ${PICOCTL_DIR}
	@strip ${BIN_DIR}/pico-ctl.linux.arm64 2>/dev/null || true

ctl-linux-amd64: go-lint
	@mkdir -p ${BIN_DIR}
	@echo "Building management cli linux/amd64 ${VERSION}"
	@GOOS=linux GOARCH=amd64 CGO_ENABLED=0 CC=${CC} CXX=${CXX} go build -trimpath ${LDFLAGS} -o ${BIN_DIR}/pico-ctl.linux.amd64 ${PICOCTL_DIR}
	@strip ${BIN_DIR}/pico-ctl.linux.amd64 2>/dev/null || true

go-lint:
	@echo "Linting Golang code"
	@go fmt ${GOFILES}
	@go vet ${GOFILESNOTEST}

ent-gen:
	@echo "Generating ent models"
	@go generate ./internal/ent

proto-gen:
	@echo "Generating proto schemes"
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

atlas-sqlite:
	@atlas schema inspect -u "ent://internal/ent/schema" --format '{{ sql . "  " }}' --dev-url "sqlite://file?mode=memory&_fk=1"

atlas-erd:
	@atlas schema inspect -u "ent://internal/ent/schema" --dev-url "sqlite://file?mode=memory&_fk=1" -w

clean:
	@rm -rf ${BIN_DIR}
