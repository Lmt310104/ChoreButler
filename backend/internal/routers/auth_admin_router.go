package routers

import "github.com/gin-gonic/gin"

func SetupAuthAdminRouter(group *gin.RouterGroup, server *APIServer) {
	group.POST("/login", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "Login",
		})
	})
}
