package main

import "fmt"

// ------------ exercício 1 ------------
/*
func ex1() {
	array := [5]int{1, 2, 3, 4, 5}

	for i, v := range array {
		fmt.Println(i, v)
	}
	fmt.Printf("%T", array)
}
*/
// ------------ exercício 2 ------------
/*
func ex2() {
	slice := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	for i, v := range slice {
		fmt.Println(i, v)
	}
	fmt.Printf("%T", slice)
}
*/
// ------------ exercício 3 ------------
/*
func ex3() {
	slice := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	fmt.Println(slice[:3])
	fmt.Println(slice[4:])
	fmt.Println(slice[1:7])
	fmt.Println(slice[2:9])
	fmt.Println(slice[2 : len(slice)-1])
}
*/
// ------------ exercício 4 ------------
/*
func ex4() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	x = append(x, 52)
	fmt.Println(x)
	x = append(x, 53, 54, 55)
	fmt.Println(x)
	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)
	fmt.Println(x)
}
*/
// ------------ exercício 5 ------------

func ex5() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	y := append(x[:3], x[6:]...)
	fmt.Println(y)
}

func main() {
	//ex1()
	//ex2()
	//ex3()
	//ex4()
	ex5()
}
