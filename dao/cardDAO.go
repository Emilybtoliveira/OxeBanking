package dao

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/Emilybtoliveira/OxeBanking/models"
	_ "github.com/lib/pq"
)

//Função para gerar data de validade do cartão; Recebe a quantidade de anos à frente em offset;
func generateValidThru(offset int) string {
	now_plus_offset := time.Now().AddDate(offset, 0, 0) //adicionando 5 anos à data atual

	valid_thru := fmt.Sprintf("%02d/%d", now_plus_offset.Month(), now_plus_offset.Year())

	return valid_thru
}

//Função que gera um número de len(min,max) digitos;
func generateRandomNumbers(min, max int) int {
	/* https://golangdocs.com/generate-random-numbers-in-golang */
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min)
}

func selectClient(user_id int) bool {
	fmt.Println("Checando se o cliente já existe..")
	query := fmt.Sprintf("SELECT * FROM client WHERE user_id = %d;", user_id)

	fmt.Println(query)
	stmt1, err1 := db.Query(query)
	CheckErr(err1)

	if stmt1.Next() != true {
		return false
	} else {
		return true
	}
}

//Função de criação de cartão físico; Recebe informações do usuário e retorna o número do cartão;
func CreateCard(user_id int, password string, owner string) (bool, error) {
	var query string

	//verificando se o user_id já é registrado no banco
	client_exists := selectClient(user_id)

	if !client_exists { //se entrar aqui, então nao existe um usuário cadastrado com esse id; nesse caso, é preciso inserí-lo em client antes de cadastrar um novo cartao
		fmt.Println("Cadastrando o cliente...")

		query = fmt.Sprintf("INSERT INTO public.client(user_id) VALUES(%d)", user_id)

		fmt.Println(query)

		stmt2, err1 := db.Exec(query)
		CheckErr(err1)
		_ = stmt2
	} else { //Regra de negocio: se já existe o cliente, é preciso conferir se ele já possui um cartão físico com status ativo
		fmt.Println("Checando se o cliente já possui cartão ativo...")

		query = fmt.Sprintf("SELECT status from physical_cards WHERE user_id = %d;", user_id)
		fmt.Println(query)

		stmt3, err1 := db.Query(query)
		CheckErr(err1)

		var status string

		for stmt3.Next() {
			err = stmt3.Scan(&status)
			CheckErr(err)

			if status == "ativo" {
				fmt.Println("Cliente já possui cartão ativo.")
				return false, err1
			}
		}
	}

	//Criando novo cartão
	fmt.Println("Gerando novo cartão...")

	//Gerando a data de vencimento do cartão
	valid_thru := generateValidThru(5)
	fmt.Printf("Validade do cartão gerada: %s\n", valid_thru)

	//Gerando o número do cartao
	card_number := generateRandomNumbers(1000000000000000, 9999999999999999)
	fmt.Printf("Número do cartão gerado: %d\n", card_number)
	/* OBS: aqui, após gerar o numero do cartão, é preciso fazer um select em physical_cards para
	checar se já não existe algum usuário com esse número de cartão. */

	//Gerando o número do cvv
	cvv := generateRandomNumbers(100, 999)
	fmt.Printf("Número do CVV gerado: %d\n", cvv)

	/* Feito isso, aqui vai a query de inserir um novo cartão em physical_cards; lembrando que nao precisa informar status nem emission_date */
	query = fmt.Sprintf("INSERT INTO public.physical_cards(user_id, card_number, four_digit_password, owner, valid_thru, cvv) VALUES (%d, %d, '%s', '%s', '%s', %d);", user_id, card_number, password, owner, valid_thru, cvv)
	fmt.Println(query)

	stmt1, err2 := db.Query(query)
	CheckErr(err2)
	_ = stmt1

	return true, err2
}

//recebe o id do usuário e retorna o número do cartão, titular, valid_thru
func GetCard(user_id int) (models.Card, error) {
	query := fmt.Sprintf("SELECT card_number, owner, valid_thru FROM physical_cards WHERE user_id = %d and status='ativo';", user_id)

	//fmt.Println(query)
	stmt1, err1 := db.Query(query)
	CheckErr(err1)

	query = fmt.Sprintf("SELECT card_function, credit_limit, set_credit_limit FROM client WHERE user_id = %d;", user_id)
	stmt2, err2 := db.Query(query)
	CheckErr(err2)

	var card models.Card

	if stmt1.Next() != true {
		//retornando sem clientes
		fmt.Println("Cliente inexistente ou sem cartões ativos.")
	} else if stmt2.Next() != true {
		fmt.Println("Cliente inexistente.")
	} else {
		//cliente encontrado
		err = stmt1.Scan(&card.Card_number, &card.Owner, &card.Valid_thru)
		CheckErr(err)

		err = stmt2.Scan(&card.Card_function, &card.Credit_limit, &card.Set_credit_limit)
		CheckErr(err)

		card.User_id = user_id
		fmt.Printf("Encontrado: %d, %d, %s, %s %s\n", card.User_id, card.Card_number, card.Owner, card.Valid_thru, card.Card_function)
	}

	return card, err
}

//recebe o id do usuário, altera status do cartão para suspenso e retorna true/false
func SuspendCard(user_id int) (bool, error) {
	client_exists := selectClient(user_id) //verifica se o user_id existe primeiro

	if !client_exists {
		fmt.Println("Cliente não existe.")
		return false, nil
	}

	fmt.Println("Cliente existe.")
	query := fmt.Sprintf("UPDATE physical_cards SET status='bloqueado' WHERE user_id = %d and status = 'ativo';", user_id)

	//fmt.Println(query)
	stmt1, err1 := db.Exec(query)

	if err1 != nil {
		fmt.Println("Problema na solicitacao")
		return false, err1
	}

	rows_affected, err2 := stmt1.RowsAffected()
	CheckErr(err2)
	if rows_affected == 0 {
		fmt.Println("Cliente não possui cartão ativo.")
		return false, err1
	}

	query = fmt.Sprintf("SELECT * from physical_cards WHERE user_id = %d and status = 'ativo';", user_id)
	stmt2, err1 := db.Query(query)
	_ = stmt2

	if err1 != nil {
		fmt.Println("Algo deu errado.")
		return false, err1
	} else {
		fmt.Println("Cartao bloqueado.")
		return true, err1
	}
}

func UpdateCardFunction(user_id int, credit_limit int, set_credit_limit int) (bool, error) {
	client_exists := selectClient(user_id) //verifica se o user_id existe primeiro

	if !client_exists {
		fmt.Println("Cliente não existe.")
		return false, nil
	}

	fmt.Println("Cliente existe.")
	query := fmt.Sprintf("UPDATE client SET card_function='Credito/Debito', credit_limit = %d, set_credit_limit=%d WHERE user_id = %d;", credit_limit, set_credit_limit, user_id)

	fmt.Println(query)
	stmt1, err1 := db.Exec(query)
	_ = stmt1

	if err1 != nil {
		fmt.Println("Problema na solicitacao")
		return false, err1
	} else {
		fmt.Println("Informações atualizadas.")
		return true, err1
	}
}
