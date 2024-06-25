package controllers

import (
	"api/API/src/banco"
	"api/API/src/modelos"
	"api/API/src/repositorios"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"oi"
)


func CriarUsuarios(w http.ResponseWriter, r *http.Request){
	corpoRequest, erro := io.ReadAll(r.Body)
	if erro != nil {
		log.Fatal(erro)
	}

	var usuario modelos.Usuario
	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil {
		log.Fatal(erro)
	}

	db, erro := banco.Conectar()
	if erro != nil {
		log.Fatal(erro)
	}

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	usuarioId, erro := repositorio.Criar(usuario)
	if erro != nil {
		log.Fatal(erro)
	}

	w.Write([]byte(fmt.Sprintf("Id inserido %d", usuarioId)))
}


func BuscarUsuarios(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Buscando todos os Usuarios!!"))
}

func BuscarUsuario(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Buscando Usuario!!"))
}

func AtualizarUsuarios(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Atualizando Usuario!!"))
}

func DeletarUsuarios(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Deletando Usuario!!"))
}