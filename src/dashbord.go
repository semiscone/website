package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func registerDashboardHandler(r *gin.Engine) {
	r.GET("/dash", dashboard)
}

func dashboard(c *gin.Context) {
	log.Info("dashboard")
	c.HTML(http.StatusOK, "dashboard.html", nil)
}
