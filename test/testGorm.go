package main

import (
	"com.higginslee/IMSystem/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:123456@tcp(localhost:33006)/ginchat?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	//auto create struct
	db.AutoMigrate(models.UserBasic{})
	//create
	user := &models.UserBasic{}
	user.Name = "HigginsLee"
	db.Create(user)
	// Read
	//fmt.Println(db.First(user, 1)) // Find
	//db.Model(user).Update("PassWord", "1234")

}
