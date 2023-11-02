package handler

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"vtb_api/internal/entities"
)

//	@Summary		Information
//	@Security		ApKeyAuth
//	@Tags			Account
//	@Accept			json
//	@Produce		json
//	@Description	Get information current User
//	@Router			/api/Account/Me [get]
func (c *Controller) InfoUser(ctx *gin.Context) {
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

//	@Summary		SignIn
//	@Tags			Account
//	@Accept			json
//	@Produce		json
//	@Description	login
//	@Param			input	body	InputUser	true	"credentials"
//	@Router			/api/Account/SignIn [post]
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

//	@Summary		SignUp
//	@Tags			Account
//	@Accept			json
//	@Produce		json
//	@Param			input	body	entities.User	true	"account info"
//	@Description	Create account
//	@Router			/api/Account/SignUp [post]
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

//	@Summary		SignOut
//	@Security		ApKeyAuth
//	@Tags			Account
//	@Accept			json
//	@Produce		json
//	@Description	Logout
//	@Router			/api/Account/SignOut [post]
func (c *Controller) SignOut(ctx *gin.Context) {
	token, err := c.getToken(ctx)
	if err != nil {
		NewErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}
	c.blackListToken[token] = struct{}{}
	ctx.Set(userCtx, "")
	ctx.JSON(http.StatusOK, "ok")
	return
}

//	@Summary		Update
//	@Security		ApKeyAuth
//	@Tags			Account
//	@Accept			json
//	@Produce		json
//	@Param			input	body	InputUser	true	"Update account"
//	@Description	Edit the user information
//	@Router			/api/Account/Update [put]
func (c *Controller) UpdateUser(ctx *gin.Context) {
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
	err = c.app.Auth.Update(id, input.Username, input.Password)
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
