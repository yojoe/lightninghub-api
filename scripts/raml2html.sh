#!/usr/bin/env bash

# requires raml2html in path https://github.com/raml2html/raml2html
raml2html ../raml10/api.raml > ../generated/html/index.html
