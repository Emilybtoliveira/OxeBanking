package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Emilybtoliveira/OxeBanking/dao"
	"github.com/Emilybtoliveira/OxeBanking/models"
)

//Função que redireciona para a função GetAllVirtualCard() em DAO/virtualCardDAO.go
func GetAllVirtualCardsHandler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()["user_id"][0]
	id, err := strconv.Atoi(params)
	dao.CheckErr(err)
	//fmt.Println(id)

	response, err := dao.GetAllVirtualCards(id)
	//fmt.Println(len(response))

	if response == nil {
		FormatResponseToJSON(w, http.StatusNotFound, nil)
	} else if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	} else {
		FormatResponseToJSON(w, http.StatusOK, response)
	}

}

//Função que redireciona para a função CreateVirtualCard() em DAO/virtualCardDAO.go
func CreateVirtualCardHandler(w http.ResponseWriter, r *http.Request) {
	var virtual_card models.VirtualCard
	_ = json.NewDecoder(r.Body).Decode(&virtual_card)

	response, err := dao.CreateVirtualCard(virtual_card.User_id, virtual_card.Owner, virtual_card.Nickname)

	if response.User_id == 0 {
		FormatResponseToJSON(w, http.StatusForbidden, response)
	} else if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
	} else {
		FormatResponseToJSON(w, http.StatusOK, response)
	}
}

//redireciona para a função RemoveVirtualCard() em DAO/virtualCardDAO.go
func RemoveVirtualCardByIDHandler(w http.ResponseWriter, r *http.Request) {
	paramUser := r.URL.Query()["user_id"][0]
	id, err := strconv.Atoi(paramUser)

	var virtual_cards models.VirtualCard
	_ = json.NewDecoder(r.Body).Decode(&virtual_cards)

	response, err := dao.RemoveVirtualCardByID(id, virtual_cards.Card_number)

	if response == false {
		FormatResponseToJSON(w, http.StatusNotFound, response)
	} else if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
	} else {
		FormatResponseToJSON(w, http.StatusOK, response)
	}
}
