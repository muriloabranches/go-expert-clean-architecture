//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"github.com/google/wire"
	"github.com/muriloabranches/Go-Expert-Clean-Architecture/internal/entity"
	"github.com/muriloabranches/Go-Expert-Clean-Architecture/internal/event"
	"github.com/muriloabranches/Go-Expert-Clean-Architecture/internal/infra/database"
	"github.com/muriloabranches/Go-Expert-Clean-Architecture/internal/infra/web"
	"github.com/muriloabranches/Go-Expert-Clean-Architecture/internal/usecase"
	"github.com/muriloabranches/Go-Expert-Clean-Architecture/pkg/events"
)

var setOrderRepositoryDependency = wire.NewSet(
	database.NewOrderRepository,
	wire.Bind(new(entity.OrderRepositoryInterface), new(*database.OrderRepository)),
)

var setEventDispatcherDependency = wire.NewSet(
	events.NewEventDispatcher,
	event.NewOrderCreated,
	wire.Bind(new(events.EventInterface), new(*event.OrderCreated)),
	wire.Bind(new(events.EventDispatcherInterface), new(*events.EventDispatcher)),
)

var setOrderCreatedEvent = wire.NewSet(
	event.NewOrderCreated,
	wire.Bind(new(events.EventInterface), new(*event.OrderCreated)),
)

func NewCreateOrderUseCase(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *usecase.CreateOrderUseCase {
	wire.Build(
		setOrderRepositoryDependency,
		setOrderCreatedEvent,
		usecase.NewCreateOrderUseCase,
	)
	return &usecase.CreateOrderUseCase{}
}

func NewFindAllOrderUseCase(db *sql.DB) *usecase.FindAllOrderUseCase {
	wire.Build(
		setOrderRepositoryDependency,
		usecase.NewFindAllOrderUseCase,
	)
	return &usecase.FindAllOrderUseCase{}
}

func NewWebOrderHandler(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *web.WebOrderHandler {
	wire.Build(
		setOrderRepositoryDependency,
		setOrderCreatedEvent,
		web.NewWebOrderHandler,
	)
	return &web.WebOrderHandler{}
}
