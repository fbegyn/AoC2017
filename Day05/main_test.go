package main

import (
	"testing"
)

func BenchmarkDay05(b *testing.B) {
	for n := 0; n < b.N; n++ {
		main()
	}
}
