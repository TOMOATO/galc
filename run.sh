#!/bin/sh
GO_FILES=$(ls src | grep .go)
echo "run from $GO_FILES"
docker run --rm -v ${PWD}/src:/go/src/galc -w /go/src/galc golang:1.12.3-alpine go run ${GO_FILES} $@
