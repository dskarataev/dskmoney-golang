package dskmoney

import (
	"dskmoney-golang/app/config"
	"github.com/gin-gonic/gin"
)

type App struct {
	Engine *gin.Engine
	Config     *config.Config
}

func NewApp() *App {
	return &App{}
}

func (this *App) Init() error {
	err := this.initConfig()
	if err != nil {
		return err
	}

	// router and HTTP server
	this.Engine = gin.Default()

	// templates
	this.Engine.LoadHTMLGlob("app/templates/*")

	// routes
	this.addRoutes()

	return err
}

func (this *App) initConfig() error {
	this.Config = config.NewConfig()
	return this.Config.Init()
}

func (this *App) Run() error {
	return this.Engine.Run(":" + this.Config.Port)
}
