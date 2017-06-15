package main

import (
"./webservice"
"github.com/go-martini/martini"
	"net/http"
	"fmt"
	//"encoding/xml"
	"encoding/xml"
	"io/ioutil"
)

func createHelloResponseString (requestContent string) string {
	return requestContent + " echo"
}

func handleHello (r *http.Request) string {

	// pobierz strumien danych
	rawbody, err := ioutil.ReadAll(r.Body);

	// sprawdz, czy jest blad
	if (err != nil) {
		fmt.Println("error?");
	}

	// dekoduj sturmien
	requestEnvelope := &webservice.SOAPEnvelope{}
	xml.Unmarshal(rawbody, &requestEnvelope)

	requestContent := "hello didney warld response";

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