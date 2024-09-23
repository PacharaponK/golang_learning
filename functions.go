package main

import (
    "fmt"
    "math"
)

func hello_world() { // define function without parameters
    fmt.Println("Hello, World!")
}

func add(num1 int, num2 int) int { // function with two parameters
    return num1 + num2
}

func power(num int, expo int) (result float64) { // named return values
    result = math.Pow(float64(num), float64(expo))
    return
}

func factorial(n int) int { // recursive function
    if n == 1 {
        return 1
    }
    return n * factorial(n-1)
}

func main() {
    hello_world()
    fmt.Println(add(5, 6))
    fmt.Println(power(2, 2))
    fmt.Println("5!", factorial(5))
}
