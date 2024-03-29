.PHONY: generate
generate : 
	go generate ./entity

.PHONY: lint
lint:
	golangci-lint run

.PHONY: vendors
vendors:
	go mod download
	go mod tidy

.PHONY: build-app
build-app:
	docker build --tag messages .
