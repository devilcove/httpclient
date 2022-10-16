// copywrite 2022 Matthew R Kasun mkasun@nusak.ca
package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func GetIP(c *gin.Context) {
	IP := c.ClientIP()
	c.String(http.StatusOK, IP)
}

func Login(c *gin.Context) {
	log.Println("login")
	var err error
	request := struct {
		User string
		Pass string
	}{}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"request": c.Request.Body, "error": err.Error()})
		return
	}
	if request.User != "demo" || request.Pass != "pass" {
		c.JSON(http.StatusBadRequest, gin.H{"request": request, "error": "invalid username or password"})
		return
	}
	expires := time.Now().Add(time.Minute * 3)
	claims := CustomClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expires),
			Issuer:    "nusak.ca",
			ID:        request.User,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	JWT, err := token.SignedString(SigningKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error creating JWT", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"JWT": JWT})
}

func Hello(c *gin.Context) {
	IP := c.ClientIP()
	c.String(http.StatusOK, IP)
}
