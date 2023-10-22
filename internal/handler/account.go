package handler

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"vtb_api/internal/entities"
)

func (c *Controller) Info(ctx *gin.Context) {
	id, err := c.getUserId(ctx)
	if err != nil {
		NewErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}
	user, err := c.app.Information(id)
	fmt.Println(user, id)
	if err != nil {
		NewErrorResponse(ctx, http.StatusForbidden, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, user)
}

type InputUser struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (c *Controller) SignIn(ctx *gin.Context) {
	var input InputUser

	if err := ctx.BindJSON(&input); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	token, err := c.app.Auth.GenerateToken(input.Username, input.Password)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}

func (c *Controller) SignUp(ctx *gin.Context) {
	var input entities.User

	if err := ctx.BindJSON(&input); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id, err := c.app.Auth.CreateUser(&input)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (c *Controller) SignOut(ctx *gin.Context) {
	id, err := c.getUserId(ctx)
	if err != nil {
		NewErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, id)
	return
}

func (c *Controller) Update(ctx *gin.Context) {
	id, err := c.getUserId(ctx)
	if err != nil {
		NewErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}
	var input InputUser
	if err := ctx.BindJSON(&input); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	err = c.app.Update(id, input.Username, input.Password)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, input.Username)
	return
}

func (c *Controller) getUserId(ctx *gin.Context) (int, error) {
	c.UserIdentity(ctx)
	id, ok := ctx.Get(userCtx)
	if !ok {
		return 0, errors.New("user id not found")
	}
	return id.(int), nil
}
