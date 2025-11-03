package v1

import "github.com/gin-gonic/gin"

type AuthRoutes struct{
	handler *AuthHandler
}

func NewAuthRoutes(handler *AuthHandler) *AuthRoutes {
	return &AuthRoutes{handler: handler}
}

func (u *AuthRoutes) RegisterRoutes(routes *gin.RouterGroup) {
	routes.POST("/register", u.handler.RegisterUser)
	routes.POST("/login", u.handler.LoginUser)
	routes.GET("/logout", u.handler.LogoutUser)
}
