package handlers

import (
	"coffeeshop-api-golang/config"
	"coffeeshop-api-golang/internal/models"
	"coffeeshop-api-golang/internal/repository"
	"coffeeshop-api-golang/pkg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type HandlerProducts struct {
	*repository.RepoProducts
}

func NewPruduct(r *repository.RepoProducts) *HandlerProducts {
	return &HandlerProducts{r}
}

func (h *HandlerProducts) GetProducts(ctx *gin.Context) {
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

	result, err := h.GetProd(&product, data.Page, data.Limit)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"data": result})

}

func (h *HandlerProducts) GetProductsBy(ctx *gin.Context) {
	name := ctx.Query("name")
	page := ctx.DefaultQuery("page", "1")
	limit := ctx.DefaultQuery("limit", "10")

	pg, _ := strconv.Atoi(page)
	lm, _ := strconv.Atoi(limit)

	data, err := h.GetProdBy(models.Meta{
		Name:  name,
		Page:  pg,
		Limit: lm,
	})

	if err != nil {
		pkg.NewRes(http.StatusBadRequest, &config.Result{
			Data: err.Error(),
		}).Send(ctx)
		return
	}

	pkg.NewRes(200, data).Send(ctx)
}

func (h *HandlerProducts) PostProduct(ctx *gin.Context) {
	product := models.Product{}

	if err := ctx.ShouldBind(&product); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	product.Image_url = ctx.MustGet("image").(string)
	result, err := h.CreateProd(&product)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	pkg.NewRes(200, result).Send(ctx)
}

func (h *HandlerProducts) UpdateProduct(ctx *gin.Context) {
	product := models.Product{}

	id := ctx.Param("product_id")

	if err := ctx.ShouldBind(&product); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product.Image_url = ctx.MustGet("image").(string)
	result, err := h.UpdateProd(&product, id)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	pkg.NewRes(200, result).Send(ctx)

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

	pkg.NewRes(200, result).Send(ctx)
}
