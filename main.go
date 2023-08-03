package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func showLoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func main() {
	r := gin.Default()
	r.Run()
}
