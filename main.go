package main

import (
	//"fmt"
	//"github.com/Emilybtoliveira/OxeBanking/models"
	//"github.com/Emilybtoliveira/OxeBanking/handlers"
	"fmt"

	"github.com/Emilybtoliveira/OxeBanking/dao"
)

func main() {
	//dao.InitDB()
	//dao.CloseDB()

	dao.CreateDB()
	fmt.Println("*---------------------------------------*")
	defer dao.CloseDB()

	dao.CreateTables()
	fmt.Println("*---------------------------------------*")
	dao.CreateCard(2000, "", "JOSE SILVA JUNIOR", "asdf1234fdsa4321")
	fmt.Println("*---------------------------------------*")
	dao.GetCard(2000)
	fmt.Println("*---------------------------------------*")
	dao.SuspendCard(2000)
	fmt.Println("*---------------------------------------*")
	//dao.CreateCard(2001, "Credito")
	//dao.CreateCard(2002, "Credito")
}
