package main

import "fmt"

func main() {
	// bool
	// string
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint 64 uintptr
	// byte (alia for uint8)
	// rune (alias for int32, represents a unicode code point)
	// float32 float64
	// complex64 complex128
	SuccessGuy := "Shin Thant Htun" //string
	i := 42                         //int
	f := 3.14                       //float
	g := 0.867 + 0.5i               //complex
	var isIt = true                 //bool
	fmt.Print(i)
	fmt.Print(f)
	fmt.Print(g)
	fmt.Print(isIt)
	fmt.Printf("The future most successful guy in the world is %v", SuccessGuy)

}
