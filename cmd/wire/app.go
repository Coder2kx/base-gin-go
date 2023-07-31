package wire

import (
	"base-gin-go/domain/repository"
	dataPkg "base-gin-go/pkg/data"
	errorPkg "base-gin-go/pkg/errors"
	jwtPkg "base-gin-go/pkg/jwt"
	passwordPkg "base-gin-go/pkg/password"
	stringPkg "base-gin-go/pkg/string"
	"base-gin-go/usecase/auth"
	"base-gin-go/usecase/product"
)

type App struct {
	// Service
	DataService     dataPkg.Service
	StringService   stringPkg.Service
	JwtService      jwtPkg.Service
	PasswordService passwordPkg.Service
	ErrorService    errorPkg.Service
	// Repository
	ProductRepository repository.ProductRepository
	UserRepository    repository.UserRepository
	// UseCase
	ProductUseCase product.UseCase
	AuthUseCase    auth.UseCase
}

func newApp(
	// Service
	dataService dataPkg.Service,
	stringService stringPkg.Service,
	jwtService jwtPkg.Service,
	passwordService passwordPkg.Service,
	errorService errorPkg.Service,
	// Repository
	productRepository repository.ProductRepository,
	userRepository repository.UserRepository,
	// UseCase
	productUseCase product.UseCase,
	authUseCase auth.UseCase,
) App {
	return App{
		// Service
		DataService:     dataService,
		StringService:   stringService,
		JwtService:      jwtService,
		PasswordService: passwordService,
		ErrorService:    errorService,
		// Repository
		ProductRepository: productRepository,
		UserRepository:    userRepository,
		// UseCase
		ProductUseCase: productUseCase,
		AuthUseCase:    authUseCase,
	}
}
