package main

import (
	"gormFirst/internal/createMysqldb"
	"gormFirst/internal/entity"
	"log"
)

func main() {
	db,_:=createMysqldb.InitMYDB()
	//userParam := entity.User{}
	//userParam.ID =0
	//userParam.Name ="name1"
	//userParam.Age = "11"
	//result := db.Create(userParam)
	userParam := entity.User{0,"nn","22"}
	result := db.Create(&userParam)
	log.Println(userParam.ID)
	log.Println("****************")
	log.Println("userID",userParam.ID)
	log.Println(result.RowsAffected)

	//查询
	//type Result struct {
	//	Name string
	//	Age  string
	//}
	//var result Result
	//db.Table("users").Select("name", "age").Where("name = ?", "name1").Scan(&result)
	//log.Println(result.Name)
	//log.Println(result.Age)
}