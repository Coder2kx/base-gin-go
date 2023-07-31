//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package wire

import (
	"base-gin-go/config"
	"base-gin-go/infra/postgresql"
	"base-gin-go/infra/postgresql/repository"
	dataPkg "base-gin-go/pkg/data"
	errorPkg "base-gin-go/pkg/errors"
	jwtPkg "base-gin-go/pkg/jwt"
	passwordPkg "base-gin-go/pkg/password"
	stringPkg "base-gin-go/pkg/string"
	"base-gin-go/usecase/auth"
	"base-gin-go/usecase/product"

	"github.com/google/wire"
)

func InitApp(config *config.Environment, database *postgresql.Database) (App, error) {
	panic(wire.Build(
		// Service
		wire.NewSet(dataPkg.NewDataService),
		wire.NewSet(stringPkg.NewStringService),
		wire.NewSet(jwtPkg.NewJwtService),
		wire.NewSet(passwordPkg.NewPasswordService),
		wire.NewSet(errorPkg.NewErrorService),
		// Repository
		wire.NewSet(repository.NewProductRepository),
		wire.NewSet(repository.NewUserRepository),
		// UseCase
		wire.NewSet(product.NewProductUseCase),
		wire.NewSet(auth.NewAuthUseCase),
		newApp,
	))
}
