package handler

import (
	"vtb_api/internal/service"
)

type Controller struct {
	app            *service.UseCase
	blackListToken map[string]struct{}
}

func NewController(app *service.UseCase) *Controller {
	return &Controller{
		app:            app,
		blackListToken: make(map[string]struct{}),
	}
}
