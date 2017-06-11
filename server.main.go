package main

import (
	//"github.com/codegangsta/martini"
	//"github.com/martini-contrib/render"
	//"fmt"
	"os"
	//"github.com/fiorix/wsdl2go/soap"
	//"github.com/fiorix/wsdl2go/soap"
	"encoding/xml"
	"fmt"
	"./webservice"
)


func main() {
	// -- client
	fmt.Println("\n----client----\n")
	
	// create request object
	request := &webservice.Request{Content: "hello"}

	// parse request object to xml
	xmlRequest, err := xml.MarshalIndent(request, " ", "  ")

	if (err != nil) {
		fmt.Printf("error encoding request: %v\n", err)
	}

	// writeout xml as string
	fmt.Println("[client] sending:")
	os.Stdout.Write(xmlRequest)
	fmt.Println()

	// -- client: SEND "xmlRequest" ---


	fmt.Println("\n----server----\n")

	// -- server: RECEIVE "xmlRequest" ---

	// fake receive xml
	receivedXMLRequest := xmlRequest
	// create empty response object
	receivedRequest := &webservice.Request{Content: "?undefined?"}
	// fill empty object "input result" with received data
	err = xml.Unmarshal([]byte(receivedXMLRequest), &receivedRequest);

	if (err != nil) {
		fmt.Printf("error decoding request: %v\n", err)
	}

	fmt.Println("\n[server] received:")
	os.Stdout.Write(receivedXMLRequest)
	fmt.Println()

	fmt.Println("\n[server] decoded request 'content': " + receivedRequest.Content)

	response := &webservice.Response{Content: receivedRequest.Content + "echoed", OriginalContent: receivedRequest.Content}
	xmlResponse, err := xml.MarshalIndent(response, " ", "  ")

	if (err != nil) {
		fmt.Printf("error: %v\n", err)
	}

	// --- server: SEND "xmlResponse"

	fmt.Println("\n----client----\n")

	// --- client: RECEIVE "xmlServerResponse"

	// fake receive xml
	xmlReceivedResponse := xmlResponse
	// create receive object
	receivedResponse := &webservice.Response{Content: "?undefined?", OriginalContent: "?undefined?"}
	// fill empty object "input result" with received data
	err = xml.Unmarshal([]byte(xmlReceivedResponse), &receivedResponse);

	if (err != nil) {
		fmt.Printf("error decoding response: %v\n", err)
	}

	fmt.Println("\n[client] decoded request \n'content': " + receivedResponse.Content + "\noriginal-content: " + receivedResponse.OriginalContent)

}
