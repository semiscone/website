package main

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"regexp"
	"time"
)

func registerAuthHandler(r *gin.Engine) {
	r.GET("/login", login)
	r.POST("/login", login)
}

func login(c *gin.Context) {
	session := sessions.Default(c)

	if c.Request.Method == "GET" {
		data := gin.H{
			"MsgInfo": session.Flashes("Info"),
			"MsgWarn": session.Flashes("Warn"),
		}
		session.Save()
		c.HTML(http.StatusOK, "login.html", data)
		return
	}

	username := c.PostForm("username")
	password := c.PostForm("password")
	h := md5.New()
	io.WriteString(h, password)
	buffer := bytes.NewBuffer(nil)
	fmt.Fprintf(buffer, "%x", h.Sum(nil))
	inputPass := buffer.String()

	userInfo := getUserInfo(username)
	if userInfo.PasswordHash != inputPass {
		msg := fmt.Sprintf("The password is not matched!")
		session.AddFlash(msg, "Warn")
		log.Info(msg)
		c.Redirect(http.StatusMovedPermanently, "/login")
	}

	userInfo.LastSeen = time.Now().Format("2006-01-02 15:04:05")
	updateUserInfo(userInfo)

	session.Set("uid", userInfo.id)
	session.Set("username", userInfo.Name)
	session.Save()

	c.Redirect(http.StatusMovedPermanently, "/dash")
}

func register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	session := sessions.Default(c)

	usererr := checkUsername(username)
	passerr := checkPassword(password)

	fmt.Println(usererr)
	if usererr == false || passerr == false {
		var msg string
		if usererr == false {
			msg = "Password error, Please to again"
		} else {
			msg = "Password error, Please to again"
		}

		session.AddFlash(msg, "Warn")
		session.Save()
		c.Redirect(302, "/regsiter")
		return
	}

	md5Password := md5.New()
	io.WriteString(md5Password, password)
	buffer := bytes.NewBuffer(nil)
	fmt.Fprintf(buffer, "%x", md5Password.Sum(nil))
	newPass := buffer.String()

	userInfo := getUserInfo(username)

	if userInfo.Username != "" {
		msg := "User already exists"
		session.AddFlash(msg, "Warn")
		session.Save()
		c.Redirect(302, "/register")
		return
	}

	now := time.Now().Format("2006-01-02 15:04:05")

	var users User
	users.Username = username
	users.PasswordHash = newPass
	users.FirstSeen = now
	users.LastSeen = now
	addUser(&users)

	// Login success and set session
	session.Set("uid", userInfo.id)
	session.Set("username", userInfo.Name)
	session.Save()
	c.Redirect(302, "/dash")

}

func checkPassword(password string) (b bool) {
	if ok, _ := regexp.MatchString("^[a-zA-Z0-9]{4,16}$", password); !ok {
		return false
	}
	return true
}

func checkUsername(username string) (b bool) {
	if ok, _ := regexp.MatchString("^[a-zA-Z0-9]{4,16}$", username); !ok {
		return false
	}
	return true
}

// Prepare check session id
func Prepare(c *gin.Context) {
	session := sessions.Default(c)
	sessionID := session.Get("uid")
	sessionUsername := session.Get("username")

	if sessionID == nil {
		c.Redirect(302, "/login")
		return
	}

	log.Infof("User [%v] login", sessionUsername)
	return
}

// AuthRequired to Authenticate User
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		sessionID := session.Get("uid")
		sessionUsername := session.Get("username")
		if sessionID == nil {
			c.Redirect(http.StatusMovedPermanently, "/login")
			return
		}

		log.Infof("User [%v] login", sessionUsername)
		// Continue down the chain to handler etc
		c.Next()
	}
}
