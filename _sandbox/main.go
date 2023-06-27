package main

import "fmt"

func main() {
	fmt.Println(1i, 3.2)
	fmt.Printf("%T %T\n", 1i, 3.2)

	type MyInt int
	xs := []MyInt{1, 2, 3}
	fmt.Println(xs)
}
