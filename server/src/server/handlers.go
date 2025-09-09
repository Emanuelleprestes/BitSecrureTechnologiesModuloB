package server

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Emanuelleprestes/InfoSmart-Solutions.git/server/controlers"
	"github.com/Emanuelleprestes/InfoSmart-Solutions.git/server/models/colaborador"
)

type UserSessao struct {
	ID    int    `json:"id"`
	Nome  string `json:"nome"`
	Cargo string `json:"cargo"`
}

// aqui vai ficar as funçoes que vao ir para para as rotas
type (
	writer   = http.ResponseWriter
	resquest = *http.Request
)

type Handlers struct {
	conn *sql.DB
}

func (h *Handlers) Login(w writer, r resquest) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Content-Type deve ser application/json", http.StatusUnsupportedMediaType)
		return
	}
	var creds struct {
		Value string `json:"value"`
		Senha string `json:"senha"`
	}
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}
	usercontroller := controlers.NewColaboradorcontroller(h.conn)
	var user *colaborador.Colaborador
	user, err = usercontroller.Login(creds.Value, creds.Senha)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	usersessao := UserSessao{
		ID:    user.ID,
		Nome:  user.Nome,
		Cargo: user.Cargo,
	}
	sessionManager.Put(r.Context(), "user", usersessao)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "ok",
		"message": "Login efetuado com sucesso",
	})
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

func (h *Handlers) Createcolaborador(w writer, r resquest) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Content-Type deve ser application/json", http.StatusUnsupportedMediaType)
		return
	}
	var colaborador colaborador.Colaborador
	controlercolaborador := controlers.NewColaboradorcontroller(h.conn)
	err := json.NewDecoder(r.Body).Decode(&colaborador)
	if err != nil {
		http.Error(
			w,
			fmt.Errorf("o json esta errado: %w", err).Error(),
			http.StatusInternalServerError,
		)
		return
	}
	err = controlercolaborador.Create(&colaborador)
	if err != nil {
		http.Error(
			w,
			fmt.Errorf("erro ao criar o usuario: %w", err).Error(),
			http.StatusInternalServerError,
		)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"Status":  "200",
		"message": "colaborador criado",
	})
}

func (h *Handlers) Deletecolaborador(w writer, r resquest) {
	fmt.Println(w.Write([]byte("")))
}

func (h *Handlers) GetcolaboladoresByName(w writer, r resquest) {
	fmt.Println(w.Write([]byte("")))
}

func (h *Handlers) Updatecolaborador(w writer, r resquest) {
}
