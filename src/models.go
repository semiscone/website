package main

import (
	"github.com/jinzhu/gorm"
	"time"
)

// User table definition
type User struct {
	gorm.Model
	id                      int `gorm:"AUTO_INCREMENT"` // set num to auto incrementable
	Email                   string
	Username                string
	RoleID                  string
	PasswordHash            string
	AccessToken             string
	AppKey                  string
	AppSecret               string
	TokenValid              time.Time
	Confirmed               bool
	Name                    string
	Location                string
	AboutMe                 string
	memberSince             time.Time
	FirstSeen               string
	LastSeen                string
	AvatarHash              string
	Posts                   string
	Releases                string
	LastAuthFailedTime      time.Time
	LastAuthFailedTimes     int
	LastSyncTime            time.Time
	ValidTime               time.Time
	Points                  int
	Phone                   string
	AutoDeliver             bool
	SyncHistory             bool
	AutoDeliverSecuredTrans bool
}
