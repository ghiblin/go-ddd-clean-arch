package usecase

import (
	"context"
	"ddd-clean-arch/internal/domain/event"
	"fmt"
)

type SendOrderEmailUseCase struct {
}

func NewSendOrderEmailUseCase() *SendOrderEmailUseCase {
	return &SendOrderEmailUseCase{}
}

func (h *SendOrderEmailUseCase) Execute(ctx context.Context, payload *event.OrderCreatedEvent) error {
	fmt.Println("--- SendOrderEmailUseCase ---")
	fmt.Printf("--- MAIL Order Created: R$ %f \n", payload.TotalPrice)
	return nil
}
