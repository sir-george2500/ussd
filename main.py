@app.route('/ussd_code/winliberia', methods=['POST', 'GET'])
def ussd():
    msisdn = request.args.get('MSISDN')
    service_code = request.args.get('SERVICE_CODE')
    push_type = request.args.get('PUSH_TYPE')
    text = request.args.get('PUSH_TEXT', '')

    if text == "":
        response = "CON Welcome to our service:\n1. Check balance\n2. View profile\n3. Exit"

    elif text == "1":
        response = "END Your profile details: Name: John Doe, Status: Active."
    else:
        response = "END Invalid option selected."

    return response
