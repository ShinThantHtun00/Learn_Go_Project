package main

import(
	"fmt"
	"example.com/greetings"
)

func main()  {
	message := greetings.Hello("Hallo")
	fmt.Println(message)
}