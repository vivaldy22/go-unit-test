package failunittest_test

import (
	"fmt"
	"testing"

	. "github.com/vivaldy22/go-unit-test/2-fail-unit-test"
)

func TestHelloWorldFail(t *testing.T) {
	actual := HelloWorld()
	if actual != "Hello World1" {
		t.Fail()
	}

	fmt.Println("ini akhir test")
}

func TestHelloWorldFailNow(t *testing.T) {
	actual := HelloWorld()
	if actual != "Hello World1" {
		t.FailNow()
	}
	fmt.Println("ini akhir test")
}

func TestHelloWorldError(t *testing.T) {
	actual := HelloWorld()
	if actual != "Hello World1" {
		t.Error("should be Hello World")
	}
	fmt.Println("ini akhir test")
}

func TestHelloWorldFatal(t *testing.T) {
	actual := HelloWorld()
	if actual != "Hello World1" {
		t.Fatal("should be Hello World")
	}
	fmt.Println("ini akhir test")
}
