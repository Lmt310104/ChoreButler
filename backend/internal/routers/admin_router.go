package routers

import "github.com/gin-gonic/gin"

func SetupAdminRouter(group *gin.RouterGroup, server *APIServer) {

	group.POST("/login")

	group.POST("/register")

	group.GET("/profile")

	group.GET("/dashboard")

	group.GET("/user-dashboard")
}
