package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vtb_api/internal/entities"
)

func (c *Controller) Info(*gin.Context) {
	return
}

func (c *Controller) SignIn(*gin.Context) {
	return
}

func (c *Controller) SignUp(ctx *gin.Context) {
	var input entities.User

	if err := ctx.BindJSON(&input); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id, err := c.app.Auth.CreateUser(input)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (c *Controller) SignOut(*gin.Context) {
	return
}

func (c *Controller) Update(*gin.Context) {
	return
}
