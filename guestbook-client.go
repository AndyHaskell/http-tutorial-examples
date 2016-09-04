package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Expected 2 command line arguments")
		return
	}

	// Get the message we want to post from the command-line
	// argument passed in and add it to the form
	msg := os.Args[1]
	formValues := make(url.Values)
	formValues.Add("message", msg)

	formValuesStr := formValues.Encode()

	// Convert our url.Values to an HTTP request body. Then use
	// http.NewRequest to put that request body on a POST request.
	requestBody := strings.NewReader(formValuesStr)
	req, err := http.NewRequest("POST",
		"http://localhost:1123/guestbook", requestBody)
	if err != nil {
		fmt.Println("Error making HTTP request: ", err)
		return
	}

	// Add a Content-Type header to the request to signify that
	// the format of the request is URL-encoded form data
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Send the request with client.Do, just like we would for
	// a GET request
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending HTTP request: ", err)
		return
	}

	if res.StatusCode == http.StatusOK {
		fmt.Println("Message successfully sent")
	} else {
		fmt.Println("Error sending POST request; HTTP status was", res.Status)
	}
}
