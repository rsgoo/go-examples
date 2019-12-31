package dao

import (
	"x-bookstore/model"
	"x-bookstore/utils"
)

//插入订单
func AddOrder(order *model.Order) (error) {
	//写sql语句
	sql := "insert into orders(id, create_time, total_count, total_amount, state, user_id) values(?,?,?,?,?,?)"
	//执行
	_, err := utils.DB.Exec(sql, order.OrderID, order.CreateTime, order.TotalCount, order.TotalAmount, order.State, order.UserID)

	if err != nil {
		return err
	}
	return nil
}

//GetOrders 获取数据库中所有的订单
func GetOrders() ([]*model.Order, error) {
	//写sql语句
	sql := "select id,create_time,total_count,total_amount,state,user_id from orders order by ctime desc"
	//执行
	rows, err := utils.DB.Query(sql)
	if err != nil {
		return nil, err
	}
	var orders []*model.Order
	for rows.Next() {
		order := &model.Order{}
		rows.Scan(&order.OrderID, &order.CreateTime, &order.TotalCount, &order.TotalAmount, &order.State, &order.UserID)
		orders = append(orders, order)
	}
	return orders, nil
}

//获取登录用户的订单
func GetMyOrders(userId int) ([]*model.Order, error) {
	//写sql语句
	sql := "select id,create_time,total_count,total_amount,state,user_id from orders where user_id = ?"
	//执行
	rows, err := utils.DB.Query(sql, userId)
	if err != nil {
		return nil, err
	}
	var orders []*model.Order
	for rows.Next() {
		order := &model.Order{}
		rows.Scan(&order.OrderID, &order.CreateTime, &order.TotalCount, &order.TotalAmount, &order.State, &order.UserID)
		orders = append(orders, order)
	}
	return orders, nil
}

//更新订单的状态，即发货和收货
//state:0 未发货 1 已发货 2 交易完成
func UpdateOrderState(orderID string, state int64) error {
	//写sql语句
	sql := "update orders set state = ? where id = ?"
	//执行
	_, err := utils.DB.Exec(sql, state, orderID)
	if err != nil {
		return err
	}
	return nil
}