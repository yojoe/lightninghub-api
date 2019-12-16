#!/usr/bin/env bash
cd `dirname $0`
# requires go-raml in $PATH (see https://github.com/Jumpscale/go-raml)
mkdir -p ../generated/go/server 2>/dev/null
go-raml -d server -l python \
             --dir ../generated/python/server \
             --ramlfile ../raml10/api.raml \
             --no-apidocs
