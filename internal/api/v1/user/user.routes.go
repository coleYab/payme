package user

import "github.com/gin-gonic/gin"

type UserRoutes struct{
	handler *UserHandler
}

func NewUserRoutes(handler *UserHandler) *UserRoutes {
	return &UserRoutes{handler: handler}
}

func (u *UserRoutes) RegisterRoutes(routes *gin.RouterGroup) {
	routes.GET("/", u.handler.FindAllUsers)
	routes.GET("/by-id/:id", u.handler.FindUserByID)
	routes.GET("/by-email/:email", u.handler.FindUserByEmail)
	routes.PUT("/:id", u.handler.UpdateUser)
	routes.PATCH("/:id/promote", u.handler.UpdateUserRole)
	routes.PATCH("/:id/change-password", u.handler.UpdateUserPassword)
	routes.DELETE("/:id", u.handler.DeleteUser)
}
