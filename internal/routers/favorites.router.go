package routers

import (
	"coffeeshop-api-golang/internal/handlers"
	"coffeeshop-api-golang/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func favorites(g *gin.Engine, d *sqlx.DB) {
	route := g.Group("/favorites")

	repo := repository.NewFavorite(d)
	handler := handlers.NewFavorite(repo)

	route.GET("/:user_id", handler.GetFavoritesByUserId)
	route.POST("/", handler.PostFavorite)
	route.PUT("/:favorite_id", handler.UpdateFavorite)
	route.DELETE("/:favorite_id", handler.DeleteFavorite)
}
