package usecase

import (
	"context"
	"ddd-clean-arch/internal/domain/event"
	"fmt"
)

type StockMovementUseCase struct {
}

func NewStockMovementUseCase() *StockMovementUseCase {
	return &StockMovementUseCase{}
}

func (h *StockMovementUseCase) Execute(ctx context.Context, payload *event.OrderCreatedEvent) error {
	fmt.Println("--- StockMovementUseCase ---")
	for _, item := range payload.Items {
		fmt.Printf("Removing %d items of product %s from stock\n", item.Quantity, item.ProductName)
	}
	return nil
}
