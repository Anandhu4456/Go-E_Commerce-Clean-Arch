package usecase

import ("github.com/Anandhu4456/go-Ecommerce/pkg/repository/interfaces"
	services"github.com/Anandhu4456/go-Ecommerce/pkg/usecase/interfaces"
)

type inventoryUsecase struct{
	invRepo interfaces.InventoryRespository
}

// constructor function
func NewInventoryUsecase(invRepo interfaces.InventoryRespository)services.InventoryUsecase{
	return &inventoryUsecase{
		invRepo: invRepo,
	}
}