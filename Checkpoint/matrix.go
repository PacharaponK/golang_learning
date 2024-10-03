package main

// import (
// 	"fmt"
// )

// func inputMatrix(M *[3][2]int) {
// 	var num int
// 	for i := 0; i < 3; i++ {
// 		for j := 0; j < 2; j++ {
// 			fmt.Printf("M[%d][%d]: ", i+1, j+1)
// 			fmt.Scan(&num)
// 			M[i][j] = num
// 		}
// 	}
// }

// func min(M [3][2]int) int {
// 	var minimun int = M[0][0]
// 	for i := 0; i < 3; i++ {
// 		for j := 0; j < 2; j++ {
// 			if M[i][j] < minimun {
// 				minimun = M[i][j]
// 			}
// 		}
// 	}
// 	return minimun
// }

// func max(M [3][2]int) int {
// 	var maximum int = M[0][0]
// 	for i := 0; i < 3; i++ {
// 		for j := 0; j < 2; j++ {
// 			if M[i][j] > maximum {
// 				maximum = M[i][j]
// 			}
// 		}
// 	}
// 	return maximum
// }

// func main() {
// 	M := [3][2]int{}
// 	inputMatrix(&M)
// 	fmt.Println()
// 	fmt.Println("Matrix")
// 	for i := 0; i < 3; i++ {
// 		fmt.Printf("%d\t", M[i][0])
// 		fmt.Printf("%d\t", M[i][1])
// 		fmt.Println()
// 	}
// 	fmt.Println()
// 	fmt.Printf("Min = %d\n", min(M))
// 	fmt.Printf("Max = %d\n", max(M))
// }
