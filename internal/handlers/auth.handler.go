package handlers

import (
	"coffeeshop-api-golang/config"
	"coffeeshop-api-golang/internal/repository"
	"coffeeshop-api-golang/pkg"

	"github.com/gin-gonic/gin"
)

type User struct {
	Email    string `db:"email" json:"email" form:"email"`
	Password string `db:"password" json:"password,omitempty"`
}

type HandlerAuth struct {
	*repository.RepoUsers
}

func NewAuth(r *repository.RepoUsers) *HandlerAuth {
	return &HandlerAuth{r}
}

func (h *HandlerAuth) Login(ctx *gin.Context) {
	var data User

	if err := ctx.ShouldBind(&data); err != nil {
		pkg.NewRes(500, &config.Result{
			Data: err.Error(),
		}).Send(ctx)
		return
	}

	users, err := h.GetAuthData(data.Email)
	if err != nil {
		pkg.NewRes(401, &config.Result{
			Data: err.Error(),
		}).Send(ctx)
		return
	}

	if err := pkg.VerifyPassword(users.Password, data.Password); err != nil {
		pkg.NewRes(401, &config.Result{
			Data: "Password salah",
		}).Send(ctx)
		return
	}

	jwtt := pkg.NewToken(users.User_id, users.Role)
	tokens, err := jwtt.Genrate()
	if err != nil {
		pkg.NewRes(500, &config.Result{
			Data: err.Error(),
		}).Send(ctx)
		return
	}

	pkg.NewRes(200, &config.Result{Data: tokens}).Send(ctx)
}
