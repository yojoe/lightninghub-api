# DO NOT EDIT THIS FILE. This file will be overwritten when re-running go-raml.

from flask import Flask, send_from_directory, send_file
from addinvoice_api import addinvoice_api
from auth_api import auth_api
from balance_api import balance_api
from create_api import create_api
from decodeinvoice_api import decodeinvoice_api
from getbtc_api import getbtc_api
from getinfo_api import getinfo_api
from getpending_api import getpending_api
from gettxs_api import gettxs_api
from getuserinvoices_api import getuserinvoices_api
from payinvoice_api import payinvoice_api

import os
dir_path = os.path.dirname(os.path.realpath(__file__))

app = Flask(__name__)

app.register_blueprint(addinvoice_api)
app.register_blueprint(auth_api)
app.register_blueprint(balance_api)
app.register_blueprint(create_api)
app.register_blueprint(decodeinvoice_api)
app.register_blueprint(getbtc_api)
app.register_blueprint(getinfo_api)
app.register_blueprint(getpending_api)
app.register_blueprint(gettxs_api)
app.register_blueprint(getuserinvoices_api)
app.register_blueprint(payinvoice_api)


@app.route('/', methods=['GET'])
def home():
    return send_file(dir_path + '/index.html')


if __name__ == "__main__":
    app.run(debug=True)