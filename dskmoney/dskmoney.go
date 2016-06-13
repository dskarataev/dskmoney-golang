package dskmoney

import (
	"dskmoney-golang/dskmoney/config"
	"github.com/gin-gonic/gin"

	"gopkg.in/pg.v4"
)

type DSKMoney struct {
	Config *config.Config
	DB     *pg.DB
	Engine *gin.Engine
}

func NewDSKMoney() *DSKMoney {
	return &DSKMoney{}
}

func (this *DSKMoney) Init() error {
	if err := this.initConfig(); err != nil {
		return err
	}

	if err := this.initDB(); err != nil {
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

	return nil
}

func (this *DSKMoney) initConfig() error {
	this.Config = config.NewConfig()
	return this.Config.Init()
}

func (this *DSKMoney) initDB() error {
	db := pg.Connect(&pg.Options{
		Addr:     this.Config.DB.Addr,
		User:     this.Config.DB.User,
		Password: this.Config.DB.Passwd,
		Database: this.Config.DB.Name,
	})
	this.DB = db

	return this.runMigrations()
}

func (this *DSKMoney) Run() error {
	return this.Engine.Run(":" + this.Config.Port)
}
