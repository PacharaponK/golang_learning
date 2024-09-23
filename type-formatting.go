package main

import (
    "fmt"
)

func main() {
    var i = 15
    var txt = "Hello World!"
    var t, f = true, false
    var ft = 3.141

    // general formatting
    fmt.Printf("\n----- General -----")
    fmt.Printf("%v\n", i)
    fmt.Printf("%#v\n", i)
    fmt.Printf("%v%%\n", i)
    fmt.Printf("%T\n", i)
    fmt.Printf("%v\n", txt)
    fmt.Printf("%#v\n", txt)
    fmt.Printf("%T\n", txt)

    // integer formatting
    fmt.Printf("\n----- Int -----")
    fmt.Printf("%b\n", i)
    fmt.Printf("%d\n", i)
    fmt.Printf("%+d\n", i)
    fmt.Printf("%o\n", i)
    fmt.Printf("%O\n", i)
    fmt.Printf("%x\n", i)
    fmt.Printf("%X\n", i)
    fmt.Printf("%#x\n", i)
    fmt.Printf("%10d\n", i)
    fmt.Printf("%-10d\n", i)
    fmt.Printf("%010d\n", i)

    // float formatting
    fmt.Printf("\n----- Float -----")
    fmt.Printf("%e\n", ft)
    fmt.Printf("%f\n", ft)
    fmt.Printf("%.2f\n", ft)
    fmt.Printf("%10.2f\n", ft)
    fmt.Printf("%g\n", ft)

    // String Formatting
    fmt.Printf("\n----- String -----")
    fmt.Printf("%s\n", txt)
    fmt.Printf("%q\n", txt)
    fmt.Printf("%8s\n", txt)
    fmt.Printf("%-8s\n", txt)
    fmt.Printf("%x\n", txt)
    fmt.Printf("% x\n", txt)

    // boolean formatting
    fmt.Printf("\n----- Bool -----")
    fmt.Printf("%t\n", t)
    fmt.Printf("%t\n", f)
}
