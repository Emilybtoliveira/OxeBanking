package models

type Card struct {
	User_id       int    `json:"id,omitempty"`
	Card_number   int    `json:"card number,omitempty"`
	Status        string `json:"card status,omitempty"`
	Password      string `json:"encrypted password,omitempty"`
	Owner         string `json:"card owner,omitempty"`
	Valid_thru    string `json:"valid thru,omitempty"`
	Cvv           int    `json:"cvv,omitempty"`
	Emission_date string `json:"emission date,omitempty"`
}
