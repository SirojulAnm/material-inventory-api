package router

import (
	"tripatra-api/router/group"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Router(DB *gorm.DB) error {
	router := gin.Default()
	router.SetTrustedProxies(nil)
	corsConfig(router)

	v1 := router.Group("v1")
	group.AuthRouter(DB, v1)
	group.TransactionRouter(DB, v1)

	err := router.Run("0.0.0.0:8000")
	if err != nil {
		return err
	}

	return nil
}
