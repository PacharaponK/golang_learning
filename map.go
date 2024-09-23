// // package main

// // import (
// //     "fmt"
// // )

// // func main() {
// //     var a = make(map[string]string) // แผนที่ a ตอนนี้ว่างเปล่า
// //     a["brand"] = "Ford"
// //     a["model"] = "Mustang"
// //     a["year"] = "1964" // ตอนนี้ a ไม่ว่างเปล่าแล้ว

// //     b := make(map[string]int)
// //     b["Oslo"] = 1
// //     b["Bergen"] = 2
// //     b["Trondheim"] = 3
// //     b["Stavanger"] = 4

// //     fmt.Printf("a\t%v\n", a)    // แสดงแผนที่ a
// //     fmt.Printf("b\t%v\n", b)    // แสดงแผนที่ b

// //     fmt.Println("data in b[\"Stavanger\"]", b["Stavanger"]) // แสดงค่าของ Stavanger
// // }


// package main

// import (
//     "fmt"
// )

// func main() {
//     var a = map[string]string{
//         "brand": "Ford",
//         "model": "Mustang",
//         "year":  "1964",
//         "day":   "",
//     }

//     val1, ok1 := a["brand"]  // ตรวจสอบคีย์ "brand"
//     val2, ok2 := a["color"]  // ตรวจสอบคีย์ "color" ที่ไม่มี
//     val3, ok3 := a["day"]    // ตรวจสอบคีย์ "day"
//     _, ok4 := a["model"]     // ตรวจสอบคีย์ "model"

//     fmt.Println(val1, ok1)  // แสดงค่าของ "brand" และสถานะการตรวจสอบ
//     fmt.Println(val2, ok2)  // แสดงค่าของ "color" และสถานะการตรวจสอบ
//     fmt.Println(val3, ok3)  // แสดงค่าของ "day" และสถานะการตรวจสอบ
//     fmt.Println(ok4)        // แสดงสถานะการตรวจสอบของ "model"
//     fmt.Println(a)          // แสดงแผนที่ a

//     delete(a, "year")      // ลบคีย์ "year" ออกจากแผนที่
//     fmt.Println(a)          // แสดงแผนที่ a หลังจากลบ
// }

package main

import (
    "fmt"
)

func main() {
    a := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}
    var b []string  // กำหนด slice ของ string
    b = append(b, "one", "two", "three", "four")  // เพิ่มสมาชิกในลำดับที่ต้องการ

    // วนลูปแสดงผลแผนที่ a โดยไม่มีลำดับที่กำหนด
    for k, v := range a {
        fmt.Printf("%v : %v, ", k, v)
    }
    fmt.Println()

    // วนลูปแสดงผลตามลำดับที่กำหนดใน slice b
    for _, element := range b {
        fmt.Printf("%v : %v, ", element, a[element])
    }
}
