SERVICE_NAME ?= store-service
DOCKER_TAG ?= latest
VERSION ?= $(GIT_VERSION)
PWD := $(shell pwd)

GIT_VERSION := $(shell git describe --tags)

all: deps check build test image

## Static Analysis ##
deps:
	go get -v github.com/golang/lint/golint
	go get github.com/constabulary/gb/...

check: deps
	go vet ./src/...
	golint ./src/...

## Compile & Build ##
build: check
	gb build -ldflags "-X frontend.Version=$(VERSION)"

## Pre-Deployment Testing ##
test: build
	make -C test

## Docker Image Build ##
image: test
	docker build --pull -t yinc2/$(SERVICE_NAME):$(DOCKER_TAG) -f Dockerfile .

## Docker Image Push ##
push: image
	docker push yinc2/$(SERVICE_NAME):$(DOCKER_TAG)
	docker tag yinc2/$(SERVICE_NAME):$(DOCKER_TAG) yinc2/$(SERVICE_NAME):$(VERSION)
	docker push yinc2/$(SERVICE_NAME):$(VERSION)

docker-build: docker-pull-golang
	@rm -rf bin/ tmp/
	@mkdir -p bin tmp/
	docker run --rm \
		-v "$$PWD/src":/$(SERVICE_NAME)/src \
		-v "$$PWD/vendor":/$(SERVICE_NAME)/vendor \
		-v "$$PWD/Makefile":/$(SERVICE_NAME)/Makefile \
		-v "$$PWD/tmp":/$(SERVICE_NAME)/bin \
		-w /$(SERVICE_NAME) \
		$(DOCKER_ENV) \
		golang:1.7 \
		make "VERSION=$(VERSION)" deps install install-race test check
	@cp tmp/$(SERVICE_NAME) tmp/$(SERVICE_NAME)-race bin/
	@rm -rf tmp/

docker-run:
	docker run -it -P --rm --name $(SERVICE_NAME) \
			               docker-registry.core.rcsops.com/$(SERVICE_NAME)

clean:
	rm -rf bin/ pkg/ test/*.pyc
	make -C test clean

.PHONY: deps install install-race test check docker-build docker-image docker-push docker-run \
		autotag version clean
