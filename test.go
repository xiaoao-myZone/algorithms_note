package main

import "fmt"

func main() {
	s := "{}[]()"
	for _, i := range s {
		fmt.Println(i)
	}
	fmt.Printf("%c\n", 124)
	fmt.Printf("%c\n", 92)
	fmt.Printf("%#v\n", s[1]-s[0])
	m := 10
	m -= (2 - 1) * (4 - 2)
	fmt.Println(m)

}
