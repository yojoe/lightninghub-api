# DO NOT EDIT THIS FILE. This file will be overwritten when re-running go-raml.

from flask import Blueprint
import handlers


create_api = Blueprint('create_api', __name__)


@create_api.route('/create', methods=['POST'])
def create_post():
    """
    Send account creation request
    It is handler for POST /create
    """
    return handlers.create_postHandler()
