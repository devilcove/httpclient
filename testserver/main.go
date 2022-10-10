// copywrite 2022 Matthew R Kasun mkasun@nusak.ca
package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type CustomClaims struct {
	jwt.RegisteredClaims
}

func main() {
	router := SetupRouter()
	router.Run("127.0.0.1:8010")
}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	//basic routes
	router.GET("/ip", GetIP)
	router.POST("/login", Login)
	return router
}

func Login(c *gin.Context) {
	SigningKey := []byte("secretphrase")
	var err error
	request := struct {
		User string
		Pass string
	}{}
	if err := c.BindJSON(&request); err != nil {
		//c.JSON(http.StatusBadRequest, gin.H{"request": c.Request.Body, "error": err.Error()})
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

func GetIP(c *gin.Context) {
	IP := c.ClientIP()
	c.String(http.StatusOK, IP)
}
