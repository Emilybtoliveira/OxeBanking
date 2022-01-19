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

	stmt1, err2 := db.Exec("DROP TABLE IF EXISTS public.client;CREATE TABLE IF NOT EXISTS public.client(    id integer NOT NULL,    user_id integer NOT NULL,    function character varying(10) COLLATE pg_catalog.\"default\" DEFAULT 'debito'::character varying,    CONSTRAINT physical_card_pkey PRIMARY KEY (id, user_id),    CONSTRAINT user_id UNIQUE (user_id))TABLESPACE pg_default;ALTER TABLE IF EXISTS public.client    OWNER to postgres;")
	checkErr(err2)

	stmt1, err3 := db.Exec("DROP TABLE IF EXISTS public.registered_cards;CREATE TABLE IF NOT EXISTS public.registered_cards(    user_id integer NOT NULL,    id integer NOT NULL,    card_number bigint NOT NULL,    status character varying(10) COLLATE pg_catalog.\"default\" NOT NULL DEFAULT 'ativo'::character varying,    six_digit_password character(100) COLLATE pg_catalog.\"default\" NOT NULL,    card_owner character varying(50) COLLATE pg_catalog.\"default\" NOT NULL,    valid_thru character varying(6) COLLATE pg_catalog.\"default\" NOT NULL,    cvv integer NOT NULL,    emission_date timestamp(3) with time zone DEFAULT CURRENT_TIMESTAMP,    CONSTRAINT registered_cards_pkey PRIMARY KEY (id),    CONSTRAINT user_id FOREIGN KEY (user_id)        REFERENCES public.client (user_id) MATCH SIMPLE        ON UPDATE CASCADE        ON DELETE CASCADE        NOT VALID)TABLESPACE pg_default;ALTER TABLE IF EXISTS public.registered_cards    OWNER to postgres;")
	checkErr(err3)

	stmt1, err4 := db.Exec("DROP TABLE IF EXISTS public.registered_virtual_cards;CREATE TABLE IF NOT EXISTS public.registered_virtual_cards(    user_id integer,    id integer NOT NULL,    card_number bigint NOT NULL,    status character varying(10) COLLATE pg_catalog.\"default\" DEFAULT 'ativo'::character varying,    card_owner character varying(50) COLLATE pg_catalog.\"default\" NOT NULL,    valid_thru character varying(6) COLLATE pg_catalog.\"default\" NOT NULL,    cvv integer NOT NULL,    emission_date timestamp(3) with time zone DEFAULT CURRENT_TIMESTAMP,    card_nickname character varying(30) COLLATE pg_catalog.\"default\" NOT NULL,    CONSTRAINT registered_virtual_cards_pkey PRIMARY KEY (id),    CONSTRAINT user_id FOREIGN KEY (user_id)        REFERENCES public.client (user_id) MATCH SIMPLE        ON UPDATE NO ACTION        ON DELETE NO ACTION)TABLESPACE pg_default;ALTER TABLE IF EXISTS public.registered_virtual_cards    OWNER to postgres;")
	checkErr(err4)


	fmt.Println(stmt1)

}

func checkErr(err error) {
    //trata erros
    if err != nil {
        panic(err.Error())
    }
}
