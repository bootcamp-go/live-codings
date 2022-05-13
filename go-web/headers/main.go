package main

import "github.com/gin-gonic/gin"

var products []request
var lastID int

type request struct {
	ID       int     `json:"id"`
	Nombre   string  `json:"nombre"`
	Tipo     string  `json:"tipo"`
	Cantidad int     `json:"cantidad"`
	Precio   float64 `json:"precio"`
}

func main() {
	r := gin.Default()
	pr := r.Group("/productos")
	pr.POST("/", Guardar())
	r.Run()
}

func Guardar() gin.HandlerFunc {
	return func(c *gin.Context) {

		token := c.GetHeader("token")
		if token != "123456" {
			c.JSON(401, gin.H{
				"error": "token inválido",
			})
			return
		}

		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		lastID++
		req.ID = lastID
		products = append(products, req)

		c.JSON(200, req)
	}
}
