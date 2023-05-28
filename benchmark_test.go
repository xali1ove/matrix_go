package main

import "testing"

func BenchmarkFindLargestGroup(b *testing.B) {
	height, width, numColors := 1000, 1000, 10
	matrix := generateMatrix(height, width, numColors)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		findLargestGroup(matrix)
	}
}
