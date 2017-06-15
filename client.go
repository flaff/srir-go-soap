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
}