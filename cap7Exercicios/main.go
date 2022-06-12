package main

import "fmt"

// ------------ exercício 1 ------------
/*
func ex1() {

	for x := 1; x <= 10000; x++ {
		fmt.Printf("%v\t", x)
	}
}
*/

// ------------ exercício 2 ------------
/*
func ex2() {
	for x := 65; x <= 90; x++ {
		fmt.Println(x)
		for j := 1; j <= 3; j++ {
			fmt.Printf("\t%#U\n", x)
		}
	}
}
*/
// ------------ exercício 3 ------------
/*
func ex3() {
	anoNascimento := 1993
	anoContar := 2088
	for anoNascimento <= anoContar {
		fmt.Println(anoNascimento)
		anoNascimento++
	}
}
*/
// ------------ exercício 4 ------------
/*
func ex4() {
	anoNascimento := 1993
	anoContar := 2088

	for {
		if anoNascimento > anoContar {
			break
		}
		fmt.Println(anoNascimento)
		anoNascimento++
	}
}
*/
// ------------ exercício 5 ------------
/*
func ex5() {
	for x := 10; x <= 100; x++ {
		fmt.Println(x % 4)
	}
}
*/
// ------------ exercício 6 ------------
/*
func ex6() {
	x := 10
	if x == 10 {
		fmt.Println("x é igual de 10")
	}
}
*/
// ------------ exercício 7 ------------
/*
func ex7() {
	x := 15
	if x == 10 {
		fmt.Println("x é igual de 10")
	} else if x < 10 {
		fmt.Println("x é menor do que 10")
	} else {
		fmt.Println("x é maior do 10")
	}
}
*/

// ------------ exercício 8 ------------
/*
func ex8() {
	sono := 2

	switch {
	case sono == 0:
		fmt.Println("topo mais exercícios!")
	case sono == 1:
		fmt.Println("sono batendo..")
	case sono == 2:
		fmt.Println("preciso dormir!")
	}
}
*/

// ------------ exercício 9 ------------

func ex9() {
	esporteFavorito := "curling"

	switch esporteFavorito {
	case "futebol":
		fmt.Println("época de Copa do Mundo é sempre bom")
	case "curling":
		fmt.Println("adoroooo! Principalmente nas olimpíadas de inverno")
	case "corrida":
		fmt.Println("um dia na São Silvestre irei")

	}
}

func main() {
	//ex1()
	//ex2()
	//ex3()
	//ex4()
	//ex5()
	//ex6()
	//ex7()
	//ex8()
	ex9()
}
