package main

import (
	"testing"
)

func BenchmarkDay08(b *testing.B) {
	for n := 0; n < b.N; n++ {
		main()
	}
}
