package main

import "testing"

func TestCreateHelloResponseString (t *testing.T) {
	msg := "my random message"
	expectedMsg := "my random message echo"
	echoedMsg := CreateHelloResponseString(msg)
	if (echoedMsg != expectedMsg) {
		t.Error("got: " + msg, "expected: " +expectedMsg)
	}
}

func TestGetRequestContent (t *testing.T) {
	mockRawbody := []byte("<Envelope xmlns=\"http://schemas.xmlsoap.org/soap/envelope/\"><Body xmlns=\"http://schemas.xmlsoap.org/soap/envelope/\"><Content>hello didney warld</Content></Body></Envelope>")
	parsedRequest := GetRequestContent(mockRawbody)
	expectedParsedRequest := "hello didney warld"
	if (parsedRequest != expectedParsedRequest) {
		t.Error("got: " + parsedRequest, "expected: " +expectedParsedRequest)
	}
}

func TestGetRequestResponse (t *testing.T) {
	parsedRequest := "hello didney warld"
	expectedResponse := "<Envelope xmlns=\"http://schemas.xmlsoap.org/soap/envelope/\"><Body xmlns=\"http://schemas.xmlsoap.org/soap/envelope/\"><Content>hello didney warld echo</Content></Body></Envelope>"
	response := GetRequestResponse(parsedRequest)
	if (response != expectedResponse) {
		t.Error("got: " + response, "expected: " +expectedResponse)
	}
}

func TestVerifyRequest (t *testing.T) {
	testmsg := "hello didney warld"
	expected := true
	computed := VerifyRequest(testmsg)

	if (expected != computed) {
		t.Error("got: false", "expected: true")
	}

	testmsg = "9789789798"
	expected = false
	computed = VerifyRequest(testmsg)

	if (expected != computed) {
		t.Error("got: true", "expected: false")
	}
}