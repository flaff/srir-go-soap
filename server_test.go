package main

import "testing"

func CreateHelloResponseStringTest (t *testing.T) {
	msg := "my random message"
	expectedMsg := "my random message echo"
	echoedMsg := CreateHelloResponseString(msg)
	if (echoedMsg != expectedMsg) {
		t.Error("got: " + msg, "expected: " +expectedMsg)
	}
}

func GetRequestContentTest (t *testing.T) {
	mockRawbody := []byte("<Envelope xmlns=\"http://schemas.xmlsoap.org/soap/envelope/\"><Body xmlns=\"http://schemas.xmlsoap.org/soap/envelope/\"><Content>hello didney warld</Content></Body></Envelope>")
	parsedRequest := GetRequestContent(mockRawbody)
	expectedParsedRequest := "hello didney warld"
	if (parsedRequest != expectedParsedRequest) {
		t.Error("got: " + parsedRequest, "expected: " +expectedParsedRequest)
	}
}

func GetRequestResponseTest (t *testing.T) {
	parsedRequest := "hello didney warld"
	expectedResponse := "<Envelope xmlns=\"http://schemas.xmlsoap.org/soap/envelope/\"><Body xmlns=\"http://schemas.xmlsoap.org/soap/envelope/\"><Content>hello didney warld echo</Content></Body></Envelope>"
	response := GetRequestResponse(parsedRequest)
	if (response != expectedResponse) {
		t.Error("got: " + response, "expected: " +expectedResponse)
	}
}