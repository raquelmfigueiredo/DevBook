package controllers

import "net/http"


func CriarUsuarios(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Criando Usuario!!"))
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