# THIS FILE IS SAFE TO EDIT. It will not be overwritten when rerunning go-raml.

from flask import jsonify, request


def payinvoice_postHandler():

    inputs = request.get_json()

    return jsonify()
