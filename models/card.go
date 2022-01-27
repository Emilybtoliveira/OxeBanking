package models

type Card struct {
	User_id          int    `json:"id,omitempty"`
	Card_number      int    `json:"card_number,omitempty"`
	Password         string `json:"encrypted_password,omitempty"`
	Owner            string `json:"card_owner,omitempty"`
	Valid_thru       string `json:"valid_thru,omitempty"`
	Cvv              int    `json:"cvv,omitempty"`
	Card_function    string `json:"card_function,omitempty"`
	Status           string `json:"status,omitempty"`
	Credit_limit     int    `json:"credit_limit,omitempty"`
	Set_credit_limit int    `json:"set_credit_limit,omitempty"`
}
