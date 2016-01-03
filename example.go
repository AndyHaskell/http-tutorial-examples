package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// Make a Client and new HTTP request
	client := http.Client{}
	req, err := http.NewRequest("GET", "http://example.com", nil)
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

	// Print out the response status code and Content-Type
	fmt.Println("Response status was:", res.Status)
	fmt.Println(res.Header.Get("Content-Type"))
	fmt.Println()

	// Print out the contents of the response body, which is
	// the HTML of example.com
	responseBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error getting response body bytes:", err)
	}
	fmt.Println(string(responseBytes))
}
