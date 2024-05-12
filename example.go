package main

import (
	"errors"
	"log"

	"github.com/gin-gonic/gin"
)

func SayHello(name string, err bool) (string, error) {
	if err {
		return "", errors.New("something wrong")
	}

	return "hello " + name, nil
}

func example() {
	router := gin.Default()

	router.GET("/", exmaple)
	router.GET("/query", queryString)
	router.GET("/params/:username/:hoby", paramString)
	router.POST("/body", reqBody)

	router.Run(":8081")

}

func exmaple(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"status": 200,
		"data":   "hello from gin",
	})

}

type qString struct {
	Limit string `form:"limit"`
	Page  string `form:"page"`
}

// ! http://localhost:8081/query?page=2&limit10
func queryString(ctx *gin.Context) {
	// page := ctx.Query("page")
	// limit := ctx.Query("limit")

	var data qString
	if err := ctx.ShouldBind(&data); err != nil {
		log.Fatal(err)
	}

	ctx.JSON(200, gin.H{
		"page":  data.Page,
		"limit": data.Limit,
	})
}

type pString struct {
	User string `uri:"username"`
	Hoby string `uri:"hoby"`
}

// ! http://localhost:8081/param/ebiebi/bola
func paramString(ctx *gin.Context) {
	// user := ctx.Param("username")
	// hoby := ctx.Param("hoby")

	var data pString
	if err := ctx.ShouldBindUri(&data); err != nil {
		log.Fatal(err)
	}

	ctx.JSON(200, gin.H{
		"user": data.User,
		"hoby": data.Hoby,
	})
}

type User struct {
	Username string `form:"username" json:"username"`
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
}

// ! http://localhost:8081/body
func reqBody(ctx *gin.Context) {

	var data User
	if err := ctx.ShouldBind(&data); err != nil {
		log.Fatal(err)
	}

	ctx.JSON(200, data)
}
