package handler

import (
	"bwastartup/helper"
	"bwastartup/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	// tangkap input user
	// map input user ke struct RegisterUserInput
	// struct userHandler dilempar sebagai parameter service

	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormaValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.ApiResponse("Register account failed", http.StatusUnprocessableEntity, "false", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	
	newUser, err := h.userService.RegisterUser(input)

	if err != nil {
		response := helper.ApiResponse("Register account failed", http.StatusBadRequest, "false", nil)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	// token, err := h.jwtService.GenerateToken()

	formatter := user.FormatUser(newUser, "tokentookentoken")

	response := helper.ApiResponse("Account has been registered", http.StatusOK, "success", formatter)


	c.JSON(http.StatusOK, response)
}

func (h *userHandler) Login(c *gin.Context) {
// input user
// input masuk ke handler, mapping input user ke input struct
// passing input strcut ke service
// dalam service mencari dengan bantuan repository user dengan nilai x
// mencocokan password

	var input user.LoginUserInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormaValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.ApiResponse("Login failed", http.StatusUnprocessableEntity, "false", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedinUser, err := h.userService.LoginUser(input)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.ApiResponse("Login failed", http.StatusUnprocessableEntity, "false", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := user.FormatUser(loggedinUser, "tokentookentoken")

	response := helper.ApiResponse("Successfully login", http.StatusOK, "success", formatter)


	c.JSON(http.StatusOK, response)
}