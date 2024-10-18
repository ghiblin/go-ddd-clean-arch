package controller

import (
	"ddd-clean-arch/internal/application/dto"
	"ddd-clean-arch/internal/application/usecase"
	"ddd-clean-arch/internal/domain/event"
	"encoding/json"
	"net/http"
)

type OrderController struct {
	createOrderUseCase         *usecase.CreateOrderUseCase
	processOrderPaymentUseCase *usecase.ProcessOrderPaymentUseCase
	stockMovementUseCase       *usecase.StockMovementUseCase
	sendOrderEmailUseCase      *usecase.SendOrderEmailUseCase
}

func NewOrderController(createOrderUseCase *usecase.CreateOrderUseCase, processOrderPaymentUseCase *usecase.ProcessOrderPaymentUseCase,
	stockMovementUseCase *usecase.StockMovementUseCase,
	sendOrderEmailUseCase *usecase.SendOrderEmailUseCase) *OrderController {
	return &OrderController{
		createOrderUseCase:         createOrderUseCase,
		processOrderPaymentUseCase: processOrderPaymentUseCase,
		stockMovementUseCase:       stockMovementUseCase,
		sendOrderEmailUseCase:      sendOrderEmailUseCase,
	}
}

func (c *OrderController) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var requestData dto.CreateOrderDTO
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	err = c.createOrderUseCase.Execute(r.Context(), requestData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (c *OrderController) ProcessOrderPayment(w http.ResponseWriter, r *http.Request) {
	var body event.OrderCreatedEvent
	json.NewDecoder(r.Body).Decode(&body)
	err := c.processOrderPaymentUseCase.Execute(r.Context(), &body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (c *OrderController) StockMovement(w http.ResponseWriter, r *http.Request) {
	var body event.OrderCreatedEvent
	json.NewDecoder(r.Body).Decode(&body)
	err := c.stockMovementUseCase.Execute(r.Context(), &body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (c *OrderController) SendOrderEmail(w http.ResponseWriter, r *http.Request) {
	var body event.OrderCreatedEvent
	json.NewDecoder(r.Body).Decode(&body)
	err := c.sendOrderEmailUseCase.Execute(r.Context(), &body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusCreated)
}
