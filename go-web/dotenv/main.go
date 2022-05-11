package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {

	api := gin.Default()
	api.GET("/prueba", PruebaEndpoint)
	api.Run(":8080")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("error al intentar cargar archivo .env")
	}

}

func PruebaEndpoint(c *gin.Context) {

	password := c.GetHeader("password")

	if password != os.Getenv("PASSWORD") {
		fmt.Println("La contraseña es", password)
		c.JSON(401, gin.H{
			"error": "token inválido",
		})
		return
	}

	fmt.Println("Hola Mundo")

}
