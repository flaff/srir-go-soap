package main

import (
"./webservice"
"github.com/go-martini/martini"
	"net/http"
	"fmt"
	//"encoding/xml"
	"encoding/xml"
	"io/ioutil"
	"log"
)

func createHelloResponseString (requestContent string) string {
	return requestContent + " echo"
}

func handleHello (r *http.Request) string {

	// pobierz strumien danych
	rawbody, err := ioutil.ReadAll(r.Body)

	if (len(rawbody) == 0) {
		fmt.Println("empty response")
		return ""
	}

	// sprawdz, czy jest blad
	if (err != nil) {
		fmt.Println("error?");
	}

	requestContent := ""

	// dekoduj strumien
	requestEnvelope := new(webservice.SOAPEnvelope)
	requestEnvelope.Body = webservice.SOAPBody{Content: &requestContent}
	err = xml.Unmarshal(rawbody, requestEnvelope)

	log.Println("Received response:");
	log.Println(string(rawbody))

	// stworz odpowiedz
	responseEnvelope := &webservice.SOAPEnvelope{}
	responseEnvelope.Body.Content = createHelloResponseString(requestContent);

	response, err := xml.MarshalIndent(responseEnvelope, "", "");

	return string(response);
}

func main () {
	m := martini.Classic()

	m.Post("/", handleHello)

	m.RunOnAddr(":8080")

}