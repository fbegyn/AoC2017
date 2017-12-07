package main

import (
	"os"
	"testing"
)

func BenchmarkDay06(b *testing.B) {
	os.Stdout, _ = os.Open(os.DevNull)
	for n := 0; n < b.N; n++ {
		main()
	}
}
