package database

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"dcs/player"
)

func Connect() *gorm.DB{
	dsn := "host=localhost user=postgres password=dadon2004 dbname=assignmentGo port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connected to DB")
	db.AutoMigrate(&player.Player{})
	return db
}

