package main

import (
	"github.com/bootcamp-go/live-codings/go-web/put/cmd/server/handler"
	"github.com/bootcamp-go/live-codings/go-web/put/internal/products"
	"github.com/gin-gonic/gin"
)

func main() {
	repo := products.NewRepository()
	service := products.NewService(repo)
	p := handler.NewProduct(service)
	r := gin.Default()
	pr := r.Group("/products")
	pr.POST("/", p.Store())
	pr.GET("/", p.GetAll())
	pr.PUT("/:id", p.Update())
	pr.PATCH("/:id", p.UpdateName())
	r.Run()
}
