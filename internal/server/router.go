package server

import "github.com/gin-gonic/gin"

func routing(router *gin.Engine) {
	api := router.Group("/api")
	{
		acc := api.Group("/Account")
		{
			acc.GET("/Me")
			acc.POST("/SignIn")
			acc.POST("/SignUp")
			acc.POST("/SignOut")
			api.PUT("/Update")
		}
		transport := api.Group("/Transport")
		{
			transport.GET("/:id")
			transport.POST("")
			transport.PUT("/:id")
			transport.DELETE("/:id")
		}
		rent := api.Group("/Rent")
		{
			rent.GET("/Transport")
			rent.GET("/:rentId")
			rent.GET("/MyHistory")
			rent.GET("/TransportHistory/:transportId")
			rent.POST("/New/:transportId")
			rent.POST("/End/:rentId")
		}
		api.POST("/Payment/Hesoyam/:accountId")
		admin := api.Group("/Admin")
		{
			acc := admin.Group("/Account")
			{
				acc.GET("")
				acc.GET("/:id")
				acc.POST("")
				acc.PUT("/:id")
				acc.DELETE("/:id")
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
