package v1

import (
	"base-gin-go/handler"
	errorPkg "base-gin-go/pkg/errors"
	"base-gin-go/usecase/auth"

	"github.com/gin-gonic/gin"
)

func initAuthRouter(
	r gin.IRouter,
	authUseCase auth.UseCase,
	errorService errorPkg.Service,
) {
	r.POST("/login", func(context *gin.Context) {
		handler.Login(context, authUseCase, errorService)
	})
}
