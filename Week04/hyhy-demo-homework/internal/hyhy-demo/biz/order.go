package biz

import "github.com/hongdagen/Go-000/Week_04/hyhy-demo-homework/internal/hyhy-demo/data"

type OrderSaveRepo interface {
	OrderSave(orderId int64)(int, error)
}

type OrderSaveBIZ struct {
	repo OrderSaveRepo
}

func NewOrderSaveBIZ(repo *data.OrderSaveRepository) *OrderSaveBIZ{
	return &OrderSaveBIZ{repo: repo}
}

func(this *OrderSaveBIZ) OrderSave(orderId int64)(int, error){
	return this.repo.OrderSave(orderId)
}