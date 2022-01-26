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

	//dao.CreateCard(2000, "asdd67a8sdaf67a6d8dsa7d8asd67a8sd7a8d6", "JOSE SILVA JUNIOR")

	router := mux.NewRouter()
	router.HandleFunc("/card/{id}", handlers.GetCardHandler).Methods("GET")
	router.HandleFunc("/card", handlers.CreateCardHandler).Methods("POST")
	router.HandleFunc("/card/function/{id}", handlers.UpdateFunctionHandler).Methods("PUT")
	router.HandleFunc("/card/status/{id}", handlers.UpdateStatusHandler).Methods("PUT")
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
