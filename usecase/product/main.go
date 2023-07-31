package product

import (
	"base-gin-go/domain/entity"
	"base-gin-go/domain/repository"
	"base-gin-go/infra/postgresql"
	dataPkg "base-gin-go/pkg/data"

	"github.com/gin-gonic/gin"
)

type UseCase interface {
	Create(ctx *gin.Context, input *CreateProductInput) (*entity.Product, error)
	Delete(ctx *gin.Context, input *DeleteProductInput) (*DeleteProductOutPut, error)
	GetByID(ctx *gin.Context, input *GetProductByIDInput) (*entity.Product, error)
	GetList(ctx *gin.Context, input *GetListProductInput) ([]*entity.Product, error)
	Update(ctx *gin.Context, input *UpdateProductInput) (*entity.Product, error)
	CreateWithTransaction(ctx *gin.Context, input *CreateProductInput) (*entity.Product, error)
}

type productUseCase struct {
	productRepository repository.ProductRepository
	dataService       dataPkg.Service
	database          *postgresql.Database
}

func NewProductUseCase(
	productRepository repository.ProductRepository,
	dataService dataPkg.Service,
	database *postgresql.Database,
) UseCase {
	return &productUseCase{
		productRepository: productRepository,
		dataService:       dataService,
		database:          database,
	}
}
