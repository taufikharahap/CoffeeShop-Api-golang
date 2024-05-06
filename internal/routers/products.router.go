package routers

import (
	"coffeeshop-api-golang/internal/handlers"
	"coffeeshop-api-golang/internal/middleware"
	"coffeeshop-api-golang/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func products(g *gin.Engine, d *sqlx.DB) {
	route := g.Group("/products")

	repo := repository.NewPruduct(d)
	handler := handlers.NewPruduct(repo)

	route.GET("/query", middleware.AuthJwt("admin", "user"), handler.GetProductsBy)
	route.GET("/", middleware.AuthJwt("admin", "user"), handler.GetProducts)
	route.POST("/", middleware.AuthJwt("admin"), middleware.UploadFile, handler.PostProduct)
	route.PATCH("/:product_id", middleware.AuthJwt("admin"), middleware.UploadFile, handler.UpdateProduct)
	route.DELETE("/:product_id", middleware.AuthJwt("admin"), handler.DeleteProduct)
}
