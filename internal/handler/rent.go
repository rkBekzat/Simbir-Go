package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary GetTransport
// @Tags Rent
// @Description Get transport which not far than radius from point
// @Router /api/Rent/Transport [get]
// TODO: solve parameters confusion
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

// @Summary GetRent
// @Security ApKeyAuth
// @Tags Rent
// @Description Get rent by id
// @Router /api/Rent/:rentId [get]
func (c *Controller) GetRentById(ctx *gin.Context) {
	id, err := c.getUserId(ctx)
	if err != nil {
		NewErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}
	rentID := ctx.Param("rentId")
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

// @Summary MyHistory
// @Security ApKeyAuth
// @Tags Rent
// @Description Get the user history of rents
// @Router /api/Rent/MyHistory [get]
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

// @Summary TranposrtHistory
// @Security ApKeyAuth
// @Tags Rent
// @Description Get history of transport
// @Router /api/Rent/TransportHistory/:trnaposrtId [get]
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

// @Summary StartRent
// @Security ApKeyAuth
// @Tags Rent
// @Description Start renting
// @Router /api/Rent/New/:transportId [post]
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
	typerenting := ctx.Query("rentType")
	renting, err := c.app.Renting.StartRenting(id, tranId, typerenting)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, renting)
}

// @Summary End
// @Security ApKeyAuth
// @Tags Rent
// @Description End renting
// @Router /api/Rent/End/:rentId [post]
func (c *Controller) EndRenting(ctx *gin.Context) {
	id, err := c.getUserId(ctx)
	if err != nil {
		NewErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}
	rId := ctx.Param("rentId")
	rentId, err := strconv.Atoi(rId)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	la, lo := ctx.Query("lat"), ctx.Query("long")
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
	err = c.app.Renting.EndRenting(id, rentId, lat, long)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, "ok")
}
