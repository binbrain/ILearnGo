package iteration

import (
    "testing"
    "fmt"
)

func TestIter(t *testing.T) {
    repeated := Repeat("a")
    expected := "aaaaa"

    if repeated != expected {
        t.Errorf("expected %q but got %q", expected, repeated)
    }
}

func TestYield(t *testing.T) {
    next := Yield(10)
    fmt.Println(next())
    fmt.Println(next())
    fmt.Println(next())
}

func BenchmarkRepeat(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Repeat("a")
    }
}
