package routers

import (
	"base-gin-go/config"
	"base-gin-go/middlewares"
	errorPkg "base-gin-go/pkg/errors"
	v1Routers "base-gin-go/routers/v1"
	"base-gin-go/usecase/auth"
	"base-gin-go/usecase/product"
	"base-gin-go/validations"
	"net/http"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func InitRouter(
	config *config.Environment,
	middleware middlewares.Middleware,
	productUseCase product.UseCase,
	authUseCase auth.UseCase,
	errorService errorPkg.Service,
) *gin.Engine {
	router := gin.New()
	router.Use(cors.New(cors.Config{
		AllowOrigins: strings.Split(config.CorsAllowOrigins, ","),
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders: []string{
			"Origin",
			"Content-Length",
			"Content-Type",
			"Access-Control-Allow-Headers",
			"Authorization",
			"X-XSRF-TOKEN",
		},
		ExposeHeaders: []string{
			"Content-Disposition",
		},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour, //nolint:gomnd // common
	}))
	router.Use(middleware.RestLogger)
	router.Use(gin.Recovery())
	// Validations
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("customEmail", validations.CustomEmail) //nolint:errcheck // no-error
	}
	apiRouter := router.Group("/api")
	apiRouter.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	v1Routers.InitV1Router(
		apiRouter.Group("/v1", middleware.Authentication),
		productUseCase,
		authUseCase,
		errorService,
	)
	return router
}
