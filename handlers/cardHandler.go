package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Emilybtoliveira/OxeBanking/dao"
	"github.com/gorilla/mux"
)

func FormatResponseToJSON(w http.ResponseWriter, statusCode int, response interface{}) {
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
	params := mux.Vars(r)
	//fmt.Println(params["id"])

	id, err := strconv.Atoi(params["id"])
	_ = err
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

}

//redireciona para a função SuspendCard() em DAO/cardDAO.go
func SuspendCardHandler(w http.ResponseWriter, r *http.Request) {

}
