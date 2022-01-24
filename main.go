package main

import (
	//"fmt"
	//"github.com/Emilybtoliveira/OxeBanking/models"
	//"github.com/Emilybtoliveira/OxeBanking/handlers"
	"github.com/Emilybtoliveira/OxeBanking/dao"
)

func main() {
	//dao.InitDB()
	//dao.CloseDB()

	dao.CreateDB()
	defer dao.CloseDB()

	dao.CreateTables()

	dao.CreateCard(2000, "", "JOSE SILVA JUNIOR", "asdd67a8sdaf67a6d8dsa7d8asd67a8sd7a8d6")
	//dao.CreateCard(2001, "Credito")
	//dao.CreateCard(2002, "Credito")
}
