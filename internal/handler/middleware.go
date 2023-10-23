package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorization = "Authorization"
	userCtx       = "userId"
)

func (c *Controller) UserIdentity(ctx *gin.Context) {
	token, err := c.getToken(ctx)
	if _, ok := c.blackListToken[token]; ok {
		NewErrorResponse(ctx, http.StatusUnauthorized, "this token expired	")
		return
	}

	userId, err := c.app.Auth.ParseToken(token)
	if err != nil {
		NewErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}
	ctx.Set(userCtx, userId)
}

func (c *Controller) getToken(ctx *gin.Context) (string, error) {
	header := ctx.GetHeader(authorization)
	if header == "" {
		return "", errors.New("empty auth header")
	}
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		return "", errors.New("invalid auth header")
	}
	return headerParts[1], nil
}
