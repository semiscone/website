package main

import (
	"time"
  "github.com/jinzhu/gorm"
)

// User table definition
type User struct {
	gorm.Model
	id                      int `gorm:"AUTO_INCREMENT"` // set num to auto incrementable
	Email                   string
	Username                string
	RoleID                  string
	PasswordHash            string
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
