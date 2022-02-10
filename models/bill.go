package models

type Bill struct {
	User_id      int     `json:"user_id,omitempty"`
	Total_amount float64 `json:"total_amount,omitempty"`
	Amount_paid  float64 `json:"amount_paid,omitempty"`
	Closing_date string  `json:"closing_date,omitempty"` //no formato aaaa-mm-dd
	Due_date     string  `json:"due_date,omitempty"` //no formato aaaa-mm-dd
	Status       string  `json:"status,omitempty"`
	Month_year   string  `json:"month_year,omitempty"` //no formato mm/aaaa
}
