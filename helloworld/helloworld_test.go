package helloworld

import "testing"

func TestGreetingMessage(t *testing.T) {
	result := GreetingMessage()
	if result != "Hello World" {
		t.Errorf("GreetingMessage() = %s; want Hello World", result)
	}
}
