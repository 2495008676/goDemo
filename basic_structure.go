package main

import (
	"fmt"
)

const (
	PI   = 3.14
	PII  = 4.14
	PIII = 5.14
)

var (
	a int    = 5
	b string = "string"
)

type struct_name struct{}

type interface_name interface{}

func main() {
	var c int
	var d bool
	var e string
	var f [1]int
	var g [1]bool
	h := 23
	fmt.Println(a, b, PI, PII, PIII, c, d, e, f, g, h)
}
