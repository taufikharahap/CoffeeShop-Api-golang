package middleware

import (
	"coffeeshop-api-golang/config"
	"coffeeshop-api-golang/pkg"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthJwt(role ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var valid bool
		var header string

		if header = ctx.GetHeader("Authorization"); header == "" {
			pkg.NewRes(401, &config.Result{
				Data: "Silahkan login",
			}).Send(ctx)
			return
		}

		if !strings.Contains(header, "Bearer") {
			pkg.NewRes(401, &config.Result{
				Data: "Invalid Header Type",
			}).Send(ctx)
			return
		}

		tokens := strings.Replace(header, "Bearer ", "", -1)
		check, err := pkg.VerifyToken(tokens)
		if err != nil {
			pkg.NewRes(401, &config.Result{
				Data: err.Error(),
			}).Send(ctx)
			return
		}

		for _, r := range role {
			if r == check.Role {
				valid = true
			}
		}

		if !valid {
			pkg.NewRes(401, &config.Result{
				Data: "You not have permission",
			}).Send(ctx)
			return
		}

		ctx.Set("userId", check.Id)
		ctx.Next()
	}
}
