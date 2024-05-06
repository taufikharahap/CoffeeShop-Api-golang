package middleware

import (
	"coffeeshop-api-golang/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UploadFile(ctx *gin.Context) {
	file, err := ctx.FormFile("image_url")
	if err != nil {
		if err.Error() == "http: no such file" {
			ctx.Set("image", "")
			ctx.Next()
			return
		}

		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Missing file"})
		return
	}

	// Open the file
	src, err := file.Open()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer src.Close()

	result, err := pkg.CloudInary(src)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to open file"})
		return
	}

	ctx.Set("image", result)
	ctx.Next()
}
