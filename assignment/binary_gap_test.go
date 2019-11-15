package main

import "testing"

func BenchmarkBinaryGap(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		BinaryGap(10000001)
	}
}
