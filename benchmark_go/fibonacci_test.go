package main

import "testing"

func BenchmarkFibonacciIterative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacciIterative(20) // Benchmarking Fibonacci of 20
	}
}

// Benchmark for Fibonacci using recursive approach
func BenchmarkFibonacciRecursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacciRecursive(20) // Benchmarking Fibonacci of 20
	}
}

func BenchmarkStringConcatenate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringConcatenate(100)
	}
}
