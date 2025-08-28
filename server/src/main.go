package main

import (
	"fmt"

	"github.com/Emanuelleprestes/InfoSmart-Solutions.git/server" // pegando o modulo do projeto
)

// essa função vai ser responsavel por configuras algumas coisas basica
// tipo conexão com o banco de dados que vai ser passado como referencia para uma estrutura
// essa função roda antes da função main
func init() {
	fmt.Println("teste")
}

func main() {
	// instaciando a estrtura do servidor
	// essa estrutura precisa do ip|porta e uma conexão com um banco de dados
	// ele esta aceitando nil porque eu recebo uma referencia para a conexão com banco de dados
	// por causa disso em toda a interação com o bancos de dados e melhor verificar se nulo ou não
	server := server.Newserver(":8080", nil)
	// logo mais isso vai virar uma gorutines, por agora não precisa
	// para poder parar o servidor via input e etc
	if err := server.Run(); err != nil {
		fmt.Println(err)
	}
}
