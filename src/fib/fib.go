package main

import (
    "fmt"
)

type fibf func(int) int

func memoFib(num int) int {
    cache := make(map[int]int)
    var fn fibf

    fn = func(num int) int {
        if val, ok := cache[num]; ok {
            return val
        }
        if num <= 1 {
            return num
        }
        cache[num] = fn(num - 1) + fn(num - 2)
        return cache[num]
    }
    return fn(num)
}

func fib(num int) int {
    if num <= 1 {
        return num
    }

    return fib(num - 1) + fib(num - 2)
}

func main() {
    fmt.Println(memoFib(89))
    fmt.Println(fib(89))
}
