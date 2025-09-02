package server

import (
	"fmt"
	"net/http"
)

// aqui vai ficar as funçoes que vao ir para para as rotas
type (
	writer   = http.ResponseWriter
	resquest = *http.Request
)

func Teste2(w writer, r resquest) {
	fmt.Println(w.Write([]byte("")))
}

func Getcolaboladores(w writer, r resquest) {
	fmt.Println(w.Write([]byte("")))
}

func Save(w writer, r resquest) {
	fmt.Println(w.Write([]byte("")))
}

func GetcolaboladoresByName(w writer, r resquest) {
	fmt.Println(w.Write([]byte("")))
}
