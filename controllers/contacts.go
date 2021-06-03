package controllers

import (
	"go-ginapi-base/httputil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (c *Controller) GetContacts(ctx *gin.Context) {
	param_id := ctx.Param("id")
	id, err := strconv.Atoi(param_id)

	if err != nil {
		httputil.CustomError(ctx, http.StatusBadRequest, err)
		return

	}
	ctx.JSON(200, gin.H{"contact_id": id})
}
