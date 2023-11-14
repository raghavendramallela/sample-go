package helloworld

import "testing"

func TestGreetsGitHub(t *testing.T) {
	result := Greet()
	if result != "Hello World" {
		t.Errorf("Greet() = %s; want Hello World", result)
	}
}
