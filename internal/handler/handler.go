package handler

import (
	"vtb_api/internal/service"
)

type Controller struct {
	app *service.UseCase
}

func NewController(app *service.UseCase) *Controller {
	return &Controller{app: app}
}
