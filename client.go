package main

import (
	"./webservice"
	"fmt"
	"io/ioutil"
)

// zwraca tresc zapytania
func getMessage () string {
	return "hello didney warld";
}

// zwraca url servera
func getUrl () string {
	return "http://127.0.0.1:8080";
}
func getFromFile() string{

	files , _ :=ioutil.ReadFile("e:\\configfile.txt",)

	return string(files[:]);
}

func getExpectedResponse (request string) string {
	return request + " echo"
}

func verifyResponse (request string, response string) bool {
	return getExpectedResponse(request) == response
}

func main () {
	// ustaw informacje o web service
	url := getFromFile();			//getUrl()
	service := webservice.NewHello_PortType(url, false, nil)

	// zrob request
	message := getMessage()
	responsePtr, err := service.SayHello(&message)

	// zweryfikuj odpowiedz
	response := *responsePtr

	// obsluga bledu
	if (err != nil) {
		fmt.Println("error occured during request");
		fmt.Println(err);
	} else {
		fmt.Println("response from server is: " + response);
	}

	// weryfikacja odpowiedzi
	if (verifyResponse(message, response)) {
		fmt.Println("correct response")
	} else {
		fmt.Println("incorrect response, expected: " + getExpectedResponse(message))
	}


}