@app.route('/ussd_code/winliberia', methods=['POST', 'GET'])
def ussd():
    msisdn = request.args.get('MSISDN')
    service_code = request.args.get('SERVICE_CODE')
    push_type = request.args.get('PUSH_TYPE')
    text = request.args.get('PUSH_MENU', '')

    if text == "":
        response = "FC Welcome to Winliberia:\n"
        response += "1) Winliberia Jackpot\n"
        response += "2) View results\n"
        response += "3) Jackpot Policy\n"
        return response
    
    elif text == "1":
        response = "FB Your profile details: Name: John Doe, Status: Active."
        return response
    else:
        response = "FB Invalid option selected."
        return response



