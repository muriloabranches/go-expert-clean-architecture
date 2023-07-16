package service

import (
	"context"

	"github.com/muriloabranches/Go-Expert-Clean-Architecture/internal/infra/grpc/pb"
	"github.com/muriloabranches/Go-Expert-Clean-Architecture/internal/usecase"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	CreateOrderUseCase  usecase.CreateOrderUseCase
	FindAllOrderUseCase usecase.FindAllOrderUseCase
}

func NewOrderService(createOrderUseCase usecase.CreateOrderUseCase, findAllOrderUseCase usecase.FindAllOrderUseCase) *OrderService {
	return &OrderService{
		CreateOrderUseCase:  createOrderUseCase,
		FindAllOrderUseCase: findAllOrderUseCase,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	dto := usecase.OrderInputDTO{
		ID:    in.Id,
		Price: float64(in.Price),
		Tax:   float64(in.Tax),
	}
	output, err := s.CreateOrderUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}
	return &pb.CreateOrderResponse{
		Id:         output.ID,
		Price:      float32(output.Price),
		Tax:        float32(output.Tax),
		FinalPrice: float32(output.FinalPrice),
	}, nil
}

func (s *OrderService) ListOrders(ctx context.Context, in *pb.Blank) (*pb.OrderList, error) {
	orders, err := s.FindAllOrderUseCase.Execute()
	if err != nil {
		return nil, err
	}

	var ordersResponse []*pb.FindAllOrderResponse
	for _, order := range orders {
		orderResponse := &pb.FindAllOrderResponse{
			Id:         order.ID,
			Price:      float32(order.Price),
			Tax:        float32(order.Tax),
			FinalPrice: float32(order.FinalPrice),
		}

		ordersResponse = append(ordersResponse, orderResponse)
	}

	return &pb.OrderList{Orders: ordersResponse}, nil
}
