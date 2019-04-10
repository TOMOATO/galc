#!/bin/sh
docker run --rm -v ${PWD}/src:/go/src/galc -w /go/src/galc golang:1.12.3-alpine go test --cover -v \
  | sed ''/PASS/s//$(printf "\033[32mPASS\033[0m")/'' \
  | sed ''/FAIL/s//$(printf "\033[31mFAIL\033[0m")/'' \
  | sed ''/error/s//$(printf "\033[31merror\033[0m")/''
