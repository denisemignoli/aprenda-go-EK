package main

import "fmt"

// ------------ exercício 1 ------------
/*
var a int = 28

func ex1() {
	fmt.Printf("%d, %b e %#x", a, a, a)
}
*/

// ------------ exercício 2 ------------
/*
func ex2() {
	a := (10 == 10)
	b := (10 != 10)
	c := (10 <= 10)
	d := (10 < 10)
	e := (10 >= 10)
	f := (10 > 10)

	fmt.Printf("%v\n%v\n%v\n%v\n%v\n%v\n", a, b, c, d, e, f)
}
*/

// ------------ exercício 3 ------------
/*
func ex3() {
	const x int = 45
	const y = 45
	fmt.Println(x, y)
}
*/

// ------------ exercício 4 ------------
/*
func ex4() {
	var x int = 10
	fmt.Printf("%d, %b, %#x\n", x, x, x)
	y := x << 1
	fmt.Printf("%d, %b, %#x\n", y, y, y)
}
*/

// ------------ exercício 5 ------------
/*
func ex5() {
	x := `Hello, testando
			de novo`
	fmt.Println(x)
}
*/

// ------------ exercício 6 ------------

const (
	_ = 1993 + iota
	b
	c
	d
	e
)

func ex6() {
	fmt.Printf("%v\t%v\t%v\t%v\t", b, c, d, e)
}

func main() {
	//ex1()
	//ex2()
	//ex3()
	//ex4()
	//ex5()
	ex6()
}
