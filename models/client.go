package models

type Client struct {
	Id               int    `json:"id"`
	User_id          int    `json:"user_id"`
	Card_function    string `json:"card function"`
	Credit_limit     string `json:"credit limit"`
	Set_credit_limit string `json:"set credit limit"`
}