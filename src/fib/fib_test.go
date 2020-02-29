package main

import (
    "testing"
)

func BenchmarkFib(b *testing.B) {
    for i := 0; i < b.N; i++ {
        fib(8)
    }
}

func BenchmarkMemoFib(b *testing.B) {
    for i := 0; i < b.N; i++ {
        memoFib(8)
    }
}

