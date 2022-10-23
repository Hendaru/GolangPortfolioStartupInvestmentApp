package main

import (
	"bwastartup/auth"
	"bwastartup/campaign"
	"bwastartup/handler"
	"bwastartup/helper"
	"bwastartup/payment"
	"bwastartup/transaction"
	"bwastartup/user"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:root@tcp(127.0.0.1:8889)/bwstartup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	//REPOSITORY YG BERHUBUNGAN DENGAN DB
	userRepository := user.NewRepository(db)
	campaignRepository := campaign.NewRepository(db)
	transactionRepository := transaction.NewRepository(db)

	//SERVICE YG BERHUBUNGAN DENGAN VALIDASI
	userService := user.NewService(userRepository)
	campaignService := campaign.NewService(campaignRepository)
	authService := auth.NewService()
	paymentService := payment.NewService()
	transactionService := transaction.NewService(transactionRepository, campaignRepository, paymentService)

	//HANDLER YG BERHUBUNGAN DENGAN RESPON EX : 200, 400, 404
	userHandler := handler.NewUserHandler(userService, authService)
	campaignHandler := handler.NewCampaignHandler(campaignService)
	transactionHandler := handler.NewTransactionHandler(transactionService)

	//ROUTING
	// userService.SaveAvatar(5, "images/1-profile.png")
	router := gin.Default()
	router.Use(cors.Default())
	router.Static("/images", "./images")
	//api version ex v1, v2 dll
	api := router.Group("/api/v1")

	//USER
	api.POST("/users", userHandler.RegisterUserHandler)
	api.POST("/sessions", userHandler.LoginHandler)
	api.POST("/email_checkers", userHandler.CheckEmailAvailabilityHandler)
	api.POST("/avatars", authMiddleware(authService, userService), userHandler.UploadAvatarHandler)
	api.POST("/user/fetch", authMiddleware(authService, userService), userHandler.FetchUser)

	//CAMPAIGN
	api.GET("/campaigns", campaignHandler.GetAllCampaignsHandler)
	api.GET("/campaigns/:id", campaignHandler.GetCampaignHandler)
	api.POST("/campaigns", authMiddleware(authService, userService), campaignHandler.CreateCampaignHandler)
	api.PUT("/campaigns/:id", authMiddleware(authService, userService), campaignHandler.UpdateCampaignHandler)
	api.POST("/campaign-image", authMiddleware(authService, userService), campaignHandler.UploadCampaignImageHandler)

	//TRANSACTION
	api.GET("/campaign/:id/transactions", authMiddleware(authService, userService), transactionHandler.GetCampaignTransactionHandler)
	api.GET("/transactions", authMiddleware(authService, userService), transactionHandler.GetUserTransactionSHandler)
	api.POST("/transactions", authMiddleware(authService, userService), transactionHandler.CreateTransactionHandler)
	api.POST("/transactions/notification", authMiddleware(authService, userService), transactionHandler.GetNotificationTransactionHandler)

	//go run main.go
	router.Run()

}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}
		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		userID := int(claim["user_id"].(float64))

		user, err := userService.GetUserByIdUserService(userID)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		c.Set("currentUser", user)

	}
}
