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
	dao.CreateTables()
}
