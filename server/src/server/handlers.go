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
	Email string `json:"email"`
	Cargo string `json:"cargo"`
}

// aqui vai ficar as funçoes que vao ir para para as rotas
type (
	writer   = http.ResponseWriter
	resquest = *http.Request
)

type colaHandler struct {
	conn *sql.DB
}

func (h *colaHandler) Me(w writer, r resquest) {
	user, ok := sessionManager.Get(r.Context(), "user").(UserSessao)
	if !ok {
		http.Error(
			w,
			fmt.Errorf("erro ao confimar o usuario: %v", ok).Error(),
			http.StatusInternalServerError,
		)
	}
	fmt.Println(user)
	userecon := controlers.NewColaboradorcontroller(h.conn)
	userdb, err := userecon.Getbyemail(user.Email)
	if err != nil {
		http.Error(
			w,
			fmt.Errorf("não atorizado a entrar: %w", err).Error(),
			http.StatusUnauthorized,
		)
		return
	}
	userdb.ID = 0
	userdb.Senha = ""
	json.NewEncoder(w).Encode(map[string]any{
		"status": "200",
		"message": map[string]string{
			"cpf":        userdb.CPF,
			"nome":       userdb.Nome,
			"cargo":      userdb.Cargo,
			"setor":      userdb.Setor,
			"status":     userdb.Status,
			"email":      userdb.Email,
			"ramal":      userdb.Ramal,
			"habilidade": userdb.Habilidades,
		},
	})
}

func (h *colaHandler) Login(w writer, r resquest) {
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
	fmt.Println("login: ", user)
	usersessao := UserSessao{
		ID:    user.ID,
		Nome:  user.Nome,
		Email: user.Email,
		Cargo: user.Cargo,
	}
	sessionManager.Put(r.Context(), "user", usersessao)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "ok",
		"message": "Login efetuado com sucesso",
	})
}

func (h *colaHandler) Getcolaboladores(w writer, r resquest) {
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
	w.Write(data)
}

func (h *colaHandler) Createcolaborador(w writer, r resquest) {
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

func (h *colaHandler) Deletecolaborador(w writer, r resquest) {
	fmt.Println(w.Write([]byte("")))
}

func (h *colaHandler) GetcolaboladoresByName(w writer, r resquest) {
	fmt.Println(w.Write([]byte("")))
}

func (h *colaHandler) Updatecolaborador(w writer, r resquest) {
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
	val := sessionManager.Get(r.Context(), "user")
	user, ok := val.(UserSessao)
	if !ok {
		http.Error(
			w,
			fmt.Errorf("erro ao confimar o usuario: %v", ok).Error(),
			http.StatusInternalServerError,
		)
	}
	if user.Nome != colaborador.Nome && user.Cargo != "gestor" {
		http.Error(
			w,
			fmt.Errorf("o usuario não pode modificar outros").Error(),
			http.StatusUnauthorized,
		)
	}
	err = controlercolaborador.Update(&colaborador)
	if err != nil {
		http.Error(
			w,
			fmt.Errorf("erro ao modifcar o usuario: %w", err).Error(),
			http.StatusInternalServerError,
		)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"Status":  "200",
		"message": "colaborador atualizar",
	})
}
