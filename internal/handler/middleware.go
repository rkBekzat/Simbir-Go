package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorization = "Authorization"
	userCtx       = "userId"
)

func (c *Controller) UserIdentity(ctx *gin.Context) {
	header := ctx.GetHeader(authorization)
	if header == "" {
		NewErrorResponse(ctx, http.StatusUnauthorized, "empty auth header")
		return
	}
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		NewErrorResponse(ctx, http.StatusUnauthorized, "invalid auth header")
	}

	userId, err := c.app.Auth.ParseToken(headerParts[1])
	if err != nil {
		NewErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}
	ctx.Set(userCtx, userId)
}
