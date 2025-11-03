package v1

import (
	"net/http"
	userdto "payme/internal/dto/users"
	"payme/internal/services"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *services.AuthServices
	userService *services.UserServices
}

func NewAuthHandler(authService *services.AuthServices, userService *services.UserServices) *AuthHandler {
	return &AuthHandler{authService: authService, userService: userService}
}

func (a *AuthHandler) RegisterUser(ctx *gin.Context) {
	var registerDto userdto.RegisterUserDto
	if err := ctx.ShouldBindJSON(&registerDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to parase the dto",
			"error":   err.Error(),
		})
		return
	}

	user, err := a.authService.RegisterUser(registerDto.Email, registerDto.FirstName, registerDto.LastName, registerDto.Password, registerDto.Role)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to register a user",
			"error":   err.Error(),
		})
		return
	}

	if err := a.userService.SaveUser(user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to create a user",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"user":  userdto.FromUserModel(user),
	})
}

func (a *AuthHandler) ForgetPassword(ctx *gin.Context) {
	// TODO: create reset token here
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func (a *AuthHandler) ResetPassword(ctx *gin.Context)  {
	// TODO: use the created reset token here
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func (a *AuthHandler) LoginUser(ctx *gin.Context)      {
	var loginDto userdto.LoginUserDto
	if err := ctx.ShouldBindJSON(&loginDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to parse the dto",
			"error":   err.Error(),
		})
		return
	}

	user, err := a.authService.LoginUser(loginDto.Email, loginDto.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to login",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"user": userdto.FromUserModel(user),
	})
}

func (a *AuthHandler) LogoutUser(ctx *gin.Context)     {
	// TODO: use cookies or other things
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
