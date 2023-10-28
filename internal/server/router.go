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
				transport.GET("")
				transport.GET("/:id")
				transport.POST("")
				transport.PUT("/:id")
				transport.DELETE("/:id")
			}
			rent := admin.Group("/Rend")
			{
				rent.GET("/:rentId")
				rent.POST("")
				rent.POST("/End/:rentId")
				rent.PUT("/:id")
				rent.DELETE("/:rentId")
			}
			admin.GET("/UserHistory/:userId")
			admin.GET("/TransportHistory/:transportId")
		}
	}
}
