package routers

import (
	"coffeeshop-api-golang/internal/handlers"
	"coffeeshop-api-golang/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func products(g *gin.Engine, d *sqlx.DB) {
	route := g.Group("/products")

	repo := repository.NewPruduct(d)
	handler := handlers.NewPruduct(repo)

	route.GET("/query", handler.GetProductsBy)
	route.POST("/", handler.PostUser)
	route.PUT("/:product_id", handler.UpdateProduct)
	route.DELETE("/:product_id", handler.DeleteProduct)
}
