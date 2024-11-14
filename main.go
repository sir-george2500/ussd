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

   
   var sessionData SessionData
    
    // Load session data if exists
    if data, err := os.ReadFile(filePath); err == nil {
        json.Unmarshal(data, &sessionData)
        // Only increment the level when moving to a new menu option
        if text != "" { 
	// Don't increment on initial request
            sessionData.Level += 1
        }
    } else {
        sessionData.Level = 1
        fmt.Println("Session Data Not Found - Starting New Session at Level 1")
    }

    var response string
    fmt.Printf("Current Level: %d, Text Received: %s\n", sessionData.Level, text)

    // USSD menu logic
    switch {
    case (text == "" || text == "255"):
        sessionData.Level = 1  
	// Reset level for initial menu
        response = "Welcome to Winliberia Raffle Draw\n"
        response += "1. Winliberia Jackpot\n"
        response += "2. View Results\n"
        response += "3. Jackpot Policy\n"
        c.Header("FreeFlow", "FC")
        
    case text == "1" && sessionData.Level == 2:
        response = "Select Currency\n"
        response += "1. LRD\n"
        response += "2. USD\n"
        c.Header("FreeFlow", "FC")
        
    case text == "2" && sessionData.Level == 2:
        response = fmt.Sprintf("Your phone number is %s", phoneNumber)
        c.Header("FreeFlow", "FB")
        
    case text == "3" && sessionData.Level == 2:
        response = "Winliberia Policy"
        c.Header("FreeFlow", "FB")
        
    case (text == "1" || text == "2") && sessionData.Level == 3:
        currency := "LRD"
        if text == "2" {
            currency = "USD"
        }
        response = fmt.Sprintf("Processing payment in %s. You will receive a message shortly.", currency)
        c.Header("FreeFlow", "FB")
        
    default:
        response = "Invalid Input!!"
        c.Header("FreeFlow", "FB")
        sessionData.Level -= 1  // Revert level increment on invalid input
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
