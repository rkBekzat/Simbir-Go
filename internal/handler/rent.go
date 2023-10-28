package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (c *Controller) Transports(ctx *gin.Context) {
	la, lo, ra, tp := ctx.Query("lat"), ctx.Query("long"), ctx.Query("radius"), ctx.Query("type")
	lat, err := strconv.ParseFloat(la, 64)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	long, err := strconv.ParseFloat(lo, 64)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	rad, err := strconv.ParseFloat(ra, 64)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ids, err := c.app.Renting.AccessTransport(lat, long, rad, tp)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, ids)
}

func (c *Controller) GetRentById(ctx *gin.Context) {
	id, err := c.getUserId(ctx)
	if err != nil {
		NewErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}
	rentID := ctx.Param("id")
	rentId, err := strconv.Atoi(rentID)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	rent, err := c.app.Renting.GetById(id, rentId)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, rent)
}

func (c *Controller) MyHistory(ctx *gin.Context) {
	id, err := c.getUserId(ctx)
	if err != nil {
		NewErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}
	history, err := c.app.Renting.History(id)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, history)
}

func (c *Controller) TransportHistory(ctx *gin.Context) {
	id, err := c.getUserId(ctx)
	if err != nil {
		NewErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}
	tId := ctx.Param("transportId")
	tranId, err := strconv.Atoi(tId)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	history, err := c.app.Renting.TransportHistory(id, tranId)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, history)
}

func (c *Controller) RentTransport(ctx *gin.Context) {
	id, err := c.getUserId(ctx)
	if err != nil {
		NewErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}
	tId := ctx.Param("transportId")
	tranId, err := strconv.Atoi(tId)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	renting, err := c.app.Renting.StartRenting(id, tranId)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, renting)
}

func (c *Controller) EndRenting(ctx *gin.Context) {
	id, err := c.getUserId(ctx)
	if err != nil {
		NewErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}
	rId := ctx.Param("transportId")
	rentId, err := strconv.Atoi(rId)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	c.app.Renting.EndRenting(id, rentId)
}
