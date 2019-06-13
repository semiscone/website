package main

import (
	"fmt"
	"github.com/go-pg/pg"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"os"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func initDatabase() {

	host := os.Getenv("DATABASE_HOST")
	port := 5432
	user := os.Getenv("DATABASE_USER")
	password := os.Getenv("DATABASE_PASSWORD")
	database := "hummingbird"

	checkAndCreateDB(fmt.Sprintf("%v:%v", host, port), user, password, database)
	connection := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, database)

	log.Infof("connection: %s", connection)
	var err error
	db, err = gorm.Open("postgres", connection)
	if err != nil {
		log.Error(err)
		panic("failed to connect database")
	}
	// defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&User{})
}

func checkAndCreateDB(addr, user, passwd, database string) bool {
	/* connect to the default DB firstly */
	defaultDB := connectDB(addr, user, passwd, "postgres")
	if defaultDB == nil {
		return false
	}

	res, err := defaultDB.Exec(
		fmt.Sprintf("SELECT 1 FROM pg_database WHERE datname = '%s'", database))

	if err != nil {
		log.Error("check target database error")
		return false
	}

	if res.RowsReturned() == 0 {
		// target database is not exist, create it
		log.Infof("target database %s not exists, try to create it ...", database)
		_, err2 := defaultDB.Exec(fmt.Sprintf("CREATE DATABASE %s", database))
		if err2 != nil {
			log.Error("create target database error")
			return false
		}

		log.Info("create target database successfully")
		return true
	}
	return true
}

func connectDB(addr, user, passwd, database string) *pg.DB {
	db := pg.Connect(&pg.Options{
		Addr:     addr,
		User:     user,
		Password: passwd,
		Database: database,
	})
	if db == nil {
		log.Error("Connect to :", database, " failed.")
		return nil
	}

	var n int
	_, err := db.QueryOne(pg.Scan(&n), "SELECT 1")
	if err != nil {
		log.Error("Ping db :", addr, " database:", database, " failed:", err)
		db.Close()
		return nil
	}
	return db
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
