package main

import "github.com/01-edu/z01"

type Door struct {
	state bool
}

const (
	OPEN  = true
	CLOSE = false
)

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func OpenDoor(ptrDoor *Door) {
	PrintStr("Door Opening...")
	(*ptrDoor).state = OPEN
}

func CloseDoor(ptrDoor *Door) {
	PrintStr("Door Closing...")
	(*ptrDoor).state = CLOSE
}

func IsDoorOpen(s *Door) bool {
	PrintStr("is the Door opened ?")
	return (*s).state
}

func IsDoorClose(s *Door) bool {
	PrintStr("is the Door closed ?")
	return !(*s).state
}

func main() {
	door := &Door{}

	OpenDoor(door)
	if IsDoorClose(door) {
		OpenDoor(door)
	}
	if IsDoorOpen(door) {
		CloseDoor(door)
	}
}
