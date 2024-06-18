package main

// go get github.com/gorilla/mux -> para conseguir utilizar as rotas.
import (
	"api/API/src/router"
	"fmt"
	"log"
	"net/http"
)

func main(){
	fmt.Println("Rodando API!!")

	r := router.Gerar()

	log.Fatal(http.ListenAndServe(":5000", r))
}