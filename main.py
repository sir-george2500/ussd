@app.route('/ussd_code/winliberia', methods=['POST', 'GET'])
def ussd():
      # Get parameters from the request
    msisdn = request.args.get('msisdndigits')
    service_code = request.args.get('shortcode')
    text = request.args.get('ussdstring', '')
    
    
    # Initialize response variables
    response_text = "CON Welcome to our service. Please select an option:\n1. Check balance\n2. View profile"

    # Handle USSD logic based on input
    if text == "1":
        response_text = "CON Your balance is $100."
    elif text == "2":
        response_text = "END Your profile details: Name: John Doe, Status: Active."
    else:
        response_text = "END Invalid option selected."

    # Return plain HTTP response as text
    return Response(response_text, content_type='text/plain')
