package dskmoney

import (
	"dskmoney-golang/dskmoney/accounts"
)

func (this *App) addRoutes() {
	this.Engine.GET("/", accounts.AccountsHandler)
	accounts.AddRoutes(this.Engine.Group("/accounts"))
}
