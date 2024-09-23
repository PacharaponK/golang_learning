package main

import ("fmt")

func main() {
    var i, j float32
	var area float32
    fmt.Print("Base: ")
    fmt.Scan(&i)
	fmt.Print("Height: ")
	fmt.Scan(&j)
	area = (i * j) / 2
    fmt.Printf("Area: %.4f", area)
}
