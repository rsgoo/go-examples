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
		w.Write([]byte("请先登录:)"))
	}

}
