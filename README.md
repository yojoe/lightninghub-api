# LightningHub API Specification

The _LightningHub API_ allows you to connect your Bitcoin/Lightning
[Bluewallet App](https://bluewallet.io) and serve multiple user accounts via a
shared Lightning Node.

This repository only contains the API specification in the RAML 1.0 format,
which can be used to generate various code stubs (client and server) and
human-readable documentation (e.g. HTML).

# Viewing the API documentation
Please see the static `index.html` in the `/docs` folder for a generated HTML
documentation (using [raml2html](https://github.com/raml2html/raml2html)).
The HTML documentation can be re-generated using the `raml2html.sh` file in the
`/scripts` folder.

# API Spec design & development
Please see [Atom](https://atom.io/) based http://apiworkbench.com/ for tools to
work with RAML 1.0 files. For other tools have a look at
https://raml.org/projects.

# Code generation
Client and server code can be automatically generated for various programming
languages and frameworks. Please see the `/docs` folder for details.
