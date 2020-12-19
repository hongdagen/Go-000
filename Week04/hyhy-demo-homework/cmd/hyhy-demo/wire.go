// +build wireinject

package main


import (
	"github.com/google/wire"
	"github.com/hongdagen/Go-000/Week_04/hyhy-demo-homework/internal/hyhy-demo/biz"
	"github.com/hongdagen/Go-000/Week_04/hyhy-demo-homework/internal/hyhy-demo/data"
	"github.com/hongdagen/Go-000/Week_04/hyhy-demo-homework/internal/hyhy-demo/service"
)

func InitOrderSave() *service.OrderSaveService{
	wire.Build(service.NewOrderSaveService,
		biz.NewOrderSaveBIZ,
		data.NewOrderSaveRepository,
	)
	return &service.OrderSaveService{}
}
