package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"vtb_api/internal/entities"
)

// @Summary		AdminGetRent
// @Security		ApiKeyAuth
// @Tags			Admin
// @Accept			json
// @Produce		json
// @Param			rentId	path	int	true	"rent id"
// @Description	Get rent by id
// @Router			/api/Admin/Rent/:rentId [get]
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

// @Summary		AdminGetUserHistory
// @Security		ApiKeyAuth
// @Tags			Admin
// @Accept			json
// @Produce		json
// @Description	Get user history
// @Router			/api/UserHistory/:userId [get]
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

// @Summary		AdminGetTransportHitory
// @Security		ApiKeyAuth
// @Tags			Admin
// @Accept			json
// @Produce		json
// @Param			transportId	path	int	true	"transport id"
// @Description	Get  transport history
// @Router			/api/TransportHistory/:transportId [get]
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

// @Summary		AdminNewRent
// @Security		ApiKeyAuth
// @Tags			Admin
// @Accept			json
// @Produce		json
// @Description	Create new rent
// @Router			/api/Admin/Rent [post]
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

// @Summary		AdminEndRent
// @Security		ApiKeyAuth
// @Tags			Admin
// @Accept			json
// @Produce		json
// @Param			rentId	path	int	true	"rent id"
// @Description	End the rent
// @Router			/api/Admin/Rent/End/:rentId [post]
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

// @Summary		AdminUpdateRent
// @Security		ApiKeyAuth
// @Tags			Admin
// @Accept			json
// @Produce		json
// @Param			id	path	int	true	"rent id"
// @Description	Update the rent
// @Router			/api/Admin/Rent/:id [put]
func (c *Controller) AdminUpdateRent(ctx *gin.Context) {
	err := c.isAdmin(ctx)
	if err != nil {
		return
	}
	rId := ctx.Param("id")
	rentId, err := strconv.Atoi(rId)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	var r entities.Rent
	if err := ctx.BindJSON(&r); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	r.Id = rentId
	err = c.app.Admin.UpdateRent(&r)
	if err != nil {
		NewErrorResponse(ctx, http.StatusForbidden, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, "ok")
}

// @Summary		AdminDeleteRent
// @Security		ApiKeyAuth
// @Tags			Admin
// @Accept			json
// @Produce		json
// @Param			rentId	path	int	true	"rent id"
// @Description	Delete the rent
// @Router			/api/Admin/Rent/:rentId [delete]
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
