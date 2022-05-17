package main

import (
	"github.com/bootcamp-go/live-codings/go-web/functional-tests/cmd/server/handler"
	"github.com/bootcamp-go/live-codings/go-web/functional-tests/internal/products"
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
	r.Run()
}
