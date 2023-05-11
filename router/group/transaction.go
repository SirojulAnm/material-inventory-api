package group

import (
	"tripatra-api/auth"
	"tripatra-api/handler"
	"tripatra-api/material"
	"tripatra-api/transaction"
	"tripatra-api/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func TransactionRouter(db *gorm.DB, main *gin.RouterGroup) {
	userRepository := user.NewRepository(db)
	transaksiRepository := transaction.NewRepository(db)
	materialRepository := material.NewRepository(db)

	authService := auth.NewService()
	userService := user.NewService(userRepository)
	transaksiService := transaction.NewService(transaksiRepository)
	materialService := material.NewService(materialRepository)

	transaksiHandler := handler.NewTransactionHandler(userService, transaksiService, materialService)

	main.POST("/transaction-submit", authMiddleware(userService, authService), transaksiHandler.TransactionSubmission)
	main.POST("/transaction-approval", authMiddleware(userService, authService), transaksiHandler.TransactionApproval)
	main.GET("/submission-list", authMiddleware(userService, authService), transaksiHandler.SubmissionList)
	main.GET("/recipient-list", authMiddleware(userService, authService), transaksiHandler.RecipientList)
	main.GET("/materials", authMiddleware(userService, authService), transaksiHandler.Materials)
	main.POST("/materials", authMiddleware(userService, authService), transaksiHandler.AddMaterials)
	main.GET("/warehouse-officer", authMiddleware(userService, authService), transaksiHandler.WarehouseOfficer)
}
