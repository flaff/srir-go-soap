package main

import (
	"./webservice"
	"fmt"
	"io/ioutil"
	"os"
	"log"

)

// zwraca tresc zapytania
func GetMessage () string {
	return "hello didney warld";
}

// zwraca url servera
func GetUrl () string {
	url := GetFromFile("ClientConfig.txt")
	if url == "" {
		url = "http://127.0.0.1:8080"
	}
	return url
}
func GetFromFile(filename string) string{

	pathtofile, _ := os.Getwd();
	pathtofile=pathtofile+"\\"+filename
	files , err :=ioutil.ReadFile(pathtofile);

	if err != nil {
		log.Fatal(err);
	}
	return string(files[:]);
}

func GetExpectedResponse (request string) string {
	return request + " echo"
}

func VerifyResponse (request string, response string) bool {
	return GetExpectedResponse(request) == response
}

func main () {
	// ustaw informacje o web service
	url := GetFromFile("configfile.txt");
	service := webservice.NewHello_PortType(url, false, nil)

	// zrob request
	message := GetMessage()
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
	if (VerifyResponse(message, response)) {
		fmt.Println("correct response")
	} else {
		fmt.Println("incorrect response, expected: " + GetExpectedResponse(message))
	}


}