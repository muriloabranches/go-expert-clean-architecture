package usecase

import "github.com/muriloabranches/Go-Expert-Clean-Architecture/internal/entity"

type FindAllOrderOutputDTO struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

type FindAllOrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewFindAllOrderUseCase(OrderRepository entity.OrderRepositoryInterface) *FindAllOrderUseCase {
	return &FindAllOrderUseCase{
		OrderRepository: OrderRepository,
	}
}

func (uc *FindAllOrderUseCase) Execute() ([]FindAllOrderOutputDTO, error) {
	orders, err := uc.OrderRepository.FindAll()
	if err != nil {
		return nil, err
	}

	var dtos []FindAllOrderOutputDTO
	for _, order := range orders {
		dtos = append(dtos, FindAllOrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		})
	}

	return dtos, nil
}
