package controllers

import (
	"encoding/json"
	"go-ginapi-base/httputil"
	"go-ginapi-base/models"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Get User
func (c *Controller) GetUser(ctx *gin.Context) {
	param_id := ctx.Param("id")
	id, _ := strconv.Atoi(param_id)

	u, err := models.GetUser(id)
	if err != nil {
		httputil.CustomError(ctx, http.StatusNotFound, err)
		return

	}

	ctx.JSON(200, u)
}

// Add User
func (c *Controller) AddUser(ctx *gin.Context) {
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
