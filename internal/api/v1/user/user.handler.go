package user

import (
	"net/http"
	userdto "payme/internal/dto/users"
	"payme/internal/services"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *services.UserServices
}

func NewUserHandler (userService *services.UserServices) *UserHandler {
	return &UserHandler{userService: userService}
}

func (u *UserHandler)FindAllUsers(ctx *gin.Context) {
	page, err := strconv.Atoi(ctx.Query("page"))
	if err != nil {
		page = 0
	}

	limit, err := strconv.Atoi(ctx.Query("page"))
	if err != nil {
		limit = 50
	}

	users, err := u.userService.FindAllUsers(page * limit + 1, limit)
	usersDto := []userdto.UserDto{}
	for _, user := range users {
		usersDto = append(usersDto, userdto.FromUserModel(user))
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"users": usersDto,
		"metadata": map[string]any{
			"first": 0,
			"next": page + 1,
			"current": page,
			"previous": max(page - 1, 0),
			"count": len(usersDto),
			"last": 100,
			"limit": limit,
		},
	});
}

func (u *UserHandler)FindUserByEmail(ctx *gin.Context) {
	email := strings.Trim(ctx.Param("email"), " ")
	if len(email) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "email is required",
		})
		return
	}

	user, err := u.userService.FindUserByEmail(email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "user not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"user": userdto.FromUserModel(user),
	})
}

func (u *UserHandler)FindUserByID(ctx *gin.Context) {
	id := strings.Trim(ctx.Param("id"), " ")
	if len(id) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "id is required",
		})
		return
	}

	user, err := u.userService.FindUserByEmail(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "user not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"user": userdto.FromUserModel(user),
	})
}


func (u *UserHandler)UpdateUserRole(ctx *gin.Context) {
	var updateUserRoleDto userdto.UpdateUserRoleDto
	if err := ctx.ShouldBindJSON(&updateUserRoleDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to parase the dto",
			"error":   err.Error(),
		})
		return
	}

	// TODO: take the updater and check if he is admin
	id := ctx.Param("id")
	newUser, err := u.userService.UpdateRole(id, updateUserRoleDto.Role);
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to update user role",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"user": userdto.FromUserModel(newUser),
	})
}

func (u *UserHandler)UpdateUser(ctx *gin.Context) {
	var updateUserDto userdto.UpdateUserDto
	if err := ctx.ShouldBindJSON(&updateUserDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to parase the dto",
			"error":   err.Error(),
		})
		return
	}

	// TODO: take the updater and check if he is admin
	id := ctx.Param("id")
	newUser, err := u.userService.UpdateUser(id, updateUserDto.FirstName, updateUserDto.LastName);
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to update user profile",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"user": userdto.FromUserModel(newUser),
	})
}

func (u *UserHandler) UpdateUserPassword(ctx *gin.Context) {
	id := strings.Trim(ctx.Param("id"), " ")
	if len(id) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "id is required",
		})
		return
	}

	var updateUserPasswordDto userdto.UpdateUserPasswordDto
	if err := ctx.ShouldBindJSON(&updateUserPasswordDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to parase the dto",
			"error":   err.Error(),
		})
		return
	}
	user, err := u.userService.UpdateUserPassword(id, updateUserPasswordDto.OldPassword, updateUserPasswordDto.NewPassword)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to update user password",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"user": userdto.FromUserModel(user),
	})
}

func (u *UserHandler)DeleteUser(ctx *gin.Context) {
	id := strings.Trim(ctx.Param("id"), " ")
	if len(id) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "id is required",
		})
		return
	}

	if err := u.userService.DeleteUser(id); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "user not found",
		})
		return
	}


	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
