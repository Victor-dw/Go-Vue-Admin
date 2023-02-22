package shop

// ShopService
type ShopService interface {
}

type Shop struct{}

// NewShopService 构造函数
func NewShopService() ShopService {
	return Shop{}
}
