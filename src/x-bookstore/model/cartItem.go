package model

//购物车中得每一个商品购物项
type CartItem struct {
	CartItemId int64   //购物项id
	Book       *Book   //图书信息
	Count      int64   //购买数量
	Amount     float64 //此购物项目中得总金额
	CartID     string  //所属购物车id（订单id）
	Ctime      string
	Mtime      string
}

func (CartItem *CartItem) GetAmount() float64 {
	bookPrice := CartItem.Book.Price
	return bookPrice * float64(CartItem.Count)
}
