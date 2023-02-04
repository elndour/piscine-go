package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 2 {
		file, err := os.Open(os.Args[1])
		if err != nil {
			fmt.Printf("Errors : %v", err.Error())
		}

		tab := make([]byte, 14)
		file.Read(tab)
		fmt.Println(string(tab))
		file.Close()
	} else if len(os.Args) < 2 {
		fmt.Println("File name missing")
	} else {
		fmt.Println("Too many arguments")
	}
}
