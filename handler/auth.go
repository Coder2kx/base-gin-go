package handler

import (
	errorPkg "base-gin-go/pkg/errors"
	errors "base-gin-go/pkg/errors/custom"
	"base-gin-go/usecase/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context, authUseCase auth.UseCase, errorService errorPkg.Service) {
	var input auth.LoginInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		errValidate := errors.NewValidateError(ctx, input, err)
		ctx.JSON(http.StatusBadRequest, errValidate)
		return
	}
	output, err := authUseCase.Login(ctx, &input)
	if err != nil {
		errConverted := errorService.ParseInternalServer(err)
		ctx.JSON(errConverted.GetHTTPCode(), errConverted)
		return
	}
	ctx.JSON(http.StatusOK, output)
}
