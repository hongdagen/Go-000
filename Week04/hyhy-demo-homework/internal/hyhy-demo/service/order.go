package service

import (
	v1 "github.com/hongdagen/Go-000/Week_04/hyhy-demo-homework/api/hyhy-demo/v1"
	"github.com/hongdagen/Go-000/Week_04/hyhy-demo-homework/internal/hyhy-demo/biz"
	"context"
)

type OrderSaveService struct {
	biz *biz.OrderSaveBIZ
}

func NewOrderSaveService(biz *biz.OrderSaveBIZ) *OrderSaveService {
	return &OrderSaveService{biz: biz}
}

func (this *OrderSaveService) OrderSave(ctx context.Context, or *v1.OrderRequest) (*v1.OrderResponse, error){
	id, _ := this.biz.OrderSave(100)
	return &v1.OrderResponse{OrderId: int64(id)}, nil
}
