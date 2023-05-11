package group

import (
	"fmt"
	"tripatra-api/auth"
	"tripatra-api/handler"
	"tripatra-api/helper"
	"tripatra-api/user"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthRouter(db *gorm.DB, main *gin.RouterGroup) {
	userRepository := user.NewRepository(db)

	authService := auth.NewService()
	userService := user.NewService(userRepository)

	userHandler := handler.NewUserHandler(userService, authService)

	main.POST("/login", userHandler.Login)
	main.POST("/register", userHandler.Register)
	main.GET("/logout", authMiddleware(userService, authService), userHandler.Logout)
	main.GET("/profile", authMiddleware(userService, authService), userHandler.Profile)
}

func authMiddleware(userService user.Service, authable auth.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		fmt.Println(tokenString)
		if tokenString == "" {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		customClaim, err := authable.ValidateToken(tokenString)
		if customClaim == nil && err != nil {
			response := helper.APIResponse("SessionExpired", http.StatusUnauthorized, "error", err.Error())
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID, err := strconv.ParseInt(customClaim.Subject, 10, 64)

		user, err := userService.GetUserByID(int(userID))
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", user)
	}
}
