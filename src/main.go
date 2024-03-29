package main

import (
	"net/http"
	"path/filepath"

	log "github.com/sirupsen/logrus"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("sessions", store))

	r.Use(gin.Logger())
	r.Delims("{{", "}}")

	r.Static("alithon", "./static/alithon")
	r.Static("dist", "./static/dist")
	r.Static("js", "./static/js")
	r.Static("plugins", "./static/plugins")

	r.HTMLRender = loadTemplates("./templates")

	r.GET("/", index)
	r.GET("/index.html", index)
	r.GET("/basic.html", basic)
	r.GET("/standard.html", standard)
	r.GET("/wide.html", wide)
	r.GET("/sidebar.html", sidebar)

	registerAuthHandler(r)
	registerDashboardHandler(r)
	return r
}

func loadTemplates(templatesDir string) multitemplate.Renderer {
	log.Printf("loadTemplates")
	r := multitemplate.NewRenderer()

	layouts, err := filepath.Glob(templatesDir + "/layout/*.html")
	if err != nil {
		log.Print(err.Error())
	}

	includes, err := filepath.Glob(templatesDir + "/includes/*.html")
	if err != nil {
		log.Print(err.Error())
	}

	// Generate our templates map from our layouts/ and includes/ directories
	for _, include := range includes {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)
		r.AddFromFiles(filepath.Base(include), files...)
	}

	pages, err := filepath.Glob(templatesDir + "/*.html")
	if err != nil {
		panic(err.Error())
	}

	// Generate our templates map from our layouts/ and includes/ directories
	for _, page := range pages {
		r.AddFromFiles(filepath.Base(page), page)
	}

	return r
}

func main() {

	router := setupRouter()

	initDatabase()

	log.Printf("HTTP Server Launching...")
	ret := router.Run(":5000")
	log.Errorf("Failed reasion: %s", ret)
}

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func basic(c *gin.Context) {
	c.HTML(http.StatusOK, "basic.html", nil)
}

func standard(c *gin.Context) {
	c.HTML(http.StatusOK, "standard.html", nil)
}

func wide(c *gin.Context) {
	c.HTML(http.StatusOK, "wide.html", nil)
}

func sidebar(c *gin.Context) {
	c.HTML(http.StatusOK, "sidebar.html", nil)
}
