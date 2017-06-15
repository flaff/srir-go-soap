package main

import (
	"./webservice"
	"./echo"
	"./message"
	"github.com/go-martini/martini"
	"net/http"
	"fmt"
	"encoding/xml"
	"io/ioutil"
	"log"
)

func CreateHelloResponseString(requestContent string) string {
	return echo.EchoMessage(requestContent)
}

func VerifyRequest (request string) bool {
	return message.GetMessage() == request
}

// zwraca tresc zapytania
func GetRequestContent (rawbody []byte) string {
	requestContent := ""

	// dekoduj strumien
	requestEnvelope := new(webservice.SOAPEnvelope)
	requestEnvelope.Body = webservice.SOAPBody{Content: &requestContent}
	xml.Unmarshal(rawbody, requestEnvelope)

	log.Println("Received response:");
	log.Println(string(rawbody))

	return requestContent
}

// stwarza odpowiedz na podstawie zapytania
func GetRequestResponse (requestContent string) string {
	// stworz odpowiedz
	responseEnvelope := &webservice.SOAPEnvelope{}
	responseEnvelope.Body.Content = CreateHelloResponseString(requestContent);
	response, err := xml.MarshalIndent(responseEnvelope, "", "");

	if (err != nil) {
		fmt.Println(err)
	}

	return string(response);
}

func handleHello(r *http.Request) string {
	// pobierz strumien danych
	rawbody, err := ioutil.ReadAll(r.Body)

	// sprawdz, czy jest blad
	if err != nil {
		fmt.Println("error?");
	}
	if len(rawbody) == 0 {
		fmt.Println("empty request")
	}

	// dekoduj zapytanie
	requestContent := GetRequestContent(rawbody)

	if (VerifyRequest(requestContent)) {
		fmt.Println("correct message from client: " + requestContent)
	} else {
		fmt.Println("INCORRECT message from client: " + requestContent)
	}

	// zwroc odpowiedz
	return GetRequestResponse(requestContent)
}

func main() {
	m := martini.Classic()

	m.Post("/", handleHello)

	m.RunOnAddr(":8080")
}
