package v1

import (
	"base-gin-go/handler"
	errorPkg "base-gin-go/pkg/errors"
	"base-gin-go/usecase/product"

	"github.com/gin-gonic/gin"
)

func initProductRouter(
	r gin.IRouter,
	productUseCase product.UseCase,
	errorService errorPkg.Service,
) {
	r.POST("", func(ctx *gin.Context) {
		handler.CreateProduct(ctx, productUseCase, errorService)
	})
	r.GET("", func(ctx *gin.Context) {
		handler.GetListProduct(ctx, productUseCase, errorService)
	})
	r.GET("/:id", func(ctx *gin.Context) {
		handler.GetProduct(ctx, productUseCase, errorService)
	})
	r.PUT("/:id", func(ctx *gin.Context) {
		handler.UpdateProduct(ctx, productUseCase, errorService)
	})
	r.DELETE("/:id", func(ctx *gin.Context) {
		handler.DeleteProduct(ctx, productUseCase, errorService)
	})
}
