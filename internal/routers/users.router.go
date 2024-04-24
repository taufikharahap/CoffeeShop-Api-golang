package routers

import (
	"coffeeshop-api-golang/internal/handlers"
	"coffeeshop-api-golang/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func users(g *gin.Engine, d *sqlx.DB) {
	route := g.Group("/users")

	repo := repository.NewUser(d)
	handler := handlers.NewUser(repo)

	route.GET("/email", handler.GetUserByEmail)
	route.POST("/", handler.PostUser)
	route.PUT("/:user_id", handler.UpdateUser)
	route.DELETE("/:user_id", handler.DeleteUser)
}
