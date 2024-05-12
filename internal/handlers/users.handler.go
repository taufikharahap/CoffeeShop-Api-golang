package handlers

import (
	"coffeeshop-api-golang/config"
	"coffeeshop-api-golang/internal/models"
	"coffeeshop-api-golang/internal/repository"
	"coffeeshop-api-golang/pkg"
	"fmt"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

type HandlerUsers struct {
	repository.RepoUsersIF
}

func NewUser(r repository.RepoUsersIF) *HandlerUsers {
	return &HandlerUsers{r}
}

func (h *HandlerUsers) GetUserByEmail(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(user.Email)

	result, err := h.GetByEmail(user.Email)
	if err != nil {
		pkg.NewRes(http.StatusBadRequest, &config.Result{
			Data: err.Error(),
		}).Send(ctx)
		return
	}

	// ctx.JSON(200, result)
	pkg.NewRes(200, result).Send(ctx)

}

func (h *HandlerUsers) PostUser(ctx *gin.Context) {
	var err error
	data := models.User{
		Role: "user",
	}

	if err := ctx.ShouldBind(&data); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	_, err = govalidator.ValidateStruct(&data)
	if err != nil {
		pkg.NewRes(401, &config.Result{
			Data: err.Error(),
		}).Send(ctx)
		return
	}

	data.Password, err = pkg.HashPassword(data.Password)
	if err != nil {
		pkg.NewRes(401, &config.Result{
			Data: err.Error(),
		}).Send(ctx)
		return
	}

	result, err := h.CreateUser(&data)
	if err != nil {
		pkg.NewRes(401, &config.Result{
			Data: err.Error(),
		}).Send(ctx)
		return
	}

	pkg.NewRes(200, result).Send(ctx)
}

func (h *HandlerUsers) UpdateUser(ctx *gin.Context) {
	var user models.User
	id := ctx.Param("user_id")
	fmt.Println(id)

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.Update(&user, id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, result)

}
func (h *HandlerUsers) DeleteUser(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBindUri(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.Delete(&user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, result)
}
