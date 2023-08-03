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
	if err := c.ShouldBindJSON(&cred); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSOn Data"})
		return
	}

	if cred.Username == "superadmin" && cred.Password == "superpassword" {
		c.JSON(http.StatusOK, gin.H{"message": "Login Sucessfull"})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Credentials"})
	}
}

func showLoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func main() {
	r := gin.Default()
	r.Run()
}
