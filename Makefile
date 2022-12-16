
GOPATH:=$(shell go env GOPATH)

.PHONY: proto
proto:
	docker run --rm -v $(shell pwd):$(shell pwd) -w $(shell pwd) -e ICODE=15234383259D5605 cap1573/cap-protoc -I ./ --go_out=./ --micro_out=./ ./proto/cart/cart.proto
	
.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o cart *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t cart:latest
