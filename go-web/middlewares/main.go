package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	api := gin.Default()
	api.GET("/prueba", PruebaMiddleware(), PruebaEndpoint)
	api.Run(":8080")
}

func PruebaEndpoint(c *gin.Context) {
	resp := map[string]string{"Hola": "Mundo"}
	c.JSON(200, resp)
}

func PruebaMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Esto es un Middleware y podemos poner logica")
		c.Next()
	}
}
