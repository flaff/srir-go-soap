package main

import "testing"

func TestReadingFromURL(t *testing.T) {
	testUrl :=GetUrl()
	expectedUrl :="http://127.0.0.1:8080"
	if testUrl !=expectedUrl{
		t.Errorf("Expected http should be http://127.0.0.1:8080 but it was %s ", testUrl)
	}
}
func TestMissingURLFile(t *testing.T){
	testUrl :=GetFromFile("testfile.txt")
	expectedUrl:=""
	if testUrl !=expectedUrl {
		t.Errorf("Expected http should be null but it was %s ", testUrl)
	}
}
func TestVerifyResponse(t *testing.T){
	test := VerifyResponse("testmessage","testmessage echo")
	if (!test) {
		t.Error("Response wasnt the same like request" )
	}
}