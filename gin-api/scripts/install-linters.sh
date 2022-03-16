#!/bin/sh

cd /tmp
go install golang.org/x/tools/cmd/goimports@latest
go install github.com/kisielk/errcheck@latest
go install golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow@latest
cd -