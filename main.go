package main

import (
	//"fmt"
	//"github.com/Emilybtoliveira/OxeBanking/models"
	//"github.com/Emilybtoliveira/OxeBanking/handlers"

	"github.com/Emilybtoliveira/OxeBanking/dao"
	"github.com/Emilybtoliveira/OxeBanking/handlers"
	"github.com/gorilla/mux"

	//"encoding/json"
	"log"
	"net/http"
)

func main() {

	dao.CreateDB()
	defer dao.CloseDB()

	dao.CreateTables()

	dao.CreateCard(2000, "", "JOSE SILVA JUNIOR", "asdd67a8sdaf67a6d8dsa7d8asd67a8sd7a8d6")
	//dao.CreateCard(2001, "Credito")
	//dao.CreateCard(2002, "Credito")

	router := mux.NewRouter()
	router.HandleFunc("/card/{id}", handlers.GetCardHandler).Methods("GET")
	router.HandleFunc("/card/{id}", handlers.CreateCardHandler).Methods("POST")
	router.HandleFunc("/contato/{id}", handlers.SuspendCardHandler).Methods("PUT")
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
