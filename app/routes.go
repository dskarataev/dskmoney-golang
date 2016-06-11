package dskmoney

import "dskmoney-golang/app/handlers"

func (this *App) addRoutes() {
	this.Engine.GET("/", handlers.Index)
	this.Engine.GET("/hello/world", handlers.HelloWorld)
}
