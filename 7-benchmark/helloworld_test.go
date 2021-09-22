package tabletest_test

import (
	"testing"

	. "github.com/vivaldy22/go-unit-test/7-benchmark"
)

func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Hello("Benchmark")
	}
}

func BenchmarkHelloSub(b *testing.B) {
	b.Run("Sub 1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Hello("Sub 1")
		}
	})
	b.Run("Sub 2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Hello("Sub 2")
		}
	})
}

func BenchmarkHelloTable(b *testing.B) {
	benchmarks := []struct {
		name  string
		param string
	}{
		{
			name:  "Hello(Enigma)",
			param: "Enigma",
		},
		{
			name:  "Hello(Camp)",
			param: "Camp",
		},
	}

	for _, bb := range benchmarks {
		b.Run(bb.name, func(b *testing.B) {
			Hello(bb.param)
		})
	}
}
