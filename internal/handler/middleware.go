package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strings"
)

const (
	authorization = "Authorization"
	userCtx       = "userId"
)

func (c *Controller) UserIdentity(ctx *gin.Context) error {
	token, err := c.getToken(ctx)
	if _, ok := c.blackListToken[token]; ok {
		return errors.New("this token expired")
	}

	userId, err := c.app.Auth.ParseToken(token)
	if err != nil {
		return err
	}
	ctx.Set(userCtx, userId)
	return nil
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
