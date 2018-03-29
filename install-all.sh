#!/bin/bash

# install ll
go build -o $GOPATH/bin/ll ll/ll.go

# ishadowx
go build -o $GOPATH/bin/ishadowx ishadowx/ishadowx.go ishadowx/flag.go ishadowx/crawler.go ishadowx/type.go
