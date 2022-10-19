package handler

import (
	"bwastartup/helper"
	"bwastartup/transaction"
	"bwastartup/user"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	service transaction.Service
}

func NewTransactionHandler(service transaction.Service) *transactionHandler {
	return &transactionHandler{service}
}

func (h *transactionHandler) GetCampaignTransactionHandler(c *gin.Context) {
	var input transaction.GetCampaignTransactionInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to  campaign transaction get id", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	transactionS, err := h.service.GetTransactionByCampaignIDService(input)
	if err != nil {
		response := helper.APIResponse("Failed to get campaign's transactions Service ", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to get campaign's trsnsactions", http.StatusOK, "success", transaction.FormatCampaignTransactionS(transactionS))
	c.JSON(http.StatusOK, response)

}

func (h *transactionHandler) GetUserTransactionSHandler(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)
	userID := currentUser.ID

	transactionS, err := h.service.GetTransactionByUserIDService(userID)
	if err != nil {
		response := helper.APIResponse("Failed to get user transactions ", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return

	}
	// fmt.Println(transactionS)
	for _, t := range transactionS {
		fmt.Println(t.Campaign.ID)
	}

	response := helper.APIResponse("Success to get User trsnsactions", http.StatusOK, "success", transaction.FormatUserTransactionS(transactionS))
	c.JSON(http.StatusOK, response)

}

func (h *transactionHandler) CreateTransactionHandler(c *gin.Context) {
	var input transaction.CreateTransactionInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errros": errors}

		response := helper.APIResponse("Failed to create transaction", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	newTransaction, err := h.service.CreatTransactionService(input)

	if err != nil {
		response := helper.APIResponse("Failed to create transaction", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
	}

	response := helper.APIResponse("Success to create transaction", http.StatusOK, "success", transaction.FormatTransaction(newTransaction))
	c.JSON(http.StatusOK, response)

}
