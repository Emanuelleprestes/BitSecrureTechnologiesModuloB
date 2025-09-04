package server

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Emanuelleprestes/InfoSmart-Solutions.git/server/controlers"
)

// aqui vai ficar as funçoes que vao ir para para as rotas
type (
	writer   = http.ResponseWriter
	resquest = *http.Request
)

type Handlers struct {
	conn *sql.DB
}

func (h *Handlers) Getcolaboladores(w writer, r resquest) {
	colabcontroller := controlers.NewColaboradorcontroller(h.conn)
	colaboradores, err := colabcontroller.Getall()
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	data, err := json.Marshal(colaboradores)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fmt.Println(w.Write(data))
}

func (h *Handlers) Savecola(w writer, r resquest) {
	fmt.Println(w.Write([]byte("")))
}

func (h *Handlers) Deletecolaborador(w writer, r resquest) {
	fmt.Println(w.Write([]byte("")))
}

func (h *Handlers) GetcolaboladoresByName(w writer, r resquest) {
	fmt.Println(w.Write([]byte("")))
}

func (h *Handlers) Updatecolaborador(w writer, r resquest) {
}
