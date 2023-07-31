package middlewares

import (
	"base-gin-go/domain/repository"
	jwtPkg "base-gin-go/pkg/jwt"
	stringPkg "base-gin-go/pkg/string"

	"github.com/gin-gonic/gin"
)

type Middleware interface {
	RestLogger(ctx *gin.Context)
	Authentication(ctx *gin.Context)
}

type middleware struct {
	jwtService     jwtPkg.Service
	stringService  stringPkg.Service
	userRepository repository.UserRepository
}

func NewMiddleware(
	jwtService jwtPkg.Service,
	stringService stringPkg.Service,
	userRepository repository.UserRepository,
) Middleware {
	return &middleware{
		jwtService:     jwtService,
		stringService:  stringService,
		userRepository: userRepository,
	}
}
