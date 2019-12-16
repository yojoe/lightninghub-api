#!/usr/bin/env bash

# requires go-raml in $PATH (see https://github.com/Jumpscale/go-raml)
mkdir -p ../generated/go/server 2>/dev/null
go-raml -d server -l go --dir ../generated/go/server \
               --ramlfile ../raml10/api.raml \
               --import-path github.com/yojoe/lightninghub-api/go/server
