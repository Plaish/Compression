package main

import (
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("data.txt")
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return
	}

	comped := Compress(data)

	err = os.WriteFile("comped.txt", comped, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
}

func Compress(input []byte) []byte {
	var comped []byte
	count := 1

	for i := range input {
		if i == len(input)-1 || input[i] != input[i+1] {
			comped = append(comped, byte(count))
			comped = append(comped, input[i])
			count = 1
		} else {
			count++
		}
	}

	return comped
}
