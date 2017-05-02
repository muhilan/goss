# !/bin/bash

	docker run --rm -v $PWD:/usr/src/goss -w /usr/src/goss golang:1.8.1 bash -c "go get -d -insecure;CGO_ENABLED=0 go build -v -o goss"