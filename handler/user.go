package handler

import (
	"bwastartup/helper"
	"bwastartup/user"
	"fmt"
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

func (h *userHandler)Login(c *gin.Context) {
	var input user.LoginInput
	err :=c.ShouldBindJSON(&input)
	if err != nil {
		errors:=helper.FormatValidationError(err)
		
		errorMessage := gin.H{"errros": errors}

		response := helper.APIResponse("Login Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedinUser, err := h.userService.Login(input)

	if err != nil {
		errorMessage := gin.H{"errros": err.Error()}
		response := helper.APIResponse("Login Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := user.FormattedUser(loggedinUser, "TOKENTOKENTOKENTOKEN")

	response := helper.APIResponse("Successfully loggedin", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK,response)

}

func (h *userHandler) CheckEmailAvailability(c *gin.Context) {
	var input user.CheckEmailInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors:=helper.FormatValidationError(err)
		errorMessage := gin.H{"errros": errors}
		response := helper.APIResponse("Email checking failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	IsEmailAvailable, err := h.userService.IsEmailAvailable(input)

	if err != nil {
		errorMessage := gin.H{"errros": "Server error"}
		response := helper.APIResponse("Login Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{
		"is_available": IsEmailAvailable,
	}

	var metaMessage string
	if IsEmailAvailable{
		metaMessage = "Email is available"
	}else{
		metaMessage = "Email has been registered"
	}

	response := helper.APIResponse(metaMessage, http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) UploadAvatar(c *gin.Context){
	
	file, err := c.FormFile("avatar")
	if err != nil {
		data := gin.H{"is_uploaded":false}
		response := helper.APIResponse("Failed to uplod avatar image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}
	userID := 6
	path := fmt.Sprintf("images/%d-%s", userID, file.Filename)

	err = c.SaveUploadedFile(file, path)
		if err != nil {
			data := gin.H{"is_uploaded":false}
			response := helper.APIResponse("Failed to uplod avatar image", http.StatusBadRequest, "error", data)
	
			c.JSON(http.StatusBadRequest, response)
			return
		}
	


	_, err =  h.userService.SaveAvatar(userID, path)
	if err != nil {
		data := gin.H{"is_uploaded":false}
		response := helper.APIResponse("Failed to uplod avatar image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded":true}
	response := helper.APIResponse("Avatar successfully uploaded", http.StatusOK, "success", data)

	c.JSON(http.StatusOK, response)
	
}

