# **Projeto OxeBanking**

Implementação do microserviço de cartões utilizando a linguagem GO

**Dupla:** [Emily Brito de Oliveira](https://github.com/Emilybtoliveira) e [José Arthur Lopes Sabino]()

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
			getCard()
			createCard()
			suspendCard()
		virtualCardDAO.go
			getAllVirtualCards()
			createVirtualCard()
			RemoveVirtualCardByID()
		billDAO.go
			getBill()
			getCreditLimit()
			adjustCreditLimit()
			deductDebt()
	handlers/
		cardHandler.go
			getCardHandler()
			createCardHandler()
			suspendCardHandler()
		virtualCardHandler.go
			getAllVirtualCardsHandler()
			createVirtualCardHandler()
			removeVirtualCardByIDHandler()
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

<br></br>

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

​				**`GetCard()`** - recebe o id do usuário e retorna algumas informações do cartão.

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

<br></br>

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

<br></br>

### Diretório Models 

*Abrigam as classes das entidades do microserviço: cartão, cartão virtual e fatura*	

```go
models/
		card.go
		virtualCard.go
		bill.go
```

---
<br></br>

## URLs HTTP

### 	Cartão

​	<span style="color:red">**`GET`** </span> `retorna dados do cartão físico`	

```http
/v1/card/{id}
```

​	<span style="color:red">**`POST`** </span>`gerar novo cartão físico`		

```http
/v1/card/
```

​	<span style="color:red">**`PUT`** </span>`atualiza status ou a função do cartão físico`	

```http
/v1/card/{id}
```

<br></br>
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
<br></br>
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

<br></br>
## Regras de Negócio

### Cartões

​	**`getCard`** 

1. A solicitação de informações do cartão físico do usuário, retorna os seguintes dados: *os ultimos 4 dígitos, titular do cartão e a função (crédito/débito).*

​	**`createCard`** 

  1. A solicitação de gerar um novo cartão físico só é efetivada após verificação da já existência de um cartão físico *não bloqueado*.
  2. Caso o mesmo esteja bloqueado, é criado um novo.
  3. Sempre que um novo usuário é registrado, `createCard` deve ser utilizado para gerar a função *débito*.
  4. `createCard` também será usado para atualizar a função do cartão de crédito para débito.

​	**`suspendCard`** 

1. Um cartão bloqueado não pode ser desbloqueado.
2. O histórico de cartões do usuário é mantido no banco (bloqueados e ativo).
3. Somente um cartão (último solicitado) poderá estar com status `ativo` .

<br></br>
### Cartões virtuais

​	**`GetAllVirtualCards`** 

1.  A solicitação de retorno de todos os cartões virtuais do usuário retorna: *o número completo do cartão, o nome do titular, o código de segurança, a validade e a função*.
2. Apenas os cartões virtuais com status `ativo` são retornados.

​		**`CreateVirtualCard`**

1. Um usuário pode ter vários cartões virtuais, portanto não há nenhum tipo de verificação prévia.

​		**`RemoveVirtualCardByID`** 

1. A solicitação de remover um cartão virtual apenas altera o status do cartão para bloqueado.
2. A mudança de status impede que o mesmo seja novamente acessado pelo usuário (não é retornado).
3. O histórico dos cartões virtuais (bloqueados e ativos) é mantido no banco de dados.

<br></br>

### Fatura

**`getBill`** 

1. A solicitação de retorno da fatura, retorna: *valor da fatura pendente de pagamento, a data de fechamento, a data de vencimento e o status (aberta/fechada)*.
2. **Até então, não será implementada a função de retornar todo o histórico de faturas do usuário.**

​	**`getCreditLimit`** 

1. A solicitação de retorno do limite de crédito retorna: *o valor atual setado pelo usuário, o valor total de limite que ele possui e o quanto do limite já está sendo usado*.

​	**`adjustCreditLimit`** 

1. O ajuste do limite só será efetivado sob verificação de que novo limite está no intervalo do limite total que o usuário possui.
2. Caso o usuário sete um novo limite com um valor menor do que o que ele já usou na fatura atual, esse limite só será válido para a proxima fatura.

​	**`deductDebt`**

1. O abate no valor da fatura é feito após verificação de que a data de recebimento do valor devido está dentro do vencimento da fatura.

