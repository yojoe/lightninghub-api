# DO NOT EDIT THIS FILE. This file will be overwritten when re-running go-raml.

from flask import Blueprint
import handlers


gettxs_api = Blueprint('gettxs_api', __name__)


@gettxs_api.route('/gettxs', methods=['GET'])
def gettxs_get():
    """
    Get a list of transactions (deposits and payments) along with their details
    for the authenticated user's custody account.
    It is handler for GET /gettxs
    """
    return handlers.gettxs_getHandler()
