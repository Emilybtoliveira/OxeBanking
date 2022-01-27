package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Emilybtoliveira/OxeBanking/dao"
	"github.com/Emilybtoliveira/OxeBanking/models"
)

func FormatResponseToJSON(w http.ResponseWriter, statusCode int, response interface{}) {
	if response == nil || response == false {
		response = "Cannot execute action. Client not found or with no active card."
	}

	json, err := json.Marshal(response)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(json)
}

//redireciona para a função GetCard() em DAO/cardDAO.go
func GetCardHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println()
	params := r.URL.Query()["user_id"][0]
	//params := mux.Vars(r)
	//fmt.Println(mux.Vars(r))

	id, err := strconv.Atoi(params)
	dao.CheckErr(err)
	fmt.Println(id)
	/* response := dao.GetCard(id)
	json.NewEncoder(w).Encode(response) */
	//io.WriteString(w, response)

	response, err := dao.GetCard(id)
	if response.User_id == 0 {
		FormatResponseToJSON(w, http.StatusNotFound, nil)
	} else if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	} else {
		FormatResponseToJSON(w, http.StatusOK, response)
	}

}

//redireciona para a função CreateCard() em DAO/cardDAO.go
func CreateCardHandler(w http.ResponseWriter, r *http.Request) {
	var card models.Card
	_ = json.NewDecoder(r.Body).Decode(&card)
	//fmt.Println(card)

	response, err := dao.CreateCard(card.User_id, card.Password, card.Owner)
	if response {
		FormatResponseToJSON(w, http.StatusOK, response)
	} else if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	} else {
		FormatResponseToJSON(w, http.StatusForbidden, response)
	}
}

//redireciona para a função SuspendCard() em DAO/cardDAO.go
func UpdateStatusHandler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()["user_id"][0]
	//params := mux.Vars(r)

	id, err := strconv.Atoi(params)
	dao.CheckErr(err)
	fmt.Println(id)

	response, err := dao.SuspendCard(id)
	if response {
		FormatResponseToJSON(w, http.StatusOK, response)
	} else if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	} else {
		FormatResponseToJSON(w, http.StatusForbidden, response)
	}
}

//redireciona para a função UpdateCardFunction() em DAO/cardDAO.go
func UpdateFunctionHandler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()["user_id"][0]
	//params := mux.Vars(r)

	id, err := strconv.Atoi(params)
	dao.CheckErr(err)
	fmt.Println(id)

	var client models.Client
	_ = json.NewDecoder(r.Body).Decode(&client)
	fmt.Println(client)

	response, err := dao.UpdateCardFunction(id, client.Credit_limit, client.Set_credit_limit)
	if response {
		FormatResponseToJSON(w, http.StatusOK, response)
	} else if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	} else {
		FormatResponseToJSON(w, http.StatusForbidden, response)
	}
}
