package main

import (
	"./webservice"
	"fmt"
	"io/ioutil"
	"os"
	"log"
	"testing"
)

// zwraca tresc zapytania
func getMessage () string {
	return "hello didney warld";
}

// zwraca url servera
func getUrl () string {
	url := getFromFile()
	if url == "" {
		url = "http://127.0.0.1:8080"
	}
	return url
}
func getFromFile() string{

	pathtofile, _ := os.Getwd();
	pathtofile=pathtofile+"\\ClientConfig.txt"
	files , err :=ioutil.ReadFile(pathtofile);

	if err != nil {
		log.Fatal(err);
	}
	return string(files[:]);
}
func TestGettingFile(t *testing.T)

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