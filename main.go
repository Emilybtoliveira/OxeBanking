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

	dao.InsertClient(2000, "")
	dao.InsertClient(2001, "Credito")
	dao.InsertClient(2002, "Credito")

	//dao.InsertPhysicalCard(2001, "JOSE SILVA LOPES", "asdd67a8sdaf67a6d8dsa7d8asd67a8sd7a8d6")
}
