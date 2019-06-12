package main

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// User table definition
type User struct {
	gorm.Model
	id                      int `gorm:"AUTO_INCREMENT"` // set num to auto incrementable
	email                   string
	username                string
	roleID                  string
	passwordHash            string
	accessToken             string
	appKey                  string
	appSecret               string
	tokenValid              time.Time
	confirmed               bool
	name                    string
	location                string
	aboutMe                 string
	memberSince             time.Time
	firstSeen               string
	lastseen                string
	avatarHash              string
	posts                   string
	releases                string
	lastAuthFailedTime      time.Time
	lastAuthFailedTimes     int
	lastSyncTime            time.Time
	validTime               time.Time
	points                  int
	phone                   string
	autoDeliver             bool
	syncHistory             bool
	autoDeliverSecuredTrans bool
}

func initDatabase() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&User{})
}
