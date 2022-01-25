package dao

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

//db connection configs
const PostgresDriver = "postgres"
const User = "postgres" //aqui deve-se mudar sempre que o usuario configurado do postgres for outro
const Host = "localhost"
const Port = "5432"
const Password = "123456" //modificar pra senha do seu usuario
const DbName = "teste"
const TableName = "AnyTable"

var db *sql.DB //aponta pra string do banco
var err error  //armazena erros

func InitDB() {
	//responsável por iniciar conexão com o banco de dados
	var DataSourceName = fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable", Host, Port, User, Password, DbName)

	fmt.Println("Dados da conexão: " + DataSourceName)

	db, err = sql.Open(PostgresDriver, DataSourceName)

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Connected!")
	}
}

func CloseDB() {
	//responsável por fechar conexão com o banco de dados
	db.Close()
	fmt.Println("Connection closed!")
}

func CreateDB() {
	//estabelecendo conexão
	conninfo := fmt.Sprintf("user=%s password=%s host=%s sslmode=disable", User, Password, Host)
	db, err := sql.Open(PostgresDriver, conninfo)
	CheckErr(err)

	//checa a existencia do database
	drop_db := "DROP DATABASE IF EXISTS " + DbName
	stmt1, err1 := db.Exec(drop_db)
	CheckErr(err1)

	//criando o database caso nao exista
	create_db := "CREATE DATABASE " + DbName + " WITH OWNER = postgres ENCODING = 'UTF8' TABLESPACE = pg_default CONNECTION LIMIT = -1;"
	stmt1, err2 := db.Exec(create_db)
	CheckErr(err2)

	_ = stmt1 //_ adicionado para nao ocorrer o erro "declared but not used"
	//fmt.Println(stmt1)
}

func CreateTables() {
	InitDB()

	stmt1, err1 := db.Exec("DROP TABLE IF EXISTS public.client;")
	CheckErr(err1)

	stmt1, err2 := db.Exec("DROP TABLE IF EXISTS public.client; CREATE SEQUENCE IF NOT EXISTS client_id_seq START 1; CREATE TABLE IF NOT EXISTS public.client(    id integer NOT NULL DEFAULT nextval('client_id_seq'::regclass),    user_id integer NOT NULL,    card_function character varying(10) COLLATE pg_catalog.\"default\" DEFAULT 'Debito'::character varying, credit_limit integer DEFAULT 0, set_credit_limit integer DEFAULT 0,    CONSTRAINT client_pkey PRIMARY KEY (id, user_id),    CONSTRAINT user_id UNIQUE (user_id))TABLESPACE pg_default;ALTER TABLE IF EXISTS public.client    OWNER to postgres;")
	CheckErr(err2)

	stmt1, err3 := db.Exec("DROP TABLE IF EXISTS public.physical_cards;CREATE TABLE IF NOT EXISTS public.physical_cards(    user_id integer NOT NULL,    card_number bigint NOT NULL,    status character varying(10) COLLATE pg_catalog.\"default\" DEFAULT 'ativo'::character varying,    four_digit_password character varying(100) COLLATE pg_catalog.\"default\" NOT NULL,    owner character varying(30) COLLATE pg_catalog.\"default\" NOT NULL,    valid_thru character varying(10) COLLATE pg_catalog.\"default\" NOT NULL,    cvv integer NOT NULL,    emission_date timestamp(3) with time zone DEFAULT CURRENT_TIMESTAMP,    CONSTRAINT physical_cards_pkey PRIMARY KEY (card_number),    CONSTRAINT user_id FOREIGN KEY (user_id)        REFERENCES public.client (user_id) MATCH SIMPLE        ON UPDATE CASCADE        ON DELETE CASCADE        NOT VALID)TABLESPACE pg_default;ALTER TABLE IF EXISTS public.physical_cards    OWNER to postgres;")
	CheckErr(err3)

	stmt1, err4 := db.Exec("DROP TABLE IF EXISTS public.virtual_cards;CREATE TABLE IF NOT EXISTS public.virtual_cards(    user_id integer NOT NULL,    card_number bigint NOT NULL,    status character varying(10) COLLATE pg_catalog.\"default\" DEFAULT 'ativo'::character varying,    owner character varying(30) COLLATE pg_catalog.\"default\" NOT NULL,    valid_thru character varying(6) COLLATE pg_catalog.\"default\" NOT NULL,    cvv integer NOT NULL,    emission_date timestamp(3) with time zone DEFAULT CURRENT_TIMESTAMP,    nickname character varying(30) COLLATE pg_catalog.\"default\",    CONSTRAINT virtual_cards_pkey PRIMARY KEY (card_number),    CONSTRAINT user_id FOREIGN KEY (user_id)        REFERENCES public.client (user_id) MATCH SIMPLE        ON UPDATE CASCADE        ON DELETE CASCADE)TABLESPACE pg_default;ALTER TABLE IF EXISTS public.virtual_cards    OWNER to postgres;COMMENT ON COLUMN public.virtual_cards.valid_thru    IS 'no formato mm/aa';")
	CheckErr(err4)

	stmt1, err5 := db.Exec("DROP TABLE IF EXISTS public.bill; CREATE SEQUENCE IF NOT EXISTS bill_bill_id_seq START 1; CREATE TABLE IF NOT EXISTS public.bill(		user_id integer NOT NULL,		bill_id integer NOT NULL DEFAULT nextval('bill_bill_id_seq'::regclass),	total_amount real DEFAULT 0.0, amount_paid real DEFAULT 0.0, closing_date date DEFAULT (CURRENT_DATE + 31),	due_date date DEFAULT (CURRENT_DATE + 38), status character varying(10) COLLATE pg_catalog.\"default\" DEFAULT 'aberta'::character varying,	month_year character varying COLLATE pg_catalog.\"default\" DEFAULT to_char(now(), 'MM/YYYY'::text), CONSTRAINT bill_pkey PRIMARY KEY (bill_id), CONSTRAINT user_id FOREIGN KEY (user_id) REFERENCES public.client (user_id) MATCH SIMPLE ON UPDATE NO ACTION ON DELETE NO ACTION	) TABLESPACE pg_default; ALTER TABLE IF EXISTS public.bill OWNER to postgres;")
	CheckErr(err5)

	_ = stmt1
	//fmt.Println(stmt1)
}

func CheckErr(err error) {
	//trata erros
	if err != nil {
		panic(err.Error())
	}
}
