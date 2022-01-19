package main

import (
  //"fmt"
  //"github.com/Emilybtoliveira/OxeBanking/models"
  //"github.com/Emilybtoliveira/OxeBanking/handlers"
	"github.com/Emilybtoliveira/OxeBanking/dao"
  
)

func main(){
	//dao.InitDB()  
	//dao.CloseDB()
	
	dao.CreateDB()
	//dao.InsertClient(20, "Credit")
	dao.InsertClient(20, "Debit")
	//dao.InsertClient(30, "Test")
	dao.CreateTables()
}
