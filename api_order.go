package mws

//OrderService
type OrderService struct {
	*Client
}

//Orders
func Orders() *OrderService {
	return &OrderService{
		Client: newClient("/Orders/2013-09-01", "2013-09-01"),
	}
}
