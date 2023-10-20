package server

import (
	"github.com/gin-gonic/gin"
	"vtb_api/internal/handler"
)

func NewServer(controller *handler.Controller) *gin.Engine {
	server := gin.Default()

	routing(server, controller)

	return server
}
