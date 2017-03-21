SERVICE_NAME ?= store-service
DOCKER_TAG ?= latest
VERSION ?= $(GIT_VERSION)
PWD := $(shell pwd)

GIT_VERSION := $(shell git describe --tags)

all: deps check build test image

# Get Dependencies #
deps:
	go get -v github.com/golang/lint/golint
	go get github.com/constabulary/gb/...

## 1. Static Check ##
check: deps
	go vet ./src/...
	golint ./src/...

## 2. Compile & Build ##
build: check
	gb build -ldflags "-X frontend.Version=$(VERSION)"

## 3. Pre-Deployment Testing ##
unit:
	gb test model -v
functional: build
	make -C test
test: unit functional

## 4. Image Build ##
image: test
	docker build --pull -t yinc2/$(SERVICE_NAME):$(DOCKER_TAG) -f Dockerfile .

## 5. Image Push ##
push: image
	docker push yinc2/$(SERVICE_NAME):$(DOCKER_TAG)
	docker tag yinc2/$(SERVICE_NAME):$(DOCKER_TAG) yinc2/$(SERVICE_NAME):$(VERSION)
	docker push yinc2/$(SERVICE_NAME):$(VERSION)

# Docker build (Optional, this can help make the build server clean)
docker-pull-golang:
	docker pull golang:1.7
docker-build: docker-pull-golang
	@rm -rf bin/ tmp/
	@mkdir -p bin tmp/
	docker run --rm \
		-v "$$PWD/src":/$(SERVICE_NAME)/src \
		-v "$$PWD/test":/$(SERVICE_NAME)/test \
		-v "$$PWD/vendor":/$(SERVICE_NAME)/vendor \
		-v "$$PWD/Makefile":/$(SERVICE_NAME)/Makefile \
		-v "$$PWD/tmp":/$(SERVICE_NAME)/bin \
		-w /$(SERVICE_NAME) \
		$(DOCKER_ENV) \
		golang:1.7 \
		make "VERSION=$(VERSION)" build
	@cp tmp/$(SERVICE_NAME) bin/
	@rm -rf tmp/

# Docker run the image
docker-run:
	docker run -it -P --rm --name $(SERVICE_NAME) \
			               docker-registry.core.rcsops.com/$(SERVICE_NAME)

clean:
	rm -rf bin/ pkg/ test/*.pyc
	make -C test clean

.PHONY: deps install install-race test check docker-build docker-image docker-push docker-run \
		autotag version clean
