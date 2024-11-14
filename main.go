type SessionData struct {
	Level int `json:"level"`
}

func UssdHandler(c *gin.Context) {
	// Set Cache-Control headers to prevent caching
	
    var sessionId, phoneNumber, serviceCode, text string

    // Get parameters based on request method
    if c.Request.Method == http.MethodPost {
        sessionId = c.PostForm("sessionId")
        serviceCode = c.PostForm("serviceCode")
        phoneNumber = c.PostForm("msisdndigits")
        text = c.PostForm("ussdstring")
    } else if c.Request.Method == http.MethodGet {
        sessionId = c.Query("dialogueID")
        serviceCode = c.Query("serviceCode")
        phoneNumber = c.Query("msisdndigits")
        text = c.Query("ussdstring")
    }

    fmt.Printf("Received Request - Method: %s, sessionId: %s, serviceCode: %s, phoneNumber: %s, text: %s\n",
        c.Request.Method, sessionId, serviceCode, phoneNumber, text)

    // Specify file path and create directory if not exists
    dirPath := "sess"
    if _, err := os.Stat(dirPath); os.IsNotExist(err) {
        os.Mkdir(dirPath, 0755)
        fmt.Println("Created directory:", dirPath)
    } else {
        fmt.Println("Directory exists:", dirPath)
    }

    filePath := filepath.Join(dirPath, fmt.Sprintf("%s.json", sessionId))
    var sessionData SessionData

    // Load session data if exists
    if data, err := os.ReadFile(filePath); err == nil {
        json.Unmarshal(data, &sessionData)
		sessionData.Level += 1
    } else {
        sessionData.Level = 1
        fmt.Println("Session Data Not Found - Starting New Session at Level 1")
    }

    var response string
    fmt.Printf("Current Level: %d, Text Received: %s\n", sessionData.Level, text)

    // USSD menu logic
    if (text == "" || text == "255") && sessionData.Level == 1 {
        response = "Welcome to Winliberia Raffle Draw\n"
        response += "1. Winliberia Jackpot\n"
        response += "2. View Results\n"
        response += "3. Jackpot Policy\n"
        c.Header("FreeFlow", "FC")

    } else if text == "1" && sessionData.Level == 2 {
        response = "Select Currency\n"
        response += "1. LRD\n"
        response += "2. USD\n"
        c.Header("FreeFlow", "FC")

    } else if text == "2" && sessionData.Level == 2 {
        response = fmt.Sprintf("Your phone number is %s", phoneNumber)
        fmt.Println("Displaying phone number:", phoneNumber)
        c.Header("FreeFlow", "FB")

    } else if text == "3" && sessionData.Level == 2 {
        response = "Winliberia Policy"
        c.Header("FreeFlow", "FB")
        fmt.Println("Displaying Jackpot Policy")

    } else if text == "1" && sessionData.Level == 3 {
        response = "Processing payment in LRD. You will receive a message shortly."
        c.Header("FreeFlow", "FB")

    } else if text == "2" && sessionData.Level == 3 {
        response = "Processing payment in USD. You will receive a message shortly."
        c.Header("FreeFlow", "FB")

    } else {
        response = "Invalid Input!!"
        c.Header("FreeFlow", "FB")
        fmt.Println("Invalid input received.")
    }

    // Save updated level to file
	sessionDataBytes, err := json.Marshal(sessionData)
	if err == nil {
		err = os.WriteFile(filePath, sessionDataBytes, 0644)
		if err != nil {
			fmt.Printf("Error Saving Session Data: %v\n", err)
		} else {
			fmt.Println("Session Data Saved Successfully")
		}
	} else {
		fmt.Printf("Error Marshalling Session Data: %v\n", err)
	}

    c.String(http.StatusOK, response)
}
