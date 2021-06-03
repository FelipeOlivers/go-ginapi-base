package controllers

import (
	"encoding/json"
	"fmt"
	"go-ginapi-base/httputil"
	"go-ginapi-base/models"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Create User
func (c *Controller) CreateUser(ctx *gin.Context) {
	body, _ := ioutil.ReadAll(ctx.Request.Body)
	var user models.User

	err := json.Unmarshal(body, &user)
	if err != nil {
		httputil.CustomError(ctx, http.StatusBadRequest, err)
		return

	}
	user.Email = "folivers@gmail.com"
	ctx.JSON(200, user)
}

// Get User
func (c *Controller) GetUser(ctx *gin.Context) {
	param_id := ctx.Param("id")
	id, err := strconv.Atoi(param_id)

	if err != nil {
		httputil.CustomError(ctx, http.StatusBadRequest, err)
		return

	}
	if id < 1 {
		httputil.CustomError(ctx, http.StatusNotFound, fmt.Errorf("User id = %d is not found.", id))
		return
	}

	u := models.User{
		ID:       id,
		Name:     "Luiz Filipe",
		Lastname: "Miranda de Oliveira",
	}

	ctx.JSON(200, u)
}
