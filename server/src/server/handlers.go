package server

import "net/http"

// aqui vai ficar as funçoes que vao ir para para as rotas
type (
	writer   = http.ResponseWriter
	resquest = *http.Request
)

func Teste2(w writer, r resquest)
func Getcolaboladores(w writer, r resquest)
func Save(w writer, r resquest)
func GetcolaboladoresByName(w writer, r resquest)
