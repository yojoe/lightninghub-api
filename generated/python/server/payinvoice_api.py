# DO NOT EDIT THIS FILE. This file will be overwritten when re-running go-raml.

from flask import Blueprint
import handlers


payinvoice_api = Blueprint('payinvoice_api', __name__)


@payinvoice_api.route('/payinvoice', methods=['POST'])
def payinvoice_post():
    """
    Request a BOLT-11 compatible Lightning invoice to be paid by the custodian.
    If the payment is successful, the respective amount + a fee is substracted
    from the authenticated user's balance.
    It is handler for POST /payinvoice
    """
    return handlers.payinvoice_postHandler()
