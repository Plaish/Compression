package main

import (
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("comped.txt")
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return
	}

	decomped := Decompress(data)

	err = os.WriteFile("decomped.txt", decomped, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
}

func Decompress(input []byte) []byte {
	var decomped []byte

	for i := 0; i < len(input); i += 2 {
		count := int(input[i])
		char := input[i+1]

		for range count {
			decomped = append(decomped, char)
		}
	}

	return decomped
}
