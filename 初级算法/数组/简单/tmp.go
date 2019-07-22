package main

import "fmt"

func main() {
	a := []map[string]string{{
		"a": "A",
	}, {
		"b": "B",
	}}
	b := 0
	fmt.Println(a)
	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T", b, b)
}
