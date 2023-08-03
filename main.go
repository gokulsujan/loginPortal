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
		cookie := &http.Cookie{
			Name:     "username",
			Value:    cred.Username,
			Path:     "/",
			HttpOnly: true,
		}
		http.SetCookie(c.Writer, cookie)
		c.JSON(http.StatusOK, gin.H{"message": "Login Sucessfull"})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Credentials"})
	}
}

func showLoginPage(c *gin.Context) {
	if _, err := c.Request.Cookie("username"); err == nil {
		c.Redirect(http.StatusSeeOther, "/home")
	} else {
		c.HTML(http.StatusOK, "index.html", nil)
	}
}

func main() {
	r := gin.Default()
	r.Static("/static", "/template")
	r.LoadHTMLGlob("template/*")
	r.GET("/", showLoginPage)
	r.POST("/login", loginHandler)
	r.Run()
}
