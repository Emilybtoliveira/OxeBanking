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

	/* //dao.CreateCard(2000, "asdd67a8sdaf67a6d8dsa7d8asd67a8sd7a8d6", "JOSE SILVA JUNIOR")
	dao.CreateVirtualCard(2001, "EMILY B OLIVEIRA", "sla")
	dao.CreateVirtualCard(2001, "EMILY B OLIVEIRA", "sla2")
	card, err := dao.CreateVirtualCard(2001, "EMILY B OLIVEIRA", "sla3")
	dao.CheckErr(err)
	dao.GetAllVirtualCards(2001)
	dao.RemoveVirtualCardByID(2001, card.Card_number)
	dao.GetAllVirtualCards(2001) */

	router := mux.NewRouter()
	router.HandleFunc("/card", handlers.GetCardHandler).Methods("GET")
	router.HandleFunc("/card", handlers.CreateCardHandler).Methods("POST")
	router.HandleFunc("/card/function", handlers.UpdateFunctionHandler).Methods("PUT")
	router.HandleFunc("/card/status", handlers.UpdateStatusHandler).Methods("PUT")
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
