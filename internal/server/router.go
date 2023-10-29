package server

import (
	"github.com/gin-gonic/gin"
	"vtb_api/internal/handler"
)

func routing(router *gin.Engine, controller *handler.Controller) {
	api := router.Group("/api")
	{
		acc := api.Group("/Account")
		{
			acc.GET("/Me", controller.InfoUser)
			acc.POST("/SignIn", controller.SignIn)
			acc.POST("/SignUp", controller.SignUp)
			acc.POST("/SignOut", controller.SignOut)
			acc.PUT("/Update", controller.UpdateUser)
		}
		transport := api.Group("/Transport")
		{
			transport.GET("/:id", controller.InfoTransport)
			transport.POST("", controller.CreateTransport)
			transport.PUT("/:id", controller.UpdateTransport)
			transport.DELETE("/:id", controller.DeleteTransport)
		}
		rent := api.Group("/Rent")
		{
			rent.GET("/Transport", controller.Transports)
			rent.GET("/:rentId", controller.GetRentById)
			rent.GET("/MyHistory", controller.MyHistory)
			rent.GET("/TransportHistory/:transportId", controller.TransportHistory)
			rent.POST("/New/:transportId", controller.RentTransport)
			rent.POST("/End/:rentId", controller.EndRenting)
		}
		api.POST("/Payment/Hesoyam/:accountId")
		admin := api.Group("/Admin")
		{
			acc := admin.Group("/Account")
			{
				acc.GET("", controller.ListAccounts)
				acc.GET("/:id", controller.GetAccountById)
				acc.POST("", controller.CreateAccountByAdmin)
				acc.PUT("/:id", controller.UpdateAccount)
				acc.DELETE("/:id", controller.DeleteAccount)
			}
			transport := admin.Group("/Transport")
			{
				transport.GET("", controller.GetTransports)
				transport.GET("/:id", controller.GetTransportById)
				transport.POST("", controller.AddTransport)
				transport.PUT("/:id", controller.UpdatesTransport)
				transport.DELETE("/:id", controller.DeletesTransport)
			}
			rent := admin.Group("/Rent")
			{
				rent.GET("/:rentId", controller.GetRentId)
				rent.POST("", controller.AdminNewRent)
				rent.POST("/End/:rentId", controller.AdminEndRent)
				rent.PUT("/:id", controller.AdminUpdateRent)
				rent.DELETE("/:rentId", controller.AdminDeleteRent)
			}
			admin.GET("/UserHistory/:userId", controller.UserHistory)
			admin.GET("/TransportHistory/:transportId", controller.AdminTransportHistory)
		}
	}
}
