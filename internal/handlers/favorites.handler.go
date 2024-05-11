package handlers

import (
	"coffeeshop-api-golang/config"
	"coffeeshop-api-golang/internal/models"
	"coffeeshop-api-golang/internal/repository"
	"coffeeshop-api-golang/pkg"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type HandlerFavorites struct {
	repository.RepoFavoritesIF
}

func NewFavorite(r repository.RepoFavoritesIF) *HandlerFavorites {
	return &HandlerFavorites{r}
}

func (h *HandlerFavorites) GetFavoritesByUserId(ctx *gin.Context) {
	var favorite models.Favorite

	type qString struct {
		Limit int `form:"limit"`
		Page  int `form:"page"`
	}

	var data qString
	if err := ctx.ShouldBind(&data); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctx.ShouldBindUri(&favorite); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.GetByUserId(&favorite, data.Page, data.Limit)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"data": result})

}

func (h *HandlerFavorites) GetFavoriteUserBy(ctx *gin.Context) {
	id := ctx.Param("user_id")
	name := ctx.Query("name")
	page := ctx.DefaultQuery("page", "1")
	limit := ctx.DefaultQuery("limit", "10")

	fmt.Println(id)

	user_id := fmt.Sprintf(`'%s'`, id)

	pg, _ := strconv.Atoi(page)
	lm, _ := strconv.Atoi(limit)

	data, err := h.GetUserIdBy(models.Meta{
		Name:  name,
		Page:  pg,
		Limit: lm,
	}, user_id)

	if err != nil {
		pkg.NewRes(http.StatusBadRequest, &config.Result{
			Data: err.Error(),
		}).Send(ctx)
		fmt.Println(err)
		return
	}

	pkg.NewRes(200, data).Send(ctx)
}

func (h *HandlerFavorites) PostFavorite(ctx *gin.Context) {
	var favorite models.Favorite

	if err := ctx.ShouldBind(&favorite); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.CreateFavorite(&favorite)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, result)
}

func (h *HandlerFavorites) UpdateFavorite(ctx *gin.Context) {
	var favorite models.Favorite
	id := ctx.Param("favorite_id")

	if err := ctx.ShouldBindJSON(&favorite); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.UpdateFav(&favorite, id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, result)

}

func (h *HandlerFavorites) DeleteFavorite(ctx *gin.Context) {
	var favorite models.Favorite

	if err := ctx.ShouldBindUri(&favorite); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.DeleteFav(&favorite)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, result)
}
