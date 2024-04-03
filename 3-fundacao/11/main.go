package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Pessoa interface {
	Desativar()
}

type Empresa struct {
	Nome string
}

func (e Empresa) Desativar() {
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado.\n", c.Nome)
}

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

func main() {
	wesley := Cliente{
		Nome:  "Wesley",
		Idade: 30,
		Ativo: true,
	}

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t\n", wesley.Nome, wesley.Idade, wesley.Ativo)

	wesley.Ativo = false
	wesley.Cidade = "São Paulo"
	wesley.Endereco.Cidade = "Rio de Janeiro"

	fmt.Println(wesley.Cidade)

	// wesley.Desativar()

	Desativacao(wesley)

	minhaEmpresa := Empresa{}

	Desativacao(minhaEmpresa)	

}
