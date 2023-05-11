package handler

import (
	"net/http"
	"tripatra-api/auth"
	"tripatra-api/helper"
	"tripatra-api/user"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
}

func (h *userHandler) Register(ctx *gin.Context) {
	var inputRegister user.InputRegister

	err := ctx.ShouldBindJSON(&inputRegister)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Register failed saat input json", http.StatusUnprocessableEntity, "error", errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.Register(inputRegister)

	if err != nil {
		response := helper.APIResponse("Register failed saat insert", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := h.authService.GenerateToken(newUser.ID)
	if err != nil {
		response := helper.APIResponse("Register account failed generate token", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(newUser)

	response := helper.APIResponse("Successfully register", http.StatusOK, "success", gin.H{"user": formatter, "token": token})

	ctx.JSON(http.StatusOK, response)
}

func (h *userHandler) Login(ctx *gin.Context) {
	var input user.LoginAdminRequest

	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Login failed saat input json", http.StatusUnprocessableEntity, "error", errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedinUser, err := h.userService.Login(input)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Login failed saat cek email atau password", http.StatusUnprocessableEntity, "error", errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	token, err := h.authService.GenerateToken(loggedinUser.ID)
	if err != nil {
		response := helper.APIResponse("Login failed saat generate token", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(loggedinUser)

	response := helper.APIResponse("Success Log In", http.StatusOK, "success", gin.H{"user": formatter, "token": token})

	ctx.JSON(http.StatusOK, response)
}

func (h *userHandler) Logout(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(user.User)

	token, err := h.authService.DeleteToken(currentUser.ID)
	if err != nil {
		response := helper.APIResponse("Logout failed saat revoke token", http.StatusBadRequest, "error", err)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Logout Success", http.StatusOK, "success", gin.H{"token": token})

	ctx.JSON(http.StatusOK, response)
}

func (h *userHandler) Profile(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(user.User)

	formatter := user.FormatUser(currentUser)

	response := helper.APIResponse("Successfully fetch user data", http.StatusOK, "success", formatter)

	ctx.JSON(http.StatusOK, response)
}
