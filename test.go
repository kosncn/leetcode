package main

import "fmt"

func main() {
	s := '1'
	r := s - '0'
	rs := int(r)
	fmt.Printf("%T: %d\n", r, r)
	fmt.Printf("%T: %d\n", rs, rs)
}
