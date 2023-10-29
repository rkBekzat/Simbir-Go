package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"vtb_api/internal/entities"
)

func handleParameters(x, y string) (int, int, error) {
	a, err := strconv.Atoi(x)
	if err != nil {
		return 0, 0, err
	}
	b, err := strconv.Atoi(x)
	if err != nil {
		return 0, 0, err
	}
	return a, b, nil
}

func (c *Controller) GetTransports(ctx *gin.Context) {
	err := c.isAdmin(ctx)
	if err != nil {
		return
	}
	st, co, tr := ctx.Query("start"), ctx.Query("count"), ctx.Query("transportType")
	start, count, err := handleParameters(st, co)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	transports, err := c.app.Admin.GetListOfTransports(start, count, tr)
	if err != nil {
		NewErrorResponse(ctx, http.StatusForbidden, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, transports)
}

func (c *Controller) GetTransportById(ctx *gin.Context) {
	err := c.isAdmin(ctx)
	if err != nil {
		return
	}
	i := ctx.Param("id")
	id, err := strconv.Atoi(i)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	transport, err := c.app.Admin.GetTransportById(id)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, transport)
}

func (c *Controller) AddTransport(ctx *gin.Context) {
	err := c.isAdmin(ctx)
	if err != nil {
		return
	}
	var tr entities.Transport
	if err := ctx.BindJSON(&tr); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	id, err := c.app.Admin.CreateTransport(&tr)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, id)
}

func (c *Controller) UpdatesTransport(ctx *gin.Context) {
	err := c.isAdmin(ctx)
	if err != nil {
		return
	}
	i := ctx.Param("id")
	id, err := strconv.Atoi(i)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	var tr entities.Transport
	if err := ctx.BindJSON(&tr); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	tr.Id = id
	err = c.app.Admin.UpdateTransport(&tr)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, "ok")
}

func (c *Controller) DeletesTransport(ctx *gin.Context) {
	err := c.isAdmin(ctx)
	if err != nil {
		return
	}
	i := ctx.Param("id")
	id, err := strconv.Atoi(i)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	err = c.app.Admin.DeleteTransport(id)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, "ok")
}
