package main

import (
	"fmt"

	docs "github.com/bootcamp-go/live-codings/go-web/swagger/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           API de ejemplo
// @version         1.0
// @description     Esta es una pequeña app para mostrar como funciona la documentación.
// @termsOfService  https://www.digitalhouse.com

// @contact.name   Bootcamp Go
// @contact.url    https://www.digitalhouse.com
// @contact.email  bootcamp-go@digitalhouse.com

// @host      localhost:8080
func main() {
	api := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	api.GET("/prueba", PruebaMiddleware(), PruebaEndpoint)
	api.Run(":8080")
}

// PruebaEndpoint godoc
// @Summary      Hola Mundo
// @Description  Devuelve un Hola Mundo
// @Accept       json
// @Produce      json
// @Success      200  {object}  map[string]string
// @Router       /prueba [get]
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
