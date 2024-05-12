package handlers

import (
	"coffeeshop-api-golang/config"
	"coffeeshop-api-golang/internal/repository"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var repoUserMock = repository.RepoMock{}
var reqBody = `{
	"email": "testing@mail.com",
	"phone": "123456789",
	"password": "abcd1234",
	"role": "user"
}`

// func TestPostUser(t *testing.T) {
// 	gin.SetMode(gin.TestMode)
// 	r := gin.Default()
// 	w := httptest.NewRecorder()

// 	handler := NewUser(&repoUserMock)
// 	exptedResult := &config.Result{Message: "1 data user created"}
// 	repoUserMock.On("CreateUser", mock.Anything).Return(exptedResult, nil)

// 	r.POST("/create", handler.PostUser)
// 	req := httptest.NewRequest("POST", "/create", strings.NewReader(reqBody))
// 	req.Header.Set("Content-type", "application/json")
// 	r.ServeHTTP(w, req)

//		assert.Equal(t, http.StatusOK, w.Code)
//		assert.JSONEq(t, `{"description": "1 data user created", "status": "OK"}`, w.Body.String())
//	}
func TestPostUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	w := httptest.NewRecorder()

	handler := NewUser(&repoUserMock)
	exptedResult := &config.Result{Message: "1 data user created"}
	repoUserMock.On("CreateUser", mock.Anything).Return(exptedResult, nil)

	r.POST("/create", handler.PostUser)
	req := httptest.NewRequest("POST", "/create", strings.NewReader(reqBody))
	req.Header.Set("Content-type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"description": "1 data user created", "status": "OK"}`, w.Body.String())
}
