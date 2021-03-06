package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

//var aa = 3
//var ss = "kkk"
//var bb = true

var (
	aa = 3
	ss = "kkk"
	bb = true
)

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d   %q\n", a, s)
}

func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func varableTypeDeduction() {
	var a, b, c, d = 3, 4, "abc", true
	b = 5
	fmt.Println(a, b, c, d)
}

func varableShort() {
	a, b, c, d := 3, 4, "abc", true
	b = 5
	fmt.Println(a, b, c, d)
}

func euler() {
	c := 3 + 4i
	fmt.Println(cmplx.Abs(c))

	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1)

	fmt.Printf("%.3f \n", cmplx.Exp(1i*math.Pi)+1)
}

func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func main() {

	fmt.Println("Hello , world!")
	variableZeroValue()
	variableInitialValue()
	varableTypeDeduction()
	varableShort()

	fmt.Println(aa, bb, ss)

	euler()

	triangle()
}
