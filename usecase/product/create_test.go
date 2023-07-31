package product

import (
	"errors"
	"testing"

	"base-gin-go/infra/postgresql"
	mockRepository "base-gin-go/mock/domain/repository"
	mockDataPkg "base-gin-go/mock/pkg/data"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := &gin.Context{}
	mockDB, _ := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	mockProductRepository := mockRepository.NewMockProductRepository(ctrl)
	mockDataService := mockDataPkg.NewMockDataService(ctrl)
	productUseCase := NewProductUseCase(mockProductRepository, mockDataService, &postgresql.Database{DB: mockDB})
	mockDataService.EXPECT().Copy(gomock.Any(), gomock.Any()).Return(errors.New("Copy fail"))
	t.Run("Test copy fail", func(t *testing.T) {
		_, err := productUseCase.Create(ctx, &CreateProductInput{})
		if err != nil && err.Error() != "Copy fail" {
			t.Errorf("Test copy fail")
		}
	})
	mockDataService.EXPECT().Copy(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
	mockProductRepository.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil, errors.New("Fail"))
	t.Run("Test create fail", func(t *testing.T) {
		_, err := productUseCase.Create(ctx, &CreateProductInput{})
		if err != nil && err.Error() != "Fail" {
			t.Errorf("Test create fail")
		}
	})
	mockProductRepository.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil, nil)
	_, err := productUseCase.Create(ctx, &CreateProductInput{})
	t.Run("Test create success", func(t *testing.T) {
		if err != nil {
			t.Errorf("Test create success fail")
		}
	})
}
