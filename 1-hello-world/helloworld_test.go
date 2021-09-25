package helloworld

import "testing"

func TestHelloWorld(t *testing.T) {
	hasil := HelloWorld()
	if hasil != "Hello Enigma" {
		panic("hasil tidak mengeluarkan Hello World")
	}
}
