package message

import "testing"

func TestGetMessage (t *testing.T) {
	msg := GetMessage()
	expectedMsg := "hello didney warld"
	if (msg != expectedMsg) {
		t.Error("got: " + msg, "expected: " +expectedMsg)
	}
}