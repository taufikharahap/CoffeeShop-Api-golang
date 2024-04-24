package handlers

import (
	"coffeeshop-api-golang/internal/models"
	"coffeeshop-api-golang/internal/repository"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandlerUsers struct {
	*repository.RepoUsers
}

func NewUser(r *repository.RepoUsers) *HandlerUsers {
	return &HandlerUsers{r}
}

func (h *HandlerUsers) GetUserByEmail(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(user)
	result, err := h.GetByEmail(&user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, result)

}

func (h *HandlerUsers) PostUser(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBind(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.CreateUser(&user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, result)
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
