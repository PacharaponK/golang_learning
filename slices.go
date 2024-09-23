// package main

// import (
//     "fmt"
// )

// func main() {
//     myslice1 := []int{}  // Empty slice of integers
//     fmt.Println(myslice1)         // []
//     fmt.Println(len(myslice1))    // Length: 0
//     fmt.Println(cap(myslice1))    // Capacity: 0

//     myslice2 := []string{"Go", "Slices", "Are", "Powerful"}  // Slice of strings
//     fmt.Println(myslice2)         // [Go Slices Are Powerful]
//     fmt.Println(len(myslice2))    // Length: 4
//     fmt.Println(cap(myslice2))    // Capacity: 4
// }


// package main

// import (
//     "fmt"
// )

// func main() {
//     arr1 := [6]int{10, 11, 12, 13, 14, 15}  // อาร์เรย์ขนาด 6
//     myslice := arr1[2:4]  // สร้าง slice จาก index 2 ถึง index 3 (ไม่รวม index 4)

//     // แสดงผล slice ที่ได้
//     fmt.Printf("myslice = %v\n", myslice)       // ค่าของ myslice คือ [12 13]
//     fmt.Printf("length = %d\n", len(myslice))   // ความยาวของ slice คือ 2 (มี 2 ตัว)
//     fmt.Printf("capacity = %d\n", cap(myslice)) // ความจุของ slice คือ 4 (จาก index 2 ถึง index สุดท้ายของอาร์เรย์)
// }


// package main

// import (
//     "fmt"
// )

// func main() {
//     myslice1 := make([]int, 5, 10)  // สร้าง slice ที่มี length 5 และ capacity 10
//     fmt.Printf("myslice1 = %v\n", myslice1)
//     fmt.Printf("length = %d\n", len(myslice1))   // Length = 5
//     fmt.Printf("capacity = %d\n", cap(myslice1)) // Capacity = 10

//     // สร้าง slice ที่มี length 5 โดยที่ capacity จะถูกกำหนดให้เท่ากับ length
//     myslice2 := make([]int, 5)
//     fmt.Printf("myslice2 = %v\n", myslice2)
//     fmt.Printf("length = %d\n", len(myslice2))   // Length = 5
//     fmt.Printf("capacity = %d\n", cap(myslice2)) // Capacity = 5
// }

package main

import (
    "fmt"
)

func main() {
    prices := []int{10, 20, 30}  // สร้าง slice ของ int ที่มีสมาชิก 10, 20, 30
    fmt.Println(prices[0])       // พิมพ์ค่าใน index 0 (คือ 10)
    fmt.Println(prices[2])       // พิมพ์ค่าใน index 2 (คือ 30)
}