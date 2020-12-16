package data


type OrderSaveRepository struct {


}

func NewOrderSaveRepository() *OrderSaveRepository{
	return &OrderSaveRepository{}
}


func(this *OrderSaveRepository) OrderSave(orderId int64)(int, error){
	// db
	return 1, nil
}