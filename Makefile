BIN_DIR=$(PWD)/bin
PICO_DIR=$(PWD)/cmd/pico
CC=gcc
CXX=g++
VERSION=$(shell git describe --abbrev=0 --tags 2>/dev/null || echo "0.0.0")
BUILD=$(shell git rev-parse HEAD)
LDFLAGS=-ldflags="-s -w -X github.com/PicoTools/pico/internal/version.gitCommit=${BUILD} -X github.com/PicoTools/pico/internal/version.gitVersion=${VERSION}"
TAGS=sqlite_foreign_keys

.PHONY: run-local
run-local: pico
	@bin/pico --config config/config.yml -d run

.PHONY: pico
pico:
	@mkdir -p ${BIN_DIR}
	@echo "Building server..."
	CGO_ENABLED=0 CC=${CC} CXX=${CXX} go build ${LDFLAGS} -tags="${TAGS}" -o ${BIN_DIR}/pico ${PICO_DIR}
	@strip bin/pico

.PHONY: pico-race
pico-race:
	@mkdir -p ${BIN_DIR}
	@echo "Building race server..."
	CC=${CC} CXX=${CXX} go build -race ${LDFLAGS} -o ${BIN_DIR}/pico.race ${PICO_DIR}

.PHONY: dep-shared
dep-shared:
	@echo "Update shared components..."
	@export GOPRIVATE="github.com/PicoTools" && go get -u github.com/PicoTools/pico-shared && go mod tidy && go mod vendor

.PHONY: ent-gen
ent-gen:
	@echo "Generating ent models..."
	@go generate ./internal/ent

.PHONY: atlas-sqlite
atlas-sqlite:
	@atlas schema inspect -u "ent://internal/ent/schema" --format '{{ sql . "  " }}' --dev-url "sqlite://file?mode=memory&_fk=1"

.PHONY: atlas-erd
atlas-erd:
	@atlas schema inspect -u "ent://internal/ent/schema" --dev-url "sqlite://file?mode=memory&_fk=1" -w

.PHONY: clean
clean:
	@rm -rf ${BINARY_DIR}
