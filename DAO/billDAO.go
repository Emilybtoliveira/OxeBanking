package DAO

import (

)

func GetBill() {
	//recebe id do usuário e retorna valor da fatura
}

func GetCreditLimit() {
	//recebe id do usuário e retorna limite máximo e limite setado pelo usuário
}

func AdjustCreditLimit() {
	//recebe id do usuário e novo limite setado pelo usuário. Retorna true/false
}

func DeductDebt() {
	//recebe id do usuário e valor para abater da fatura. Retorna true/false
}
