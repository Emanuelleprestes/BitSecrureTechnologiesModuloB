package server

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Emanuelleprestes/InfoSmart-Solutions.git/server/controlers"
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

func (h *Handlers) Loginbyemail(w writer, r resquest) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Content-Type deve ser application/json", http.StatusUnsupportedMediaType)
		return
	}

	var creds struct {
		Email string `json:"email"`
		Senha string `json:"senha"`
	}

	// decodifica o JSON do corpo
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}
	usercontroller := controlers.NewColaboradorcontroller(h.conn)
	user, err := usercontroller.Loginbyemail(creds.Email, creds.Senha)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	usersessao := UserSessao{
		ID:    user.ID,
		Nome:  user.Nome,
		Cargo: user.Cargo,
	}
	sessionManager.Put(r.Context(), "user", usersessao)
	sessionID := sessionManager.Token(r.Context())
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "ok",
		"message": "Login efetuado com sucesso",
		"session": sessionID,
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
