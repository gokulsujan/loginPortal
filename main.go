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
		c.Redirect(http.StatusSeeOther, "/home")
	} else {
		c.HTML(http.StatusOK, "index.html", gin.H{"Error": "Invalid Credentials"})
	}
}

func showLoginPage(c *gin.Context) {
	c.Header("Cache-Control", "no-store, must-revalidate")
	if _, err := c.Request.Cookie("username"); err == nil {
		c.Redirect(http.StatusSeeOther, "/home")
	} else {
		c.HTML(http.StatusOK, "index.html", nil)
	}
}

func showHomePage(c *gin.Context) {
	c.Header("Cache-Control", "no-store, must-revalidate")
	if Cookie, err := c.Request.Cookie("username"); err == nil {

		c.HTML(http.StatusAccepted, "home.html", gin.H{"username": Cookie.Value})
	} else {
		c.Redirect(http.StatusSeeOther, "/")
	}
}

func logoutHandler(c *gin.Context) {
	c.Header("Cache-Control", "no-store, must-revalidate")
	if _, err := c.Request.Cookie("username"); err == nil {
		cookie := &http.Cookie{
			Name:     "username",
			Value:    "",
			Path:     "/",
			HttpOnly: true,
			MaxAge:   -1,
		}
		http.SetCookie(c.Writer, cookie)
	}
	c.Redirect(http.StatusSeeOther, "/")
}

func main() {
	r := gin.Default()
	r.Static("/static", "/template")
	r.LoadHTMLGlob("template/*")
	r.GET("/", showLoginPage)
	r.GET("/home", showHomePage)
	r.POST("/login", loginHandler)
	r.GET("/login", showLoginPage)
	r.GET("/logout", logoutHandler)
	r.Run()
}
