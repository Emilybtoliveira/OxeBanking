# **PROJETO OXEBANKING**

**Implementação do microserviço de <span style="color:#6959CD"> cartões </span> utilizando a linguagem GO**

**Dupla:** Emily Brito de Oliveira e José Arthur Lopes Sabino

## Funcionalidades

1. Solicitar novo cartão crédito/débito

2. Bloquear cartão crédito/débito

3. Criar/Remover cartão virtual

4. Fatura (Ajustar limite, Pagar fatura)

   

## Estrutura do Projeto

### Hierarquia dos arquivos e funções

```go
oxebanking_cartao/
	DAO/
		connectionDAO.go
			initDB()
			closeDB()
		cardDAO.go
			GetCard()
			CreateCard()
			suspendCard()
		virtualCardDAO.go
			GetAllVirtualCards()
			CreateVirtualCard()
			RemoveVirtualCardByID()
		billDAO.go
			getBill()
			getCreditLimit()
			adjustCreditLimit()
			deductDebt()
	handlers/
		cardHandler.go
			GetCardHandler()
			CreateCardHandler()
			suspendCardHandler()
		virtualCardHandler.go
			GetAllVirtualCardsHandler()
			CreateVirtualCardHandler()
			RemoveVirtualCardByIDHandler()
		billHandler.go
			getBillHandler()
			getCreditLimitHandler()
			adjustCreditLimitHandler()
			deductDebtHandler()
	models/
		card.go
		virtualCard.go
		bill.go
    database.sql	
    main.go
```





### Detalhamento das funções

### Diretório DAO 

*Lida com o acesso ao banco de dados.*

```go
DAO/
	connectionDAO.go
```

​				**`initDB()`** - responsável por iniciar conexão com o banco de dados

​				**`closeDB()`** - responsável por fechar conexão com o banco de dados

```go
	cardDAO.go
```

​				**`GetCard()`** - recebe o id do usuário e retorna o número do cartão e o tipo (crédito/debito).

​				**`CreateCard()`** - recebe o id do usuário e senha desejada para criação (faz verificação prévia da já existência de um). Retorna o número do cartão.

​				**`SuspendCard()`** - recebe o id do usuário, altera status do cartão para suspenso e retorna true/false.

```go
	virtualCardDAO.go
```

​				**`GetAllVirtualCards()`** - recebe o id do usuário e retorna os números dos cartões virtuais existentes.

​				**`CreateVirtualCard()`** - recebe o id do usuário e retorna as informações geradas do cartão virtual.

​				**`RemoveVirtualCardByID()`** - recebe o id do usuário e o número do cartão virtual a ser removido. Retorna true/false.

```go
	billDAO.go
```

​				**`getBill()`** - recebe id do usuário e retorna valor da fatura.

​				**`getCreditLimit()`** - recebe id do usuário e retorna limite máximo e limite setado pelo usuário.

​				**`adjustCreditLimit()`** - recebe id do usuário e novo limite setado pelo usuário. Retorna true/false.

​				**`deductDebt()`** - recebe id do usuário e valor para abater da fatura. Retorna true/false.	



### Diretório Handler 

*Define como redirecionar as requisições HTTP recebidas.*		

```go
handlers/
	cardHandler.go
```

​				**`GetCardHandler()`** - redireciona para a função `GetCard()` em `DAO/cardDAO.go`.

​				**`CreateCardHandler()`** - redireciona para a função `CreateCard()` em `DAO/cardDAO.go`. 

​				**`suspendCardHandler()`** - redireciona para a função `suspendCard()` em `DAO/cardDAO.go`.		

```go
	virtualCardHandler.go
```

​				**`GetAllVirtualCardsHandler()`** - redireciona para a função `GetAllVirtualCard()` em `DAO/virtualCardDAO.go`.

​				**`CreateVirtualCardHandler()`**- redireciona para a função `CreateVirtualCard()` em `DAO/virtualCardDAO.go`.

​				**`RemoveVirtualCardByIDHandler()`**- redireciona para a função `RemoveVirtualCard()` em `DAO/virtualCardDAO.go`.

```go
	billHandler.go
```

​				**`getBillHandler()`**- redireciona para a função `getBill()` em `DAO/billDAO.go`.

​				**`getCreditLimitHandler()`** - redireciona para a função `getCreditLimit()` em `DAO/billDAO.go`.

​				**`adjustCreditLimitHandler()`** - redireciona para a função `adjustCreditLimit()` em `DAO/billDAO.go`.

​				**`deductDebtHandler()`**  - redireciona para a função `deductDebt()` em `DAO/billDAO.go`.





### Diretório Models 

*Abrigam as classes das entidades do microserviço: cartão, cartão virtual e fatura*	

```go
models/
		card.go
		virtualCard.go
		bill.go
```



### URLs HTTP

### 	Cartão

​	<span style="color:red">**`GET`** </span> `retorna dados do cartão físico`	

```http
/v1/card/{id}
```

​	<span style="color:red">**`POST`** </span>`gerar novo cartão físico`		

```http
/v1/card/
```

​	<span style="color:red">**`PUT`** </span>`atualiza status do cartão físico para bloqueado`	

```http
/v1/card/{id}
```

​	

### 	Cartão Virtual

​	<span style="color:red">**`GET`** </span> `retorna todos os cartões virtuais do usuário`	

```http
/v1/virtualcard/{id}
```

​	<span style="color:red">**`POST`** </span>`gera novo cartão virtual`		

```http
/v1/virtualcard/{id}
```

​	<span style="color:red">**`DELETE`** </span>`remove cartão virtual`	

```http
/v1/virtualcard/{id}
```

### 	

### 	Fatura

​	<span style="color:red">**`GET`** </span> `retorna o valor da fatura`	

```http
/v1/bill/{id}
```

​	<span style="color:red">**`GET`** </span>`retorna o limite total e o limite setado pelo usuário`		

```http
/v1/bill/{id}/limit
```

​	<span style="color:red">**`PUT`** </span>`altera o limite setado pelo usuário`	

```http
/v1/bill/{id}/limit
```

​	<span style="color:red">**`PUT`** </span>`altera o valor da fatura do usuário`	

```http
/v1/bill/{id}
```



## Regras de Negócio

