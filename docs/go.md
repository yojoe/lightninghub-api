# Generating code stubs for Go
This approach uses https://github.com/Jumpscale/go-raml to generate code stubs
from the RAML 1.0 file that allow for **round trip engineering**.

# Prerequisites
- [Go](https://golang.org/doc/install) >= v1.13
- [go-raml](https://github.com/Jumpscale/go-raml)

# Generate Go server stubs
See the ready-to-use `raml2go-server.sh` script in the `/scripts` folder, which is basically just a single command:
```
go-raml -d server -l go \
          --dir ../generated/go/server \
          --ramlfile ../raml10/api.raml \
          --import-path github.com/yojoe/lightninghub-api/go/server \
          --no-apidocs
```

`go-raml` will generate a lot of files which can be divided into two categories:
- files which **will be overwritten** when go-raml is re-run
- files which **won't be overwritten** when go-raml is re-run

You're custom implementation of the LightningHub API should be done in those
files which start with the following comment line:
```
// THIS FILE IS SAFE TO EDIT. It will not be overwritten when rerunning go-raml.
```

If you follow these recommendations, round-trip engineering should work fine.

# Generate Go client stubs
Please see `go-raml help client` for details. It should be almost identical to generating the server code above.
