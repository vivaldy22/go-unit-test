package failunittest_test

import (
	"testing"

	. "github.com/vivaldy22/go-unit-test/2-fail-unit-test"
)

func TestHelloWorldFail(t *testing.T) {
	actual := HelloWorld()
	if actual != "Hello World" {
		t.Fail()
	}
}

func TestHelloWorldFailNow(t *testing.T) {
	actual := HelloWorld()
	if actual != "Hello World" {
		t.FailNow()
	}
}

func TestHelloWorldError(t *testing.T) {
	actual := HelloWorld()
	if actual != "Hello World" {
		t.Error("should be Hello World")
	}
}

func TestHelloWorldFatal(t *testing.T) {
	actual := HelloWorld()
	if actual != "Hello World" {
		t.Fatal("should be Hello World")
	}
}
