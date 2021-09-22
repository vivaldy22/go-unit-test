package helloworld_test

import (
	"testing"

	. "github.com/vivaldy22/go-unit-test/1-hello-world"
)

func TestHelloWorld(t *testing.T) {
	actual := HelloWorld()
	if actual != "Hello World" {
		panic("Actual result is not Hello World")
	}
}

