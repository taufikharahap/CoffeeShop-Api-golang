package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func New(db *sqlx.DB) *gin.Engine {
	router := gin.Default()

	users(router, db)
	products(router, db)
	favorites(router, db)
	auth(router, db)

	return router
}
