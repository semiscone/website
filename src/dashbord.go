package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func registerDashboardHandler(r *gin.Engine) {
	private := r.Group("/private")
	{
		private.GET("/dash", dashboard)
		private.GET("/dashboard", dashboard)
	}
	private.Use(AuthRequired())

	r.GET("/d", dashboard)
}

func dashboard(c *gin.Context) {
	log.Info("dashboard")
	c.HTML(http.StatusOK, "dashboard.html", nil)
}
