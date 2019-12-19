package model

//购物车Id
type Cart struct {
	CartID      string      //购物车id
	CartItems   []*CartItem //购物车中得所有图书购物项
	TotalCount  int64       //所有图书的数量
	TotalAmount float64     //所有的费用
	UserID      int         //当前购物车所属的用户
}

func (cart *Cart) GetTotalCount() int64 {
	var totalCount int64 = 0
	for _, v := range cart.CartItems {
		totalCount += v.Count
	}
	return totalCount
}

func (cart *Cart) GetTotalAmount() float64 {
	var totalAmount float64 = 0
	for _, v := range cart.CartItems {
		totalAmount += v.GetAmount()
	}
	return totalAmount
}
