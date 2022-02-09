package main

import "fmt"

func main() {
	fmt.Println("Chris")

	m1 := make(map[int]string)
	m1[1] = "kyriakh"
	fmt.Println(m1)

	for _, value := range m1 {
		fmt.Println(value)
	}

}
