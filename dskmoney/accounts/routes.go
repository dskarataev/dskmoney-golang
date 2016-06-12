package accounts

import (
	"github.com/gin-gonic/gin"
)

func AddRoutes(router *gin.RouterGroup) {
	router.GET("/", AccountsHandler)
}
