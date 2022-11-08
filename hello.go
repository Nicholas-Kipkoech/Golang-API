package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "welcome to Golang-API for starter",
			"username": "Nicholas kipkoech",
			"country": "Kenya",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}