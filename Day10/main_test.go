package main

import (
	"testing"
)

func BenchmarkDay10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		main()
	}
}
