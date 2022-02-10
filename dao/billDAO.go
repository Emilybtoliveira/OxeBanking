package dao

import (
	"fmt"
	"regexp"
	"time"

	"github.com/Emilybtoliveira/OxeBanking/models"
	_ "github.com/lib/pq"
)

func IsThereAnOpenBill(user_id int) bool {
	var closing_date, month_year string
	//now := time.Now()

	query := fmt.Sprintf("SELECT closing_date, month_year FROM bill WHERE user_id = %d AND status = 'aberta'", user_id)
	stmt, err := db.Query(query)
	CheckErr(err)

	/* Supõe que só existe uma fatura aberta*/
	if stmt.Next() {
		err = stmt.Scan(&closing_date, &month_year)
		CheckErr(err)

		today := fmt.Sprintf("%d-%02d-%02d", time.Now().Year(), time.Now().Month(), time.Now().Day())
		if today >= closing_date {
			query := fmt.Sprintf("UPDATE bill SET status = 'fechada' WHERE user_id = %d AND month_year = '%s'", user_id, month_year)
			stmt, err := db.Query(query)
			CheckErr(err)

			_ = stmt

			fmt.Println("Fatura em aberto fechada.")
			return false
		}

		return true
	} else {
		fmt.Println("Não há faturas em aberto.")
		return false
	}
}

func ValidateDateFormat(date string) bool {
	regex := regexp.MustCompile("(0?[1-9]|1[012])/((19|20)\\d\\d)")
	return regex.MatchString(date)
}

//Função que gera a fatura do mes
func GenerateBill(user_id int, closing_date string, due_date string, month_year string) (bool, error) {
	var query string

	client_exists := selectClient(user_id)

	if !client_exists {
		print("Cliente não encontrado.")
		return false, err
	}

	/* CRIAR FORMA DE VALIDAR O FORMATO DE month_year */
	//ValidateDateFormat(month_year)

	/* CRIAR FORMA DE VALIDAR O FORMATO DE due_date E closing_date */

	/* Aqui, é verificado antes se já existe alguma fatura aberta. Se esta estiver ainda no prazo pra estar aberta, então não pode ser gerada nova fatura.
	Se a fatura não estiver mais no prazo, ela é fechada, e uma nova fatura pode ser gerada.*/
	havia_fatura_aberta := IsThereAnOpenBill(user_id)

	if havia_fatura_aberta {
		fmt.Println("Cliente possui fatura em aberto.")
		return false, err
	}

	query = fmt.Sprintf("SELECT * FROM bill WHERE user_id = %d AND month_year = '%s'", user_id, month_year)
	stmt, err := db.Exec(query)
	CheckErr(err)

	rows_affected, err := stmt.RowsAffected()
	CheckErr(err)

	if rows_affected != 0 {
		fmt.Println("Cliente já possui fatura referente a esse mês e ano.")
		return false, err
	}

	query = fmt.Sprintf("INSERT INTO bill(user_id, closing_date, due_date, month_year) VALUES (%d, '%s', '%s', '%s')", user_id, closing_date, due_date, month_year)
	stmt2, err := db.Query(query)
	CheckErr(err)

	_ = stmt2

	if err != nil {
		fmt.Println("Problema na solicitacao")
		return false, err
	} else {
		fmt.Println("Fatura gerada.")
		return true, err
	}
}

//Função que recebe id do usuário e retorna valor da fatura
func GetBill(user_id int) (models.Bill, error) {
	var bill models.Bill
	var query string

	query = fmt.Sprintf("SELECT * FROM bill WHERE user_id = %d AND status = 'aberta'", user_id)
	stmt, err := db.Query(query)
	CheckErr(err)

	if !stmt.Next() {
		fmt.Println("Cliente não existe ou não possui faturas em aberto.")
	} else {
		err = stmt.Scan(&bill.User_id, &bill.Total_amount, &bill.Amount_paid, &bill.Closing_date, &bill.Due_date, &bill.Status, &bill.Month_year)
		CheckErr(err)
	}
	return bill, err
}

//Função que recebe id do usuário e retorna limite máximo e limite setado pelo usuário
func GetCreditLimit(user_id int) (models.Client, error) {
	var client models.Client
	var query string

	query = fmt.Sprintf("SELECT credit_limit,set_credit_limit FROM client WHERE user_id=%d", user_id)
	fmt.Println(query)

	stmt, err := db.Query(query)
	CheckErr(err)

	if !stmt.Next() {
		fmt.Println("Cliente não existe.")
	} else {
		err = stmt.Scan(&client.Credit_limit, &client.Set_credit_limit)
		CheckErr(err)
	}

	return client, err
}

//Função que recebe id do usuário e novo limite setado pelo usuário. Retorna true/false
func AdjustCreditLimit(user_id, credit_limit, set_credit_limit int) (bool, error) {
	var query string
	var returned_credit_limit int

	/* Verifica primeiro se o cliente tem função crédito */
	query = fmt.Sprintf("SELECT credit_limit FROM client WHERE user_id = %d;", user_id)
	stmt, err := db.Query(query)
	CheckErr(err)

	if stmt.Next() {
		err = stmt.Scan(&returned_credit_limit)
	} else {
		fmt.Println("Cliente não existe.")
		return false, err
	}

	if returned_credit_limit == 0 {
		fmt.Println("Cliente não possui função crédito.")
		return false, err
	} else if set_credit_limit > credit_limit {
		fmt.Println("Cliente não pode setar um limite maior que o que possui.")
		return false, err
	}

	query = fmt.Sprintf("UPDATE client SET credit_limit = %d, set_credit_limit=%d WHERE user_id = %d;", credit_limit, set_credit_limit, user_id)

	fmt.Println(query)
	stmt1, err := db.Exec(query)
	_ = stmt1

	if err != nil {
		fmt.Println("Problema na solicitacao")
		return false, err
	} else {
		fmt.Println("Informações atualizadas.")
		return true, err
	}
}

//Função que recebe id do usuário e valor para abater da fatura. Retorna true/false
func DeductDebt(user_id int, valor float64) (models.Bill, error) {
	var bill models.Bill
	var query string
	var valor_fatura, valor_ja_pago float64

	query = fmt.Sprintf("SELECT total_amount, amount_paid FROM bill WHERE user_id = %d and fatura='aberta'", user_id)
	stmt, err := db.Query(query)
	CheckErr(err)

	if stmt.Next() {
		err = stmt.Scan(&valor_fatura, &valor_ja_pago)
	} else {
		fmt.Println("Cliente não existe ou não possui faturas em aberto.")
		return bill, err
	}

	if valor > valor_fatura {
		fmt.Println("Valor a ser pago maior que o valor da fatura.")
		return bill, err
	} else {
		query = fmt.Sprintf("UPDATE bill SET total_amount = %f, amount_paid = %f WHERE user_id=%d", valor_fatura-valor, valor_ja_pago+valor, user_id)
		fmt.Println(query)
		stmt1, err := db.Exec(query)
		_ = stmt1

		if err != nil {
			fmt.Println("Problema na solicitacao")
		} else {
			fmt.Println("Fatura atualizada.")

		}
	}
	return GetBill(user_id)
}
