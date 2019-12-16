# LightningHub API Specification

The _LightningHub API_ allows you to connect your Bitcoin/Lightning
[Bluewallet App](https://bluewallet.io) and serve multiple user accounts via a
shared Lightning Node.

This repository only contains the API specification in the RAML 1.0 format,
which can be used to generate various code stubs (client and server) and
human-readable documentation (e.g. HTML).

# Viewing the API documentation
Please see the static [index.html](docs/index.html) in the [/docs](docs/) folder for a generated HTML
documentation (using [raml2html](https://github.com/raml2html/raml2html)).
The HTML documentation can be re-generated using the [raml2html.sh](scripts/raml2html.sh) file in the
[/scripts](/scripts) folder.

# API Spec design & development
Please see [Atom](https://atom.io/) based http://apiworkbench.com/ for tools to
work with RAML 1.0 files. For more tools have a look at
https://raml.org/projects.

Please have a look into the [raml10/api.raml](raml10/api.raml) file, which is 
the root document (starting point) of the LightningHub API specification.

# Code generation
Client and server code can be automatically generated for various programming
languages and frameworks. Please see the [/docs](/docs) folder for details.
