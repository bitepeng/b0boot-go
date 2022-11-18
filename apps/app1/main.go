package app

import (
	"b0go/core/engine"
	"embed"
	"io/fs"

	"github.com/gin-gonic/gin"
)

type AppConfig struct {
	Name string
}

var (
	app    *engine.AppConfig
	appId  = "app1"
	config = new(AppConfig)

	//go:embed ui/dist
	uiFS embed.FS
)

func init() {
	uiDist, _ := fs.Sub(uiFS, "ui/dist")
	app = &engine.AppConfig{
		Name:   appId,
		Type:   engine.APP_APP,
		Config: config,
		UIFS:   uiDist,
		Run:    run,
	}
	engine.AppInstall(app)
}

func run() {
	engine.Gin.Use(engine.CorsMiddleware())
	engine.GET(appId, "/ping", "{}", "ping", ping)
	engine.POST(appId, "/ping", "{}", "ping", ping)

}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": appId + " pong " + config.Name,
	})
}
