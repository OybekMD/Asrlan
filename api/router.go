package api

import (
	_ "asrlan-monolight/api/docs" // swag
	v1 "asrlan-monolight/api/handlers/v1"
	"asrlan-monolight/api/middleware"

	"asrlan-monolight/api/tokens"
	"asrlan-monolight/config"
	"asrlan-monolight/storage"

	"github.com/casbin/casbin/v2"
	"github.com/gin-contrib/cors"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

// Option ...
type Option struct {
	Conf     *config.Config
	Storage  storage.StorageI
	Enforcer *casbin.Enforcer
}

// @Title       	Asrlan-Monolithic
// @securityDefinitions.apikey BearerAuth
// @In          	header
// @Name        	Authorization
func New(option *Option) *gin.Engine {
	gin.SetMode(option.Conf.GinMode)
	router := gin.New()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowCredentials = true
	corsConfig.AllowHeaders = []string{"*"}
	corsConfig.AllowBrowserExtensions = true
	corsConfig.AllowMethods = []string{"*"}
	router.Use(cors.New(corsConfig))

	jwtHandler := tokens.JWTHandler{
		SigninKey: option.Conf.SigningKey,
	}

	handlerV1 := v1.New(&v1.HandlerV1Options{
		Cfg:        option.Conf,
		Storage:    option.Storage,
		JWTHandler: jwtHandler,
		Enforcer:   option.Enforcer,
	})

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.NewAuthorizer(option.Enforcer, jwtHandler, *option.Conf))

	api := router.Group("/v1")

	// Login-Passwords
	api.POST("/login", handlerV1.Login)
	api.POST("/signup", handlerV1.Signup)
	api.POST("/verify", handlerV1.Verify)
	api.POST("/forgot", handlerV1.ForgotPassword)
	api.POST("/resetpassword", handlerV1.ResetPassword)
	// api.GET("/token/:refresh", handlerV1.Token)
	// api.GET("/users/:login", handlerV1.GetInfo)

	// Fileup
	api.POST("/pdfupload", handlerV1.UploadPDFFile)
	api.POST("/badgeupload", handlerV1.BadgeImageFile)
	api.POST("/avatarupload", handlerV1.AvatarImageFile)
	api.POST("/soundupload", handlerV1.UploadSoundFile)
	api.POST("/videoupload", handlerV1.UploadVideoFile)
	router.Static("/media", "./media")

	// Dashboard
	api.GET("/navbar", handlerV1.GetNavbar)
	api.GET("/dashboard", handlerV1.GetDashboard)
	api.GET("/leaderboard", handlerV1.GetLeaderboard)

	// User
	api.GET("/user/:id", handlerV1.GetUser)
	api.DELETE("/user/:id", handlerV1.DeleteUser)
	api.PUT("/user", handlerV1.UpdateUser)
	api.PUT("/userpassword", handlerV1.UpdateUserPassword)

	// ---------Problems!!!!!!!!
	// Activity
	api.POST("/activity", handlerV1.CreateActivity)
	api.GET("/activitys", handlerV1.GetAllGroupedByMonth)
	api.GET("/activitysch", handlerV1.GetAllGroupedByChoice)

	// Social
	api.PUT("/social", handlerV1.UpdateSocial)
	api.GET("/social/:id", handlerV1.GetSocial)

	// Badge
	api.GET("/badge/:id", handlerV1.GetBadge)
	api.POST("/badge", handlerV1.CreateBadge)
	api.PUT("/badge", handlerV1.UpdateBadge)
	api.DELETE("/badge/:id", handlerV1.DeleteBadge)
	api.GET("/badges", handlerV1.ListBadges)

	// Language
	api.GET("/language/:id", handlerV1.GetLanguage)
	api.POST("/language", handlerV1.CreateLanguage)
	api.PUT("/language", handlerV1.UpdateLanguage)
	api.DELETE("/language/:id", handlerV1.DeleteLanguage)
	api.GET("/languages", handlerV1.ListLanguages)
	api.GET("/languagesforregister", handlerV1.LanguagesForRegister)

	// Level
	api.GET("/level/:id", handlerV1.GetLevel)
	api.POST("/level", handlerV1.CreateLevel)
	api.PUT("/level", handlerV1.UpdateLevel)
	api.DELETE("/level/:id", handlerV1.DeleteLevel)
	api.GET("/levels", handlerV1.ListLevels)
	api.GET("/levelsforregister", handlerV1.LevelsForRegister)
	api.GET("/levelsforcourse", handlerV1.LevelsForCourse)

	// Topic
	api.GET("/topic/:id", handlerV1.GetTopic)
	api.POST("/topic", handlerV1.CreateTopic)
	api.PUT("/topic", handlerV1.UpdateTopic)
	api.DELETE("/topic/:id", handlerV1.DeleteTopic)
	api.GET("/topics", handlerV1.ListTopics)

	// Lesson
	api.GET("/lesson/:id", handlerV1.GetLesson)
	api.POST("/lesson", handlerV1.CreateLesson)
	api.PUT("/lesson", handlerV1.UpdateLesson)
	api.DELETE("/lesson/:id", handlerV1.DeleteLesson)
	api.GET("/lessons", handlerV1.ListLessons)

	// Content
	api.GET("/content/:id", handlerV1.GetContent)
	api.POST("/content", handlerV1.CreateContent)
	api.PUT("/content", handlerV1.UpdateContent)
	api.DELETE("/content/:id", handlerV1.DeleteContent)
	api.GET("/contents", handlerV1.ListContents)
	
	// ContentFile
	api.GET("/contentfile/:id", handlerV1.GetContentFile)
	api.POST("/contentfile", handlerV1.CreateContentFile)
	api.DELETE("/contentfile/:id", handlerV1.DeleteContentFile)

	// Book
	api.GET("/book/:id", handlerV1.GetBook)
	api.POST("/book", handlerV1.CreateBook)
	api.PUT("/book", handlerV1.UpdateBook)
	api.DELETE("/book/:id", handlerV1.DeleteBook)
	api.GET("/books/:id", handlerV1.ListBooks)

	// Profile
	api.GET("/profile", handlerV1.Profile)

	url := ginSwagger.URL("swagger/doc.json")
	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router
}
