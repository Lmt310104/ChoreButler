package routers

import (
	docs "chorebutler/docs"
	"chorebutler/pkg/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"time"
)

type APIServer struct {
	Router   *gin.Engine
	Validate *validator.Validate
	Cfg      utils.Config
}

func NewApiServer(Validator *validator.Validate, config utils.Config) (*APIServer, error) {
	r := gin.Default()
	if config.GinMode != "release" {
		r.Use(cors.New(cors.Config{
			AllowAllOrigins:  true,
			AllowCredentials: true,
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
			MaxAge:           12 * time.Hour,
		}))
	}
	setupBasicRoutes(r)
	return &APIServer{
		r,
		Validator,
		config,
	}, nil
}

func setupBasicRoutes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello world!",
		})
	})
	r.GET("/health_check", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "I'm alive!",
		})
	})
}

func (server *APIServer) Start(address string) error {
	return server.Router.Run(address)
}

func (server *APIServer) SetupRouter() {
}

func (server *APIServer) SetupSwagger(swaggerURL string) {
	docs.SwaggerInfo.BasePath = "/"
	server.Router.GET(swaggerURL+"/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
