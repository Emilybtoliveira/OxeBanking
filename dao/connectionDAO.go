package dao

import (
	"fmt"
	"database/sql"
	_"github.com/lib/pq"
)

//db connection configs
const PostgresDriver = "postgres"
const User = "postgres"
const Host = "localhost"
const Port = "5432"
const Password = "654321"
const DbName = "database"
const TableName = "AnyTable"

var db *sql.DB //aponta pra string do banco
var err error //armazena erros

func InitDB(){
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

func CloseDB(){
	//responsável por fechar conexão com o banco de dados
	db.Close()
}

func CreateDB(){
	//estabelecendo conexão
	conninfo := fmt.Sprintf("user=%s password=%s host=%s sslmode=disable", User, Password, Host)
    db, err := sql.Open(PostgresDriver, conninfo)
	checkErr(err)
	
	//fazendo a limpa
	stmt1, err1 := db.Exec("DROP DATABASE IF EXISTS database");
	checkErr(err1)

	//criando a database
	stmt1, err2 := db.Exec("CREATE DATABASE database WITH OWNER = postgres ENCODING = 'UTF8' TABLESPACE = pg_default CONNECTION LIMIT = -1; ")
	checkErr(err2)

	fmt.Println(stmt1)

}

func CreateTables(){
	InitDB()
	
	stmt1, err1 := db.Exec("DROP TABLE IF EXISTS public.client;")
	checkErr(err1)

	stmt1, err2 := db.Exec("DROP TABLE IF EXISTS public.client; CREATE SEQUENCE IF NOT EXISTS client_id_seq START 1; CREATE TABLE IF NOT EXISTS public.client(    id integer NOT NULL DEFAULT nextval('client_id_seq'::regclass),    user_id integer NOT NULL,    card_function character varying(10) COLLATE pg_catalog.\"default\" DEFAULT 'debito'::character varying,    CONSTRAINT client_pkey PRIMARY KEY (id, user_id),    CONSTRAINT user_id UNIQUE (user_id))TABLESPACE pg_default;ALTER TABLE IF EXISTS public.client    OWNER to postgres;")
	checkErr(err2)

	stmt1, err3 := db.Exec("DROP TABLE IF EXISTS public.physical_cards;CREATE TABLE IF NOT EXISTS public.physical_cards(    user_id integer NOT NULL,    card_number bigint NOT NULL,    status character varying(10) COLLATE pg_catalog.\"default\" DEFAULT 'ativo'::character varying,    four_digit_password character varying(100) COLLATE pg_catalog.\"default\" NOT NULL,    owner character varying(30) COLLATE pg_catalog.\"default\" NOT NULL,    valid_thru character varying(6) COLLATE pg_catalog.\"default\" NOT NULL,    cvv integer NOT NULL,    emission_date timestamp(3) with time zone DEFAULT CURRENT_TIMESTAMP,    CONSTRAINT physical_cards_pkey PRIMARY KEY (card_number),    CONSTRAINT user_id FOREIGN KEY (user_id)        REFERENCES public.client (user_id) MATCH SIMPLE        ON UPDATE CASCADE        ON DELETE CASCADE        NOT VALID)TABLESPACE pg_default;ALTER TABLE IF EXISTS public.physical_cards    OWNER to postgres;")
	checkErr(err3)

	stmt1, err4 := db.Exec("DROP TABLE IF EXISTS public.virtual_cards;CREATE TABLE IF NOT EXISTS public.virtual_cards(    user_id integer NOT NULL,    card_number bigint NOT NULL,    status character varying(10) COLLATE pg_catalog.\"default\" DEFAULT 'ativo'::character varying,    owner character varying(30) COLLATE pg_catalog.\"default\" NOT NULL,    valid_thru character varying(6) COLLATE pg_catalog.\"default\" NOT NULL,    cvv integer NOT NULL,    emission_date timestamp(3) with time zone DEFAULT CURRENT_TIMESTAMP,    nickname character varying(30) COLLATE pg_catalog.\"default\",    CONSTRAINT virtual_cards_pkey PRIMARY KEY (card_number),    CONSTRAINT user_id FOREIGN KEY (user_id)        REFERENCES public.client (user_id) MATCH SIMPLE        ON UPDATE CASCADE        ON DELETE CASCADE)TABLESPACE pg_default;ALTER TABLE IF EXISTS public.virtual_cards    OWNER to postgres;COMMENT ON COLUMN public.virtual_cards.valid_thru    IS 'no formato mm/aa';")
	checkErr(err4)

	fmt.Println(stmt1)
}

func InsertClient(user_id int64, card_function string){
	InitDB()
	//esse foi meu codigo
	//query := fmt.Sprintf("INSERT INTO public.client(user_id, card_function) VALUES(%d, '%s')", user_id, card_function)
	//fmt.Println(query)
	
	//stmt1, err1 := db.Exec(query)
	//checkErr(err1) 
	
	//fmt.Println(stmt1)
	
	//esse eh a base do vapordev
	sqlStatement := fmt.Sprintf("INSERT INTO public.client VALUES ($2, $3)")

    insert, err := db.Prepare(sqlStatement)
    checkErr(err)

    result, err := insert.Exec(5, "Debit")
    checkErr(err)

    affect, err := result.RowsAffected()
    checkErr(err)

    fmt.Println(affect)
	
	CloseDB()
}

func checkErr(err error) {
    //trata erros
    if err != nil {
        panic(err.Error())
    }
}
