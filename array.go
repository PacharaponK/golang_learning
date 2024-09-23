package main

import (
    "fmt"
)

func main() {
    var arr1 = [3]int{1, 2, 3}
    arr2 := [5]int{4, 5, 6, 7, 8}      // fully initialized
    arr3 := [5]int{}                    // not initialized, all elements default to 0
    arr4 := [5]int{11, 12}              // partially initialized, remaining elements default to 0
    arr5 := [5]int{1: 10, 2: 40}        // initialize only specific elements
    arr6 := [...]int{1, 2, 3, 4, 5, 6}  // length is inferred

    fmt.Println(arr1)   // [1, 2, 3]
    fmt.Println(arr2)   // [4, 5, 6, 7, 8]
    fmt.Println(arr3)   // [0, 0, 0, 0, 0]
    fmt.Println(arr4)   // [11, 12, 0, 0, 0]
    fmt.Println(arr5)   // [0, 10, 40, 0, 0]
    fmt.Println(arr6)   // [1, 2, 3, 4, 5, 6]
    fmt.Println(arr1[1]) // 2

    arr2[2] = 20        // Modify the element at index 2 of arr2
    fmt.Println(arr2)   // [4, 5, 20, 7, 8]
    fmt.Println(len(arr2)) // Length of arr2, which is 5
}
