.PHONY: wireformat

bin_path=./bin

wireformat:
	go build -o $(bin_path)/wireformat cmd/wireformat/wireformat.go

http-resolver:
	go build -o $(bin_path)/http-resolver cmd/dns/http-resolver.go

fmt:
	go fmt ./...

