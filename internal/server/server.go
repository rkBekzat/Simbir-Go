package server

import "github.com/gin-gonic/gin"

func NewServer() {
	server := gin.Default()

	routing(server)
}
