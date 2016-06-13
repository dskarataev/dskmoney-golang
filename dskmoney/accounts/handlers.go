package accounts

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AccountsHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
