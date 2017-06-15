package main

import (
	"./webservice"
	"./echo"
	"./message"
	"fmt"
	"io/ioutil"
	"os"
	"log"
	"bufio"
)

// zwraca tresc zapytania
func GetMessage () string {
	return message.GetMessage();
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
	return echo.EchoMessage(request)
}

func VerifyResponse (request string, response string) bool {
	return GetExpectedResponse(request) == response
}

func _main () {
	// ustaw informacje o web service
	url := GetFromFile("ClientConfig.txt");
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

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter to continue")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
}