SRC := $(shell find .. -type f -name '*.go' -not -path '*/vendor/*')
bin/dhop: *.go vendor
	go build -o $@

vendor: Gopkg.*
	dep ensure

clean:
	rm -rf ./vendor ./bin
