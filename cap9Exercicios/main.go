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

/*
func ex5() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	y := append(x[:3], x[6:]...)
	fmt.Println(y)
}
*/
// ------------ exercício 6 ------------
/*
func ex6() {
	slice := make([]string, 26, 26)
	slice = []string{"Acre", "Alagoas", "Amapá", "Amazonas",
		"Bahia", "Ceará", "Espírito Santo", "Goiás", "Maranhão",
		"Mato Grosso", "Mato Grosso do Sul", "Minas Gerais", "Pará",
		"Paraíba", "Paraná", "Pernambuco", "Piauí", "Rio de Janeiro",
		"Rio Grande do Norte", "Rio Grande do Sul", "Rondônia", "Roraima",
		"Santa Catarina", "São Paulo", "Sergipe", "Tocantins"}

	fmt.Println(len(slice), cap(slice))

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}
}
*/
// ------------ exercício 7 ------------
/*
func ex7() {
	ss := [][]string{
		[]string{"Denise", "Mignoli", "Falar sobre signos"},
		[]string{"Lera", "Mechis", "Viajar"},
		[]string{"Olga", "Markov", "Escrever"},
	}

	for _, v := range ss {
		fmt.Println(v)
	}

	fmt.Println("\n")

	for _, v := range ss {
		fmt.Println(v[0])
		for _, item := range v {
			fmt.Println("\t", item)
		}
	}
}
*/
// ------------ exercício 8 ------------
/*
func ex8() {
	exemplo := map[string][]string{
		"mignoli_denise": []string{"falar sobre signos"},
		"mechis_lera":    []string{"viajar"},
		"markov_olga":    []string{"escrever"},
	}

	for k, v := range exemplo {
		fmt.Println(k)
		for i, h := range v {
			fmt.Println("\t", i, "-", h)
		}
	}
}
*/
// ------------ exercício 9 ------------
/*
func ex9() {
	exemplo := map[string][]string{
		"mignoli_denise": []string{"falar sobre signos", "estudar idiomas"},
		"mechis_lera":    []string{"viajar", "tirar fotos"},
		"markov_olga":    []string{"escrever", "assistir filme"},
	}

	exemplo["joaquina_maria"] = []string{"estudar", "sair com os amigos"}

	for k, v := range exemplo {
		fmt.Println(k)
		for i, h := range v {
			fmt.Println("\t", i, "-", h)
		}
	}
}
*/

// ------------ exercício 10 ------------

func ex10() {
	exemplo := map[string][]string{
		"mignoli_denise": []string{"falar sobre signos", "estudar idiomas"},
		"mechis_lera":    []string{"viajar", "tirar fotos"},
		"markov_olga":    []string{"escrever", "assistir filme"},
	}

	exemplo["joaquina_maria"] = []string{"estudar", "sair com os amigos"}

	delete(exemplo, "joaquina_maria")

	for k, v := range exemplo {
		fmt.Println(k)
		for i, h := range v {
			fmt.Println("\t", i, "-", h)
		}
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
	//ex9()
	ex10()
}
