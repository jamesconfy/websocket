package cmd

import (
	"io"
	"log"
	"os"

	"project-name/cmd/middleware"
	"project-name/cmd/routes"

	mysql "project-name/internal/database"
	"project-name/internal/logger"
	repo "project-name/internal/repository"
	"project-name/internal/service"

	"project-name/utils"

	_ "project-name/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Setup() {
	config, err := utils.LoadConfig("./")
	if err != nil {
		log.Println("Error loading configurations: ", err)
	}

	addr := config.ADDR
	if addr == "" {
		addr = "8000"
	}

	dsn := config.DATA_SOURCE_NAME
	if dsn == "" {
		log.Println("DSN cannot be empty")
	}

	secret := config.SECRET_KEY_TOKEN
	if secret == "" {
		log.Println("Please provide a secret key token")
	}

	host := config.HOST
	if host == "" {
		log.Println("Please provide an email host name")
	}

	port := config.PORT
	if port == "" {
		log.Println("Please provide an email port")
	}

	passwd := config.PASSWD
	if passwd == "" {
		log.Println("Please provide an email password")
	}

	email := config.EMAIL
	if email == "" {
		log.Println("Please provide an email address")
	}

	connection, err := mysql.NewMySQLServer(dsn)
	if err != nil {
		log.Println("Error Connecting to DB: ", err)
		return
	}
	defer connection.Close()
	conn := connection.GetConn()

	gin.DefaultWriter = io.MultiWriter(os.Stdout, logger.NewLogger())
	gin.DisableConsoleColor()

	router := gin.New()
	v1 := router.Group("/api/v1")
	v1.Use(gin.Logger())
	v1.Use(gin.Recovery())
	router.Use(middleware.CORS())

	// Auth Repository
	authRepo := repo.NewAuthRepo(conn)

	// User Repository
	userRepo := repo.NewUserRepo(conn)

	// Email Service
	emailSrv := service.NewEmailSrv(email, passwd, host, port)

	// Token Service
	authSrv := service.NewAuthService(authRepo, secret)

	// Validation Service
	validatorSrv := service.NewValidationStruct()

	// Cryptography Service
	cryptoSrv := service.NewCryptoSrv()

	// Home Service
	homeSrv := service.NewHomeSrv()

	// User Service
	userSrv := service.NewUserSrv(userRepo, authRepo, validatorSrv, cryptoSrv, authSrv, emailSrv)

	// Routes
	routes.HomeRoute(v1, homeSrv)
	routes.UserRoute(v1, userSrv)
	routes.ErrorRoute(router)

	// Documentation
	v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":" + addr)
}
