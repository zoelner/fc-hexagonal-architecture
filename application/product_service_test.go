package application_test

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/zoelner/hexagonal-architecture/application"
	mock_application "github.com/zoelner/hexagonal-architecture/application/mocks"
	"testing"
)

func TestProductService_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_application.NewMockProductInterface(ctrl)
	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)
	persistence.EXPECT().Get(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.ProductService{
		Persistence: persistence,
	}

	result, err := service.Get("abc")
	require.Nil(t, err)
	require.Equal(t, product, result)
}

func TestProductService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_application.NewMockProductInterface(ctrl)
	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)
	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.ProductService{
		Persistence: persistence,
	}

	result, err := service.Create("Product 1", 10)
	require.Nil(t, err)
	require.Equal(t, product, result)
}

func TestProductService_Enable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_application.NewMockProductInterface(ctrl)
	product.EXPECT().Enable().Return(nil)

	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)
	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.ProductService{
		Persistence: persistence,
	}

	result, err := service.Enable(product)
	require.Nil(t, err)
	require.Equal(t, product, result)
}

func TestProductService_Disable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_application.NewMockProductInterface(ctrl)
	product.EXPECT().Disable().Return(nil)

	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)
	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.ProductService{
		Persistence: persistence,
	}

	result, err := service.Disable(product)
	require.Nil(t, err)
	require.Equal(t, product, result)
}
