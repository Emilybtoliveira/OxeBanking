package models

type Client struct {
	Id               int    `json:"id,omitempty"`
	User_id          int    `json:"user_id,omitempty"`
	Card_function    string `json:"card_function,omitempty"`
	Credit_limit     int    `json:"credit_limit,omitempty"`
	Set_credit_limit int    `json:"set_credit_limit,omitempty"`
}