package pkg

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func Server(router *gin.Engine) *http.Server {
	var addr string = "0.0.0.0:8081"
	if port := os.Getenv("PORT"); port != "" {
		addr = ":" + port
	}

	server := &http.Server{
		Addr:         addr,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 15,
		Handler:      router,
	}

	return server
}
