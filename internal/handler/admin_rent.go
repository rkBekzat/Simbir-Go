package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"vtb_api/internal/entities"
)

func (c *Controller) GetRentId(ctx *gin.Context) {
	err := c.isAdmin(ctx)
	if err != nil {
		return
	}
	i := ctx.Param("rentId")
	id, err := strconv.Atoi(i)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	r, err := c.app.Admin.GetRentById(id)
	if err != nil {
		NewErrorResponse(ctx, http.StatusForbidden, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, r)
}

func (c *Controller) UserHistory(ctx *gin.Context) {
	err := c.isAdmin(ctx)
	if err != nil {
		return
	}
	u := ctx.Param("userId")
	id, err := strconv.Atoi(u)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	history, err := c.app.Admin.GetUserHistory(id)
	if err != nil {
		NewErrorResponse(ctx, http.StatusForbidden, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, history)
}

func (c *Controller) AdminTransportHistory(ctx *gin.Context) {
	err := c.isAdmin(ctx)
	if err != nil {
		return
	}
	t := ctx.Param("transportId")
	id, err := strconv.Atoi(t)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	history, err := c.app.Admin.GetTransportHistory(id)
	if err != nil {
		NewErrorResponse(ctx, http.StatusForbidden, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, history)
}

func (c *Controller) AdminNewRent(ctx *gin.Context) {
	err := c.isAdmin(ctx)
	if err != nil {
		return
	}
	var r entities.Rent
	if err := ctx.BindJSON(&r); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	id, err := c.app.Admin.NewRent(&r)
	if err != nil {
		NewErrorResponse(ctx, http.StatusForbidden, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, id)
}

func (c *Controller) AdminEndRent(ctx *gin.Context) {
	err := c.isAdmin(ctx)
	if err != nil {
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
	err = c.app.Admin.EndRent(rentId, lat, long)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, "ok")
}

func (c *Controller) AdminUpdateRent(ctx *gin.Context) {
	err := c.isAdmin(ctx)
	if err != nil {
		return
	}
	var r entities.Rent
	if err := ctx.BindJSON(&r); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	err = c.app.Admin.UpdateRent(&r)
	if err != nil {
		NewErrorResponse(ctx, http.StatusForbidden, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, "ok")
}

func (c *Controller) AdminDeleteRent(ctx *gin.Context) {
	err := c.isAdmin(ctx)
	if err != nil {
		return
	}
	rId := ctx.Param("rentId")
	rentId, err := strconv.Atoi(rId)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	err = c.app.Admin.DeleteRent(rentId)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, "ok")
}
