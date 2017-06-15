package echo

import "testing"

func TestEcho (t *testing.T) {
	msg := "my random message"
	expectedMsg := "my random message echo"
	echoedMsg := EchoMessage(msg)
	if (echoedMsg != expectedMsg) {
		t.Error("got: " + msg, "expected: " +expectedMsg)
	}
}