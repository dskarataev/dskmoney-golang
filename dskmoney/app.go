package dskmoney

import (
	"dskmoney-golang/dskmoney/config"
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
	if this.Config.Env == config.ProductionEnv {
		gin.SetMode(gin.ReleaseMode)
	}
	this.Engine = gin.Default()

	// templates
	this.Engine.LoadHTMLGlob("dskmoney/templates/*")

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
