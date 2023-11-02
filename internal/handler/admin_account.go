package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"vtb_api/internal/entities"
)

func (c *Controller) isAdmin(ctx *gin.Context) error {
	id, err := c.getUserId(ctx)
	if err != nil {
		NewErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return err
	}
	user, err := c.app.Auth.Information(id)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return err
	}
	if user.IsAdmin == false {
		NewErrorResponse(ctx, http.StatusForbidden, "User not admin")
		return err
	}
	return nil
}

//	@Summary		Accounts
//	@Security		ApKeyAuth
//	@Tags			Admin
//	@Accept			json
//	@Produce		json
//	@Description	Get the list of users
//	@Router			/api/Admin/Account [get]
func (c *Controller) ListAccounts(ctx *gin.Context) {
	err := c.isAdmin(ctx)
	if err != nil {
		return
	}
	st, co := ctx.Query("start"), ctx.Query("count")
	start, count, err := handleParameters(st, co)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	accounts, err := c.app.Admin.GetAccounts(start, count)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, accounts)
}

//	@Summary		GetAccount
//	@Security		ApKeyAuth
//	@Tags			Admin
//	@Accept			json
//	@Produce		json
//	@Param			id	path	int	true	"id of user"
//	@Description	Get the user by id
//	@Router			/api/Admin/Account/:id [get]
func (c *Controller) GetAccountById(ctx *gin.Context) {
	err := c.isAdmin(ctx)
	if err != nil {
		return
	}
	ID := ctx.Param("id")
	id, err := strconv.Atoi(ID)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	account, err := c.app.Admin.GetAccountById(id)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, account)
}

//	@Summary		CreateUser
//	@Security		ApKeyAuth
//	@Tags			Admin
//	@Accept			json
//	@Produce		json
//	@Description	Admin create the user
//	@Router			/api/Admin/Account [post]
func (c *Controller) CreateAccountByAdmin(ctx *gin.Context) {
	err := c.isAdmin(ctx)
	if err != nil {
		return
	}
	var user entities.User
	if err := ctx.BindJSON(&user); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	account, err := c.app.Admin.CreateAccount(&user)
	if err != nil {
		return
	}
	ctx.JSON(http.StatusOK, account)
}

//	@Summary		UpdateAccount
//	@Security		ApKeyAuth
//	@Tags			Admin
//	@Accept			json
//	@Produce		json
//	@Param			id	path	int	true	"id of user"
//	@Description	Admin update the user information
//	@Router			/api/Admin/Account [put]
func (c *Controller) UpdateAccount(ctx *gin.Context) {
	err := c.isAdmin(ctx)
	if err != nil {
		return
	}
	ID := ctx.Param("id")
	id, err := strconv.Atoi(ID)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	var user entities.User
	if err := ctx.BindJSON(&user); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	user.Id = id
	err = c.app.Admin.UpdateAccount(&user)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, "ok")
}

//	@Summary		DeleteUser
//	@Security		ApKeyAuth
//	@Tags			Admin
//	@Accept			json
//	@Produce		json
//	@Param			id	path	int	true	"id of user"
//	@Description	Admin remove the user
//	@Router			/api/Admin/Account [delete]
func (c *Controller) DeleteAccount(ctx *gin.Context) {
	err := c.isAdmin(ctx)
	if err != nil {
		return
	}
	ID := ctx.Param("id")
	id, err := strconv.Atoi(ID)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	err = c.app.Admin.DeleteAccount(id)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, "ok")
}
