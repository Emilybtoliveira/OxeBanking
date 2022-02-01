package dao

import (
	"fmt"

	"github.com/Emilybtoliveira/OxeBanking/models"
	_ "github.com/lib/pq"
)

func GetVirtualCard(user_id, card_number int) models.VirtualCard {
	var query string
	var virtual_card models.VirtualCard

	query = fmt.Sprintf("SELECT owner, valid_thru, cvv, nickname FROM public.virtual_cards WHERE user_id = %d and card_number = %d;", user_id, card_number)
	stmt1, err1 := db.Query(query)
	CheckErr(err1)

	query = fmt.Sprintf("SELECT card_function, credit_limit, set_credit_limit FROM client WHERE user_id = %d;", user_id)
	stmt2, err1 := db.Query(query)
	CheckErr(err1)

	if stmt2.Next() != true {
		fmt.Println("Cliente inexistente.")
	} else {
		if stmt1.Next() {
			err1 = stmt1.Scan(&virtual_card.Owner, &virtual_card.Valid_thru, &virtual_card.Cvv, &virtual_card.Nickname)
			CheckErr(err1)

			err1 = stmt2.Scan(&virtual_card.Card_function, &virtual_card.Credit_limit, &virtual_card.Set_credit_limit)
			CheckErr(err1)

			virtual_card.Card_number = card_number
			virtual_card.User_id = user_id
		}
	}

	return virtual_card
}

//Função que recebe o id do usuário e retorna informações de todos os cartões virtuais existentes
func GetAllVirtualCards(user_id int) ([]models.VirtualCard, error) {
	var query string

	query = fmt.Sprintf("SELECT card_number, owner, valid_thru, cvv, nickname FROM public.virtual_cards WHERE user_id = %d and status = 'ativo';", user_id)
	fmt.Println(query)

	stmt1, err1 := db.Query(query)
	CheckErr(err1)

	query = fmt.Sprintf("SELECT card_function, credit_limit, set_credit_limit FROM client WHERE user_id = %d;", user_id)
	stmt2, err1 := db.Query(query)
	CheckErr(err1)

	var virtual_card models.VirtualCard
	var cards_list []models.VirtualCard

	if stmt2.Next() != true {
		fmt.Println("Cliente inexistente.")
	} else {
		//cliente encontrado
		rows := 0

		for stmt1.Next() {
			//fmt.Println("ENTREI")
			err = stmt1.Scan(&virtual_card.Card_number, &virtual_card.Owner, &virtual_card.Valid_thru, &virtual_card.Cvv, &virtual_card.Nickname)
			CheckErr(err)

			err = stmt2.Scan(&virtual_card.Card_function, &virtual_card.Credit_limit, &virtual_card.Set_credit_limit)
			CheckErr(err)

			virtual_card.User_id = user_id

			cards_list = append(cards_list, virtual_card)

			rows += 1
		}

		if rows == 0 {
			//retornando sem cartões
			fmt.Println("Cliente inexistente ou sem cartões virtuais ativos.")
		} else {
			//retornando com cartões
			for i := 0; i < rows; i++ {
				fmt.Println(cards_list[i])
			}

			return cards_list, err1
		}
	}

	return nil, err1
}

//Função de criação de cartão virtual; Recebe informações do usuário e retorna o número do cartão;
func CreateVirtualCard(user_id int, owner string, nickname string) (models.VirtualCard, error) {
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
	}

	//Criando novo cartão
	fmt.Println("Gerando novo cartão virtual...")

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
	query = fmt.Sprintf("INSERT INTO public.virtual_cards(user_id, card_number, owner, valid_thru, cvv, nickname) VALUES (%d, %d, '%s', '%s', %d, '%s');", user_id, card_number, owner, valid_thru, cvv, nickname)
	fmt.Println(query)

	stmt1, err1 := db.Query(query)
	CheckErr(err1)
	_ = stmt1

	virtual_card := GetVirtualCard(user_id, card_number)
	fmt.Println(virtual_card)

	return virtual_card, err1
}

//Função que recebe o id do usuário e o número do cartão virtual a ser removido, retornando true/false
func RemoveVirtualCardByID(user_id, card_number int) (bool, error) {
	var query string

	client_exists := selectClient(user_id)

	if !client_exists {
		fmt.Println("Cliente não existe.")
		return false, nil
	}

	fmt.Println("Cliente existe.")
	query = fmt.Sprintf("UPDATE public.virtual_cards SET status='bloqueado' WHERE user_id = %d and card_number = %d;", user_id, card_number)
	//fmt.Println(query)
	stmt1, err1 := db.Exec(query)

	if err1 != nil {
		fmt.Println("Problema na solicitacao")
		return false, err1
	}

	rows_affected, err1 := stmt1.RowsAffected()
	CheckErr(err1)

	if rows_affected == 0 {
		fmt.Println("Cliente não possui cartão ativo.")
		return false, err1
	}

	fmt.Println("Cartão bloqueado.")
	return true, err1
}
