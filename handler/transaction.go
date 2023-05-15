package handler

import (
	"net/http"
	"tripatra-api/email"
	"tripatra-api/helper"
	"tripatra-api/material"
	"tripatra-api/transaction"
	"tripatra-api/user"

	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	userService        user.Service
	transactionService transaction.Service
	materialService    material.Service
}

func NewTransactionHandler(userService user.Service, transactionService transaction.Service, materialService material.Service) *transactionHandler {
	return &transactionHandler{userService, transactionService, materialService}
}

func (h *transactionHandler) TransactionSubmission(ctx *gin.Context) {
	var inputTransaction transaction.InputTransaction

	err := ctx.ShouldBindJSON(&inputTransaction)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Transaction failed saat input json", http.StatusUnprocessableEntity, "error", errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// cekWarehouseStatus := false
	// if inputTransaction.WarehouseCategory == transaction.AddWarehouse {
	// 	cekWarehouseStatus = true
	// } else if inputTransaction.WarehouseCategory == transaction.TakeWarehouse {
	// 	cekWarehouseStatus = true
	// }

	// if !cekWarehouseStatus {
	// 	response := helper.APIResponse("Transaction Error WarehouseCategory tidak dikenal", http.StatusUnprocessableEntity, "error", nil)
	// 	ctx.JSON(http.StatusUnprocessableEntity, response)
	// 	return
	// }

	// material, err := h.materialService.FindByID(inputTransaction.MaterialID)
	// if material.ID == 0 {
	// 	response := helper.APIResponse("TransactionApproval Error, Material No matching records found", http.StatusUnprocessableEntity, "error", nil)
	// 	ctx.JSON(http.StatusUnprocessableEntity, response)
	// 	return
	// }

	currentUser := ctx.MustGet("currentUser").(user.User)

	inputTransaction.SenderID = currentUser.ID
	inputTransaction.Status = transaction.StatusReceive
	newTransaction, err := h.transactionService.Add(inputTransaction)
	if err != nil {
		response := helper.APIResponse("Transaction failed saat insert", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	emailSender, err := h.userService.GetUserByID(newTransaction.SenderID)
	if err != nil {
		response := helper.APIResponse("Failed saat insert get emailSender", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	emailReceiver, err := h.userService.GetUserByID(newTransaction.ReceiverID)
	if err != nil {
		response := helper.APIResponse("Failed saat insert get emailReceiver", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	materialName, err := h.materialService.FindByID(newTransaction.MaterialID)
	if err != nil {
		response := helper.APIResponse("Failed saat insert get materialName", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	err = email.SendingEmail(emailReceiver.Email, materialName.Name, emailSender.Email)
	if err != nil {
		response := helper.APIResponse("Failed saat sending email", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Successfully Submit Transaksi", http.StatusOK, "success", newTransaction)

	ctx.JSON(http.StatusOK, response)
}

func (h *transactionHandler) TransactionApproval(ctx *gin.Context) {
	var inputApproval transaction.InputApproval

	err := ctx.ShouldBindJSON(&inputApproval)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("TransactionApproval failed saat input json", http.StatusUnprocessableEntity, "error", errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	cekStatusTransaction := false
	if inputApproval.Status == transaction.StatusIssue {
		cekStatusTransaction = true
	} else if inputApproval.Status == transaction.StatusUpdated {
		cekStatusTransaction = true
	} else if inputApproval.Status == transaction.StatusDeleted {
		cekStatusTransaction = true
	}

	if !cekStatusTransaction {
		response := helper.APIResponse("TransactionApproval Error TransactionStatus tidak dikenal", http.StatusUnprocessableEntity, "error", nil)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	transactionData, err := h.transactionService.FindByID(inputApproval.TransactionID)
	if transactionData.ID == 0 {
		response := helper.APIResponse("TransactionApproval Error, Transaction No matching records found", http.StatusUnprocessableEntity, "error", nil)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	materialData, err := h.materialService.FindByID(transactionData.MaterialID)
	if materialData.ID == 0 {
		response := helper.APIResponse("TransactionApproval Error, Material No matching records found", http.StatusUnprocessableEntity, "error", nil)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var resutQuantity int
	if inputApproval.Status == transaction.StatusUpdated && transactionData.WarehouseCategory == transaction.AddWarehouse {
		resutQuantity = materialData.Quantity + transactionData.Quantity
	}
	if inputApproval.Status == transaction.StatusUpdated && transactionData.WarehouseCategory == transaction.TakeWarehouse {
		resutQuantity = materialData.Quantity - transactionData.Quantity
	}
	if inputApproval.Status != transaction.StatusUpdated {
		resutQuantity = materialData.Quantity
	}

	materialResult, err := h.materialService.Update(materialData.ID, resutQuantity)
	if err != nil {
		response := helper.APIResponse("TransactionApproval Error Saat Update Material", http.StatusUnprocessableEntity, "error", err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	_, err = h.transactionService.Update(transactionData.ID, inputApproval.Status, inputApproval.Reason)
	if err != nil {
		response := helper.APIResponse("TransactionApproval Error Saat Update Transaction", http.StatusUnprocessableEntity, "error", err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIResponse("Successfully update item "+materialResult.Name+" to Warehouse", http.StatusOK, "success", transactionData)

	ctx.JSON(http.StatusOK, response)
}

func (h *transactionHandler) SubmissionList(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(user.User)

	transactions, err := h.transactionService.FindBySender(currentUser.ID)

	if err != nil {
		response := helper.APIResponse("Failed saat get transaksi by sender", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Successfully get transaksi", http.StatusOK, "success", transactions)

	ctx.JSON(http.StatusOK, response)
}

func (h *transactionHandler) RecipientList(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(user.User)

	transactions, err := h.transactionService.FindByReceiver(currentUser.ID)

	if err != nil {
		response := helper.APIResponse("Failed saat get transaksi by receiver", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Successfully get transaksi", http.StatusOK, "success", transactions)

	ctx.JSON(http.StatusOK, response)
}

func (h *transactionHandler) Materials(ctx *gin.Context) {
	materials, err := h.materialService.GetAll()

	if err != nil {
		response := helper.APIResponse("Failed saat get Materials", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Successfully get Materials", http.StatusOK, "success", materials)

	ctx.JSON(http.StatusOK, response)
}

func (h *transactionHandler) AddMaterials(ctx *gin.Context) {
	var InputMaterial material.InputMaterial

	err := ctx.ShouldBindJSON(&InputMaterial)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Transaction failed saat input json", http.StatusUnprocessableEntity, "error", errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := ctx.MustGet("currentUser").(user.User)
	if currentUser.Role == "user" {
		if err != nil {
			response := helper.APIResponse("Failed saat Save Materials, user tidak boleh add materials", http.StatusBadRequest, "error", nil)
			ctx.JSON(http.StatusBadRequest, response)
			return
		}
	}

	material, err := h.materialService.Add(InputMaterial)

	if err != nil {
		response := helper.APIResponse("Failed saat Save Materials", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Successfully Save Materials", http.StatusOK, "success", material)

	ctx.JSON(http.StatusOK, response)
}

func (h *transactionHandler) WarehouseOfficer(ctx *gin.Context) {
	materials, err := h.userService.WarehouseOfficer()

	if err != nil {
		response := helper.APIResponse("Failed saat get WarehouseOfficer", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Successfully get WarehouseOfficer", http.StatusOK, "success", materials)

	ctx.JSON(http.StatusOK, response)
}
