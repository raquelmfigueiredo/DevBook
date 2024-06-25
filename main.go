package main

// go get github.com/gorilla/mux -> para conseguir utilizar as rotas.
import (
	"api/API/src/config"
	"api/API/src/router"
	"fmt"
	"log"
	"net/http"
)

func main(){
	config.Carregar()
	r := router.Gerar()

	fmt.Printf("Escutando na porta %d", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}