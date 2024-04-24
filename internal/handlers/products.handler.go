package handlers

import (
	"coffeeshop-api-golang/internal/models"
	"coffeeshop-api-golang/internal/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandlerProducts struct {
	*repository.RepoProducts
}

func NewPruduct(r *repository.RepoProducts) *HandlerProducts {
	return &HandlerProducts{r}
}

func (h *HandlerProducts) GetProductsBy(ctx *gin.Context) {
	var product models.Product

	type qString struct {
		Limit int `form:"limit"`
		Page  int `form:"page"`
	}
	var data qString
	if err := ctx.ShouldBind(&data); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.GetBy(&product, data.Page, data.Limit)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"data": result})

}

func (h *HandlerProducts) PostUser(ctx *gin.Context) {
	var product models.Product

	if err := ctx.ShouldBind(&product); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.CreateProduct(&product)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, result)
}

func (h *HandlerProducts) UpdateProduct(ctx *gin.Context) {
	var product models.Product
	id := ctx.Param("product_id")

	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.UpdateProd(&product, id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, result)

}

func (h *HandlerProducts) DeleteProduct(ctx *gin.Context) {
	var product models.Product

	if err := ctx.ShouldBindUri(&product); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.DeleteProd(&product)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, result)
}
