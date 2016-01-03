package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// Make a Client and new HTTP request
	client := http.Client{}
	req, err := http.NewRequest("GET", "https://upload.wikimedia.org/wikipedia/commons/1/18/Bradypus.jpg", nil)
	if err != nil {
		fmt.Println("Error creating HTTP request:", err)
		return
	}

	// Send the request and get back the HTTP response
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}

	// Print out the contents of the response body, which is
	// the encoding of a picture of a three-toed sloth.
	responseBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error getting response body bytes:", err)
	}
	fmt.Println(string(responseBytes))
}
