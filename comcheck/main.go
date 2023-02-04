package main

import (
	"fmt"
	"os"
)

func main() {
	A := os.Args[1:]
	if len(A) == 0 {
		return
	}
	if len(A) != 0 {
		for i := 0; i < len(A); i++ {
			if A[i] == "01" || A[i] == "galaxy" || A[i] == "galaxy 01" {
				fmt.Println("Alert!!!")
				break
			}
		}
	}
}
