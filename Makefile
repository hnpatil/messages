.PHONY: generate
generate : 
	go generate ./entity

.PHONY: lint
lint:
	golangci-lint run