.PHONY: wireformat

bin_path=./bin

wireformat:
	go build -o $(bin_path)/wireformat cmd/wireformat/wireformat.go

fmt:
	go fmt ./...

