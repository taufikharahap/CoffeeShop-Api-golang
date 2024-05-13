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

func TestGetByEmail(t *testing.T) {
	data := config.Result{
		Data: map[string]interface{}{
			"user_id": "b458ac8b-1223-43f8-a66b-6668a9d49a06",
			"email":   "testing@mail.com",
			"phone":   "1234556799",
			"birth":   "2000-01-01",
			"gender":  "laki-laki",
		},
	}
	var body = `{
		"email": "testing@mail.com"
	}`

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	w := httptest.NewRecorder()

	handler := NewUser(&repoUserMock)
	exptedResult := &data
	repoUserMock.On("GetByEmail", mock.Anything).Return(exptedResult, nil)

	r.GET("/user", handler.GetUserByEmail)
	req := httptest.NewRequest("GET", "/user", strings.NewReader(body))
	req.Header.Set("Content-type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{
		"status": "OK", 
		"data": {
			"user_id": "b458ac8b-1223-43f8-a66b-6668a9d49a06",
			"email":   "testing@mail.com",
			"phone":   "1234556799",
			"birth":   "2000-01-01",
			"gender":  "laki-laki"
		}
	}`, w.Body.String())
}

func TestPostUser(t *testing.T) {
	var reqBody = `{
		"email": "testing@mail.com",
		"phone": "123456789",
		"password": "abcd1234",
		"role": "user"
	}`
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

func TestUpdateUser(t *testing.T) {
	var reqBody = `{
		"first_name" : "nama",
		"last_name" : "aja",
		"email" : "user@mail.com",
		"phone" : "01234567999",
		"password" : "user",
		"birth" : "2000-01-01",
		"gender" : "laki-laki",
		"image" : "http://db.foto.com"
	}`
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	w := httptest.NewRecorder()

	handler := NewUser(&repoUserMock)
	exptedResult := &config.Result{Message: "1 data user updated"}
	repoUserMock.On("Update", mock.Anything, mock.Anything).Return(exptedResult, nil)

	r.PATCH("/update/:id", handler.UpdateUser)
	req := httptest.NewRequest("PATCH", "/update/b458ac8b-1223-43f8-a66b-6668a9d49a06", strings.NewReader(reqBody))
	req.Header.Set("Content-type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"status": "OK", "description": "1 data user updated"}`, w.Body.String())
}
func TestDeleteUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	w := httptest.NewRecorder()

	handler := NewUser(&repoUserMock)
	exptedResult := &config.Result{Message: "1 data user deleted"}
	repoUserMock.On("Delete", mock.Anything).Return(exptedResult, nil)

	r.DELETE("/delete/:id", handler.DeleteUser)
	req := httptest.NewRequest("DELETE", "/delete/b458ac8b-1223-43f8-a66b-6668a9d49a06", nil)
	req.Header.Set("Content-type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"status": "OK", "description": "1 data user deleted"}`, w.Body.String())
}
