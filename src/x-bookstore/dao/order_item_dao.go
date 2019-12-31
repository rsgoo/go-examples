package dao

import (
	"x-bookstore/model"
	"x-bookstore/utils"
)

//AddOrderItem 向数据表中插入订单项
func AddOrderItem(orderItem *model.OrderItem) error {
	//写sql语句
	sql := "insert into order_items(count,amount,title,author,price,img_path,order_id) values(?,?,?,?,?,?,?)"
	//执行
	_, err := utils.DB.Exec(sql, orderItem.Count, orderItem.Amount, orderItem.Title, orderItem.Author, orderItem.Price, orderItem.ImgPath, orderItem.OrderID)
	if err != nil {
		return err
	}
	return nil
}

//根据订单号获取订单详情
func GetOrderItemsByOrderId(orderId string) ([]*model.OrderItem, error) {
	sql := "select id,count,amount,title,author,price,img_path,order_id from order_items where order_id = ?"

	rows, err := utils.DB.Query(sql, orderId)

	if err != nil {
		return nil, err
	}

	var orderItems []*model.OrderItem

	for rows.Next() {
		orderItem := &model.OrderItem{}
		rows.Scan(orderItem.OrderID)
		rows.Scan(&orderItem.OrderItemID, &orderItem.Count, &orderItem.Amount, &orderItem.Title, &orderItem.Author, &orderItem.Price, &orderItem.ImgPath, &orderItem.OrderID)
		//添加到切片中
		orderItems = append(orderItems, orderItem)
	}

	return orderItems, nil
}
