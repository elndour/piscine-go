package main

import "fmt"

type donnie struct {
	Name     string
	Live     float32
	Age      int
	Aircraft int
}

func Pilot(ptr *donnie) {
	ptr.Name = "Donnie"
	ptr.Live = 100.0
	ptr.Age = 24
	ptr.Aircraft = 1
}

func main() {
	donnies := donnie{}
	Pilot(&donnies)
	fmt.Println(donnies)
}
