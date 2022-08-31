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

func NewUserHandler(userService user.Service )*userHandler{
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	
	var input user.RegisterUserInput
	err :=c.ShouldBindJSON(&input)
	if err != nil {
		errors:=helper.FormatValidationError(err)
		
		errorMessage := gin.H{"errros": errors}

		response := helper.APIResponse("Register account failed from body", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err :=h.userService.RegisterUser(input)
	if err != nil {
		errors:=helper.FormatValidationError(err)
		errorMessage := gin.H{"errros": errors}
		response := helper.APIResponse("Register account failed to dataBase", http.StatusBadRequest, "error", errorMessage)

		c.JSON(http.StatusBadRequest, response)
		
		return
	}

	//token, err := h.jwtService.GenerateToken()

	formatter := user.FormattedUser(newUser, "TOKENTOKENTOKENTOKEN")

	response := helper.APIResponse("Account has been registered successfully", http.StatusOK, "success", formatter)


	c.JSON(http.StatusOK,response)
}

