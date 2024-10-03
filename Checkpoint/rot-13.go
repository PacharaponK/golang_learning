package main

import (
	"bufio"
	"fmt"
	"os"
)

func rot13(text string) string {
	var encryptText string
	// fmt.Println(text)
	for _, char := range text {
		// fmt.Println(char)
		if char >= 'a' && char <= 'z' {
			encryptText += string((char-'a'+13)%26 + 'a')
		} else if char >= 'A' && char <= 'Z' {
			encryptText += string((char-'A'+13)%26 + 'A')
		} else {
			encryptText += string(char)
		}
	}
	return encryptText

}

func main() {
	var option int
	fmt.Println("Select 2 options")
	fmt.Println("- 1 encrypt with ROT-13")
	fmt.Println("- 2 decrypt with ROT-13")
	fmt.Println()
	fmt.Print("Choose option: ")
	fmt.Scan(&option)
	fmt.Scanln()
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	if option == 1 {
		fmt.Printf("Ciphertext is %s", rot13(text))
	} else {
		fmt.Printf("Plaintext is %s", rot13(text))
	}

}
