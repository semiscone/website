package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"os"
	"time"
)

var db *gorm.DB

func initDatabase() {
	log.Print("Sleep for database")

	time.Sleep(10 * time.Second)

	host := os.Getenv("POSTGRES_HOST")
	port := 5432
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	database := os.Getenv("POSTGRES_DATABASE")

	conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, database)
	log.Printf("conn: %s", conn)

	var err error
	db, err = gorm.Open("postgres", conn)
	if err != nil {
		log.Fatal(err)
		panic("failed to connect database")
	}
	//defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&User{})

}

func getUserInfo(username string) *User {
	var user User
	db.Where("name = ?", username).First(&user)
	return &user
}

func updateUserInfo(user *User) error {
	err := db.Save(&user).Error
	return err
}

func addUser(user *User) error {
	err := db.Create(user).Error
	return err
}
