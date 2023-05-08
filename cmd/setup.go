package cmd

import (
	"fmt"
	"io"
	"log"
	"os"

	middleware "project-name/cmd/middleware"
	routes "project-name/cmd/routes"
	_ "project-name/docs"
	sql "project-name/internal/database"
	"project-name/internal/logger"
	repo "project-name/internal/repository"
	service "project-name/internal/service"
	utils "project-name/utils"

	gin "github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	addr         string
	mode         string
	dsn          string
	secret       string
	email        string
	email_passwd string
	email_host   string
	email_port   string
)

func Setup() {
	router := gin.New()
	v1 := router.Group("/api/v1")
	v1.Use(gin.Logger())
	v1.Use(gin.Recovery())
	router.Use(middleware.CORS())

	db, err := sql.New(dsn)
	if err != nil {
		log.Println("Error Connecting to DB: ", err)
	}
	defer db.Close()
	conn := db.GetConn()

	// Auth Repository
	authRepo := repo.NewAuthRepo(conn)

	// User Repository
	userRepo := repo.NewUserRepo(conn)

	// Email Service
	emailSrv := service.NewEmailSrv(email, email_passwd, email_host, email_port)

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

func init() {
	godotenv.Load(".env")

	addr = utils.AppConfig.ADDR
	if addr == "" {
		addr = "8000"
	}

	secret = utils.AppConfig.SECRET_KEY_TOKEN
	if secret == "" {
		log.Println("Please provide a secret key token")
	}

	mode = utils.AppConfig.MODE
	if mode == "development" {
		loadDev()
	}

	if mode == "production" {
		loadProd()
	}

}

func loadDev() {
	gin.SetMode(gin.DebugMode)

	host := utils.AppConfig.POSTGRES_HOST
	username := utils.AppConfig.POSTGRES_USERNAME
	passwd := utils.AppConfig.POSTGRES_PASSWORD
	dbname := utils.AppConfig.POSTGRES_DBNAME

	dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, username, passwd, dbname)
	if dsn == "" {
		log.Println("DSN cannot be empty")
	}

	email_host = utils.AppConfig.HOST
	if email_host == "" {
		log.Println("Please provide an email host name")
	}

	email_port = utils.AppConfig.PORT
	if email_port == "" {
		log.Println("Please provide an email port")
	}

	email_passwd = utils.AppConfig.PASSWD
	if email_passwd == "" {
		log.Println("Please provide an email password")
	}

	email = utils.AppConfig.EMAIL
	if email == "" {
		log.Println("Please provide an email address")
	}
}

func loadProd() {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.MultiWriter(os.Stdout, logger.NewLogger())
	gin.DisableConsoleColor()
}

var _ = loadProd
