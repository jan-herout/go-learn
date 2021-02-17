package main

import "fmt"

func main() {	
	// "byte" string - poslední dva bajty jsou diakritika ve Win1250
	s := string([]byte{'a', 'b', 0x0A, 'c', 0xA7, 0xD8})
	
	// dej mi třetí a čtvrtý *bajt* (!!)
	fmt.Printf("bytes, hex: % x\n", s[3:5])
	
	// range přes index - BYTE
	for i:=0; i<len(s); i++ {
		fmt.Printf("iterace %d, hex: %02x\n", i, s[i])
	}
	
	// range + value - RUNE
	for _, c := range s {
		fmt.Printf("rune=%q typ=%T hodnota=%v\n", c, c, c)
	}		
	_ = s[0:4]
}
