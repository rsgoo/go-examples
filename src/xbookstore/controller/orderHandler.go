package controller

import (
	"net/http"
	"xbookstore/dao"
	"xbookstore/utils"
	"xbookstore/model"
	"time"
	"html/template"
)

func Checkout(w http.ResponseWriter, r *http.Request) {
	//获取session
	_, session := dao.IsLogin(r)
	userId := session.UserID

	//获取购物车
	cart, _ := dao.GetCartByUserID(userId)

	orderId := utils.CreateUUID()

	//创建订单
	order := &model.Order{
		OrderID:     orderId,
		CreateTime:  time.Now().Format("2006-01-02 15:04:05"),
		TotalCount:  cart.TotalCount,
		TotalAmount: cart.TotalAmount,
		State:       0,
		UserID:      int64(userId),
	}

	//保存订单到数据表
	dao.AddOrder(order)

	//保存订单项目
	cartItems := cart.CartItems
	for _, value := range cartItems {
		orderItem := &model.OrderItem{
			Count:   value.Count,
			Amount:  value.Amount,
			Title:   value.Book.Title,
			Author:  value.Book.Author,
			Price:   value.Book.Price,
			ImgPath: value.Book.ImgPath,
			OrderID: orderId,
		}
		//保存到数据表
		dao.AddOrderItem(orderItem)

		//更新当前当前购物项中的库存和销量
		book := value.Book
		book.Sales = book.Sales + int(value.Count)
		book.Stock = book.Stock - int(value.Count)
		//更新图书的信息
		dao.UpdateBook(book)
	}

	//结完账后清空购物车
	dao.DeleteCartByCartID(cart.CartID)

	//将订单号设置到session中
	session.OrderID = orderId

	//解析模板
	t := template.Must(template.ParseFiles("views/pages/cart/checkout.html"))
	//执行
	t.Execute(w, session)
}

//GetOrders 获取所有订单
func GetOrders(w http.ResponseWriter, r *http.Request) {
	//调用dao中获取所有订单的函数
	orders, _ := dao.GetOrders()
	//解析模板
	t := template.Must(template.ParseFiles("views/pages/order/order_manager.html"))

	//for _, value := range orders {
	//	fmt.Println(value)
	//}
	//os.Exit(1)

	//执行
	t.Execute(w, orders)
}

func GetOrderInfo(w http.ResponseWriter, r *http.Request) {
	orderId := r.FormValue("orderId")
	ordersItems, _ := dao.GetOrderItemsByOrderId(orderId)
	//解析模板
	t := template.Must(template.ParseFiles("views/pages/order/order_info.html"))
	//执行
	t.Execute(w, ordersItems)
}

//获取登录用户订单
func GetMyOrders(w http.ResponseWriter, r *http.Request) {
	_, session := dao.IsLogin(r)
	orders, _ := dao.GetMyOrders(session.UserID)

	//将订单设置到session中
	session.Orders = orders
	//解析模板
	t := template.Must(template.ParseFiles("views/pages/order/order.html"))
	//执行
	t.Execute(w, session)
}

//发货
func SendOrder(w http.ResponseWriter, r *http.Request) {
	//获取订单id
	orderId := r.FormValue("orderId")
	dao.UpdateOrderState(orderId, 1)
	GetOrders(w, r)
}

//收到货
func TakeOrder(w http.ResponseWriter, r *http.Request)  {
	//获取订单id
	orderId := r.FormValue("orderId")
	dao.UpdateOrderState(orderId, 2)
	GetMyOrders(w, r)
}
