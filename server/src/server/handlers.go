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

func Getcolaboladores(w writer, r resquest) {
	fmt.Println(w.Write([]byte("")))
}

func Savecola(w writer, r resquest) {
	fmt.Println(w.Write([]byte("")))
}

func Deletecolaborador(w writer, r resquest) {
	fmt.Println(w.Write([]byte("")))
}

func GetcolaboladoresByName(w writer, r resquest) {
	fmt.Println(w.Write([]byte("")))
}
