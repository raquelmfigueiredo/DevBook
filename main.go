package main

// go get github.com/gorilla/mux -> para conseguir utilizar as rotas.
import (
	"api/API/config"
	"api/API/src/router"
	"fmt"
	"log"
	"net/http"
)

func main(){
	config.Carregar()

	fmt.Println("Rodando API!!")
	r := router.Gerar()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}