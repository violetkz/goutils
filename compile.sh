#!/bin/env bash

export GOPATH=`pwd`
export GOBIN=$GOPATH/bin

(cd src/config/ && go install )
(cd src/test/   && go install && cp -f config.txt $GOBIN)
