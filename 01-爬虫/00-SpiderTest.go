package main

import (
	"testProject/01-爬虫/db"
	"testProject/01-爬虫/models"
)

func main() {
	//ktmlInfo := models.DasKtdmInfo{
	//	ID: 1,
	//	URL: "fffff",
	//	Title: "test",
	//	CreateTime: time.Now(),
	//}
	ktmlInfo := models.DasKtdmInfo{
		Url: "fffff",
		Title: "test",
	}

	engin := db.InitMysqlEngin()
	engin.Insert(ktmlInfo)
}
