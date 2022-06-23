package main

import "fmt"

// ------------ exercício 1 ------------
/*
type pessoa struct {
	nome      string
	sobrenome string
	sabores   []string
}

func ex1() {
	pessoa1 := pessoa{
		nome:      "Lera",
		sobrenome: "Mechis",
		sabores:   []string{"baunilha", "abacaxi", "morango"},
	}

	pessoa2 := pessoa{
		nome:      "Olga",
		sobrenome: "Markov",
		sabores:   []string{"café", "chocolate", "pêssego"},
	}

	fmt.Printf("%v %v\n", pessoa1.nome, pessoa1.sobrenome)
	for _, v := range pessoa1.sabores {
		fmt.Println("-", v)
	}

	fmt.Println("\n")

	fmt.Printf("%v %v\n", pessoa2.nome, pessoa2.sobrenome)
	for _, v := range pessoa2.sabores {
		fmt.Println("-", v)
	}
}
*/
// ------------ exercício 2 ------------
/*
type pessoa struct {
	nome      string
	sobrenome string
	sabores   []string
}

func ex2() {

	mapa := make(map[string]pessoa)

	mapa["Mechis"] = pessoa{
		nome:      "Lera",
		sobrenome: "Mechis",
		sabores:   []string{"baunilha", "abacaxi", "morango"},
	}

	mapa["Markov"] = pessoa{
		nome:      "Olga",
		sobrenome: "Markov",
		sabores:   []string{"café", "chocolate", "pêssego"},
	}

	for _, v := range mapa {
		fmt.Println("Meu nome é", v.nome, "e meus sorvetes favoritos são:")
		for _, v := range v.sabores {
			fmt.Println("–", v)
		}
		fmt.Println("\n")
	}
}
*/
// ------------ exercício 3 ------------
/*
type veiculo struct {
	portas int
	cor    string
}

type caminhonete struct {
	veiculo
	traçãoNasQuatro bool
}

type sedan struct {
	veiculo
	modeloLuxo bool
}

func ex3() {
	automovel1 := caminhonete{
		veiculo: veiculo{
			portas: 4,
			cor:    "amarelo",
		},
		traçãoNasQuatro: true,
	}
	automovel2 := sedan{
		veiculo: veiculo{
			portas: 2,
			cor:    "preto",
		},
		modeloLuxo: false,
	}
	fmt.Println(automovel1)
	fmt.Println(automovel2)
	fmt.Println(automovel1.traçãoNasQuatro)
	fmt.Println(automovel2.cor)
}
*/
// ------------ exercício 4 ------------

func ex4() {
	x := struct {
		nome      string
		sabor     string
		ondetem   []string
		vaibemcom map[string]string
	}{
		nome:    "Shawarma",
		sabor:   "salgado",
		ondetem: []string{"em muitos lugares", "em São Paulo"},
		vaibemcom: map[string]string{
			"no café da manhã": "nunca tentei",
			"no almoço":        "uma delícia",
			"na janta":         "é um lanche leve",
		},
	}

	fmt.Println(x)
}

func main() {
	//ex1()
	//ex2()
	//ex3()
	ex4()
}
