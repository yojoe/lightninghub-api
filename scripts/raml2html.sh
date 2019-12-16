#!/usr/bin/env bash

# requires raml2html in $PATH (see https://github.com/raml2html/raml2html)
raml2html ../raml10/api.raml > ../docs/index.html
