package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Emilybtoliveira/OxeBanking/dao"
)

//redireciona para a função GetAllVirtualCard() em DAO/virtualCardDAO.go
func GetAllVirtualCardsHandler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()["user_id"][0]
	id, err := strconv.Atoi(params)
	dao.CheckErr(err)
	fmt.Println(id)

	response, err := dao.GetAllVirtualCards(id)

	for i := range response {
		if response[i].User_id == 0 {
			FormatResponseToJSON(w, http.StatusNotFound, nil)
		} else if err != nil {
			http.Error(w, "Internal error", http.StatusInternalServerError)
			return
		} else {
			FormatResponseToJSON(w, http.StatusOK, response[i])
		}
	}
}

//redireciona para a função CreateVirtualCard() em DAO/virtualCardDAO.go
func CreateVirtualCardHandler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()["user_id"][0]
	id, err := strconv.Atoi(params)

	owner := r.URL.Query()["owner"][0]
	nickname := r.URL.Query()["nickname"][0]

	response, err := dao.CreateVirtualCard(id, owner, nickname)

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
	paramCard := r.URL.Query()["card_number"][0]

	id, err := strconv.Atoi(paramUser)
	cardId, err := strconv.Atoi(paramCard)

	response, err := dao.RemoveVirtualCardByID(id, cardId)

	if response == false {
		FormatResponseToJSON(w, http.StatusNotFound, response)
	} else if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
	} else {
		FormatResponseToJSON(w, http.StatusOK, response)
	}
}
