package controller

import (
	"net/http"
	"fmt"
	"xbookstore/dao"
	"xbookstore/model"
	"xbookstore/utils"
	"html/template"
	"strconv"
	"encoding/json"
)

func AddBook2Cart(w http.ResponseWriter, r *http.Request) {
	//判断是否登录
	flag, session := dao.IsLogin(r)

	if flag {
		//获取图书id
		bookID := r.FormValue("bookId")
		fmt.Println("要添加图书的id是：", bookID)

		//获取图书信息
		book, _ := dao.GetBookById(bookID)

		//判断数据库中是否有当前用户的购物车
		cart, _ := dao.GetCartByUserID(session.UserID)

		//用户有购物车
		if cart != nil {
			//当前登录的用户已经有购物车
			//1：判断当前购物车中是否有这本图书
			cartItem, _ := dao.GetCartItemByBookIDAndCartID(bookID, cart.CartID)
			if cartItem != nil {
				//将该用户的购物车中的购物项的图书数量+1处理

				//购物车中的购物项
				cts := cart.CartItems

				for _, v := range cts {
					//当前的购物项的图书与购物车中已有购物项图书一样
					if v.Book.ID == cartItem.Book.ID {
						fmt.Println("更新图书数量...")
						//将购物项中的图书的数量加1
						v.Count += 1
						//更新数据库中该购物项的图书的数量
						dao.UpdateBookCount(v)
					}
				}

			} else {
				fmt.Println("当前购物车中没有该图书对应的购物项")
				//购物车中的购物项还没有该图书，此时需要创建一个购物项并添加到数据表中

				//创建购物项
				cartItem := &model.CartItem{
					Book:   book,
					Count:  1,
					CartID: cart.CartID,
				}

				//添加到购物项切片中
				cart.CartItems = append(cart.CartItems, cartItem)

				//将新创建的购物项添加到数据库中
				dao.AddCartItem(cartItem)
			}

			//更新购物车中的总数量和总金额
			dao.UpdateCart(cart)

		} else {
			//当前用户还没有购物车，需要创建一个购物车并添加到数据表中

			//创建购物车
			cartID := utils.CreateUUID()
			cart := &model.Cart{
				CartID: cartID,
				UserID: session.UserID,
			}

			//购物车中的购物项
			var cartItems []*model.CartItem

			//创建购物项
			cartItem := &model.CartItem{
				Book:   book,
				Count:  1,
				CartID: cartID,
			}

			//将购物项添加到切片中
			cartItems = append(cartItems, cartItem)

			//3将切片设置到cart中
			cart.CartItems = cartItems

			//将购物车保存到数据库中
			dao.AddCart(cart)
		}
		w.Write([]byte("您刚刚将《" + book.Title + "》添加到了购物车！"))
	} else {
		w.Write([]byte("请先登录！"))
	}
}

//获取购物车信息
func GetCartInfo(w http.ResponseWriter, r *http.Request) {
	_, session := dao.IsLogin(r)

	//huo去用户id
	userID := session.UserID

	//从数据表中获取对应的购物车
	cart, _ := dao.GetCartByUserID(userID)
	//存在购物车
	if cart != nil {
		session.Cart = cart
		t := template.Must(template.ParseFiles("views/pages/cart/cart.html"))
		t.Execute(w, session)
	} else {
		//该用户还没有购物车
		//解析模板文件
		t := template.Must(template.ParseFiles("views/pages/cart/cart.html"))
		//执行
		t.Execute(w, session)
	}
}

//清空购物车
func DeleteCart(w http.ResponseWriter, r *http.Request) {
	cartId := r.FormValue("cartId")
	dao.DeleteCartByCartID(cartId)
	//调用GetCartInfo函数再次查询购物车信息
	GetCartInfo(w, r)
}

//删除购物车中的购物项
func DeleteCartItem(w http.ResponseWriter, r *http.Request) {
	cartItemId := r.FormValue("cartItemId")
	iCartItemId, _ := strconv.ParseInt(cartItemId, 10, 64)

	//获取登录用户的session
	_, session := dao.IsLogin(r)

	//获取用户的id
	userID := session.UserID

	//获取该用户的购物车
	cart, _ := dao.GetCartByUserID(userID)
	cartItems := cart.CartItems

	for k, v := range cartItems {
		if v.CartItemId == iCartItemId {

			//将购物车中的购物想移除
			cartItems = append(cartItems[:k], cartItems[k+1:]...)
			//将删除购物项之后的切片再次赋给购物车中的切片
			cart.CartItems = cartItems
			//将当前购物项从数据库中删除
			dao.DeleteCartItemByID(cartItemId)
		}
	}
	dao.UpdateCart(cart)

	GetCartInfo(w, r)
}

//更新购物车中的购物项目
func UpdateCartItem(w http.ResponseWriter, r *http.Request) {
	//获取要更新的购物项的id
	cartItemID := r.FormValue("cartItemId")
	//将购物项的id转换为int64
	iCartItemId, _ := strconv.ParseInt(cartItemID, 10, 64)

	//获取用户输入的图书的数量
	bookCount := r.FormValue("bookCount")
	iBookCount, _ := strconv.ParseInt(bookCount, 10, 64)

	//获取session
	_, session := dao.IsLogin(r)

	//获取用户的id
	userID := session.UserID

	//获取该用户的购物车
	cart, _ := dao.GetCartByUserID(userID)

	//获取购物车中的所有的购物项
	cartItems := cart.CartItems

	//遍历得到每一个购物项
	for _, v := range cartItems {
		//寻找要更新的购物项
		if v.CartItemId == iCartItemId {
			//这个就是我们要更新的购物项
			//将当前购物项中的图书的数量设置为用户输入的值
			v.Count = iBookCount
			//更新数据库中该购物项的图书的数量和金额小计
			dao.UpdateBookCount(v)
		}
	}

	//更新购物车中的图书的总数量和总金额
	dao.UpdateCart(cart)

	//调用获取购物项信息的函数再次查询购物车信息
	cart, _ = dao.GetCartByUserID(userID)

	// GetCartInfo(w, r)
	//获取购物车中图书的总数量
	totalCount := cart.TotalCount

	//获取购物车中图书的总金额
	totalAmount := cart.TotalAmount

	var amount float64
	//获取购物车中更新的购物项中的金额小计
	cIs := cart.CartItems
	for _, v := range cIs {
		if iCartItemId == v.CartItemId {
			//这个就是我们寻找的购物项，此时获取当前购物项中的金额小计
			amount = v.Amount
		}
	}
	//创建Data结构
	data := model.Data{
		Amount:      amount,
		TotalAmount: totalAmount,
		TotalCount:  totalCount,
	}
	//将data转换为json字符串
	jsonData, _ := json.Marshal(data)
	//响应到浏览器
	w.Write(jsonData)
}
