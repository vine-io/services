GIT_COMMIT=$(shell git rev-parse --short HEAD)
GIT_TAG=$(shell git describe --abbrev=0 --tags --always --match "v*")
GIT_VERSION=github.com/vine-io/services/pkg/runtime/doc
CGO_ENABLED=0
BUILD_DATE=$(shell date +%s)
LDFLAGS=-X $(GIT_VERSION).GitCommit=$(GIT_COMMIT) -X $(GIT_VERSION).GitTag=$(GIT_TAG) -X $(GIT_VERSION).BuildDate=$(BUILD_DATE)

build-tag:
	sed -i "" "s/GitTag     = ".*"/GitTag     = \"$(GIT_TAG)\"/g" pkg/runtime/doc.go
	sed -i "" "s/GitCommit  = ".*"/GitCommit  = \"$(GIT_COMMIT)\"/g" pkg/runtime/doc.go
	sed -i "" "s/BuildDate  = ".*"/BuildDate  = \"$(BUILD_DATE)\"/g" pkg/runtime/doc.go

proto: proto-apiserver proto-auth

proto-apiserver:
	cd $(GOPATH)/src && \
	protoc -I=$(GOPATH)/src --gogo_out=:. --vine_out=:. --validator_out=:. github.com/vine-io/services/api/service/apiserver/v1/apiserver.proto

proto-auth:
	cd $(GOPATH)/src && \
	protoc -I=$(GOPATH)/src --gogo_out=:. --vine_out=:. --validator_out=:. github.com/vine-io/services/api/service/auth/v1/auth.proto

install:

build:

clean:

.PHONY: build-tag install build clean
