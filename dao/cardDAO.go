package dao

import (
	"fmt"
	"math/rand"
	"time"

	_ "github.com/lib/pq"
)

type card struct {
	user_id       int64
	card_number   int64
	status        string
	password      string
	owner         string
	cvv           int16
	emission_date string
}

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

//Função de criação de cartão físico; Recebe informações do usuário e retorna o número do cartão;
func CreateCard(user_id int64, card_function string, owner string, password string) bool {

	fmt.Println("Checando se o cliente já existe..")
	query := fmt.Sprintf("SELECT * FROM client WHERE user_id = %d;", user_id) //verificando se o user_id já é registrado no banco

	fmt.Println(query)
	stmt1, err1 := db.Query(query)
	CheckErr(err1)

	if stmt1.Next() != true { //se entrar aqui, então nao existe um usuário cadastrado com esse id; nesse caso, é preciso inserí-lo em client antes de cadastrar um novo cartao
		//esse abaixo é um exemplo de como ler o retorno de um select
		/* var client models.Client

		err = stmt1.Scan(&client.Id, &client.User_id, &client.Card_function, &client.Credit_limit, &client.Set_credit_limit)
		CheckErr(err)

		fmt.Printf("%d %d %s\n", client.Id, client.User_id, client.Card_function) */

		fmt.Println("Cadastrando o cliente...")

		query = ""

		if card_function == "" {
			query = fmt.Sprintf("INSERT INTO public.client(user_id) VALUES(%d)", user_id)
		} else {
			query = fmt.Sprintf("INSERT INTO public.client(user_id, card_function) VALUES(%d, '%s')", user_id, card_function)
		}
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

			fmt.Printf("%s\n", status)

			if status == "ativo" {
				fmt.Println("Cliente já possui cartão ativo.")
				return false
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
	query = fmt.Sprintf("INSERT INTO public.physical_cards(user_id, card_number, four_digit_password, owner, valid_thru, cvv) VALUES (%d, %d, '%s', '%s', %s, %d);", user_id, card_number, password, owner, valid_thru, cvv)
	fmt.Printf("Inserindo os seguintes dados: ID: %d, Numero do Cartao: %d, Senha: '%s', Proprietario: '%s', Validade: %s, CVV: %d\n", user_id, card_number, password, owner, valid_thru, cvv)
	stmt1, err2 := db.Query(query)
	CheckErr(err2)
	return true
}

func GetCard(user_id int64) bool {
	//recebe o id do usuário e retorna algumas informações do cartão
	query := fmt.Sprintf("SELECT * FROM physical_cards WHERE user_id = %d;", user_id) //verificando se o user_id já é registrado no banco

	//fmt.Println(query)
	stmt1, err1 := db.Query(query)
	CheckErr(err1)

	if stmt1.Next() != true {
		//retornando sem clientes
		fmt.Println("Cliente inexistente")
		return false
	} else {
		//cliente encontrado
		useless_var := 0
		cardInfo := card{}
		err = stmt1.Scan(&cardInfo.user_id, &cardInfo.card_number, &cardInfo.status, &cardInfo.password, &cardInfo.owner, &useless_var, &cardInfo.cvv, &cardInfo.emission_date)
		CheckErr(err)

		fmt.Printf("Nome do cliente: %s\nNumero do cartao: %d, Status do cartao: %s\n", cardInfo.owner, cardInfo.card_number, cardInfo.status)
	}

	return true
}

func SuspendCard(user_id int64) bool {
	//recebe o id do usuário, altera status do cartão para suspenso e retorna true/false

	query := fmt.Sprintf("UPDATE physical_cards SET status='suspenso' WHERE user_id = %d;", user_id) //verificando se o user_id já é registrado no banco

	//fmt.Println(query)
	stmt1, err1 := db.Query(query)
	_ = stmt1

	if err1 != nil {
		fmt.Println("Problema na solicitacao")
		return false
	}

	fmt.Println("Cartao suspenso")
	return true
}
