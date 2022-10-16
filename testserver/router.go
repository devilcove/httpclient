// copywrite 2022 Matthew R Kasun mkasun@nusak.ca
package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var SigningKey []byte

type CustomClaims struct {
	jwt.RegisteredClaims
}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	//basic routes
	router.GET("/ip", GetIP)
	router.POST("/login", Login)
	//authenticated routes
	restricted := router.Group("/api", Auth)
	{
		restricted.GET("/hello", Hello)
	}
	return router
}

func Auth(c *gin.Context) {
	if len(c.Request.Header["Authorization"]) == 0 {
		fail(c, "no auth header")
		return
	}
	id, status := getFromJWT(c.Request.Header["Authorization"][0])
	log.Println(id, status, time.Now())
	if status == 1 {
		fail(c, "token expired")
		return
	}
	if status == 2 {
		fail(c, "invalid token")
	}
	if id != "demo" {
		fail(c, "no such user")
		return
	}

	c.Next()
}

func fail(c *gin.Context, message string) {
	c.JSON(http.StatusUnauthorized, gin.H{"message": message})
	c.Abort()
}

func getFromJWT(auth string) (string, int) {
	//parts := strings.Split(auth, " ")
	log.Println(auth)
	//if len(parts) < 2 {
	//return "", nil
	//}
	token, err := jwt.ParseWithClaims(auth, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return SigningKey, nil
	})
	if err != nil {
		log.Println("error from jwt parse ", err)
		return "", 1
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		log.Println("claims ", claims, "token ", token.Valid)
		return claims.ID, 0
	}
	return "", 2
}
