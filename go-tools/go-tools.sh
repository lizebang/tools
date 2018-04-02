#!/bin/bash

cd $GOPATH/bin
rm gocode gopkgs go-outline go-symbols guru gorename godef godoc goreturns golint dlv

go get -u -v github.com/nsf/gocode
go get -u -v github.com/uudashr/gopkgs/cmd/gopkgs
go get -u -v github.com/ramya-rao-a/go-outline
go get -u -v github.com/acroca/go-symbols
go get -u -v golang.org/x/tools/cmd/guru
go get -u -v golang.org/x/tools/cmd/gorename
go get -u -v github.com/rogpeppe/godef
go get -u -v golang.org/x/tools/cmd/godoc
go get -u -v sourcegraph.com/sqs/goreturns
go get -u -v github.com/golang/lint/golint
go get -u -v github.com/derekparker/delve/cmd/dlv
