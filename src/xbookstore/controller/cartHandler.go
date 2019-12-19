package controller

import (
	"net/http"
	"fmt"
	"xbookstore/dao"
	"xbookstore/model"
	"xbookstore/utils"
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
		if cart != nil {
			//当前登录的用户已经有购物车

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
	}

}
