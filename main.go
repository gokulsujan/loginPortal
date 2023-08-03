package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func loginHandler(c *gin.Context) {
	var cred Credentials
	cred.Username = c.PostForm("username")
	cred.Password = c.PostForm("password")

	if cred.Username == "superadmin" && cred.Password == "superpassword" {
		c.JSON(http.StatusOK, gin.H{"message": "Login Sucessfull"})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Credentials"})
	}
}

func showLoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func main() {
	r := gin.Default()
	r.Static("/static", "/template")
	r.LoadHTMLGlob("template/*")
	r.GET("/", showLoginPage)
	r.POST("/login", loginHandler)
	r.Run()
}
