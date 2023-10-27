package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"vtb_api/internal/entities"
)

func (c *Controller) InfoTransport(ctx *gin.Context) {
	s := ctx.Param("id")
	id, err := strconv.Atoi(s)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	tr, err := c.app.Transport.GetById(id)
	if err != nil {
		NewErrorResponse(ctx, http.StatusForbidden, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, tr)
}

func (c *Controller) CreateTransport(ctx *gin.Context) {
	id, err := c.getUserId(ctx)
	if err != nil {
		NewErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}
	var tr entities.Transport
	if err := ctx.BindJSON(&tr); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	id, err = c.app.Transport.AddTransport(id, &tr)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, id)
}

func (c *Controller) UpdateTransport(ctx *gin.Context) {
	id, err := c.getUserId(ctx)
	if err != nil {
		NewErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}
	var tr entities.Transport
	if err := ctx.BindJSON(&tr); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	err = c.app.Transport.Update(id, &tr)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, "ok")
}

func (c *Controller) DeleteTransport(ctx *gin.Context) {
	id, err := c.getUserId(ctx)
	if err != nil {
		NewErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}
	s := ctx.Param("id")
	transportId, err := strconv.Atoi(s)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	err = c.app.Transport.Delete(id, transportId)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
}
