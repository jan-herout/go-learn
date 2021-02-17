package main

func main() {
	var b1,b2 byte = 0, 0xFF
	diff := b1 - 1
	if diff == b2 {
		panic("kaboom")
	}
}
