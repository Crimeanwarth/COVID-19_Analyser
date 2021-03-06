package API

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// This Package is made for API request and JSON parsing.
// We want to get more functions in order to make it easy to use for future work.

// *************** Function Declarations *******************

func HelloFunc() { // Hello initialisation function
	fmt.Println("Hello API!")
}

func Availability(url string) (bool, *http.Response) { // Returns true if API is
	response, err := http.Get(url)

	if err == nil && response != nil {
		return true, response
	} else {
		log.Fatal(err)
		return false, response
	}
}

func ReadBody(response *http.Response) []byte { // Reads and gives you body as string
	data, err := ioutil.ReadAll(response.Body) // gets the body from the API
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func URLGenerator(address string, CountryCode string) string {
	newAddress := address + CountryCode
	return newAddress
}
