package main

import "fmt"

func main() {
	var name string
	fmt.Println("Name:")
	fmt.Scanln(&name)
	fmt.Println(name)
}
