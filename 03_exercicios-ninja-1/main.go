package main

import "fmt"

// ------------ exercício 1 ------------
/*
func ex1() {
	x := 42
	y := "James Bond"
	z := true

	fmt.Println(x, y, z)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
*/

// ------------ exercício 2 ------------
/*
var x int
var y string
var z bool

func ex2() {
	fmt.Printf("%v\n%v\n%v", x, y, z)
}
*/

// ------------ exercício 3 ------------
/*
var x int = 42
var y string = "James Bond"
var z bool = true

func ex3() {
	s := fmt.Sprintf("%v\t%v\t%v", x, y, z)
	fmt.Println(s)
}
*/

//  ------------ exercício 4 ------------
/*
type hotdog int

var x hotdog

func ex4() {
	fmt.Printf("%v\t%T\n", x, x)
	x = 42
	fmt.Printf("%v\t%T\n", x, x)

}
*/

//  ------------ exercício 5 ------------
type hotdog int

var x hotdog
var y int

func ex5() {
	fmt.Printf("%v\t%T\n", x, x)
	x = 42
	fmt.Printf("%v\t%T\n", x, x)
	y = int(x)
	fmt.Printf("%v\t%T\n", y, y)
}

func main() {
	//ex1()
	//ex2()
	//ex3()
	//ex4()
	ex5()
}
