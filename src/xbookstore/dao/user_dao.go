package dao

import (
	"xbookstore/model"
	"xbookstore/utils"
	"fmt"
	"os"
)

//查询用户信息
func CheckUserNameAndPassword(username, password string) (*model.User, error) {
	sqlStr := "select id, username, password, email, ctime,mtime from users where username = ? and password = ?";
	row := utils.DB.QueryRow(sqlStr, username, password)

	user := &model.User{}
	row.Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.Email,
		&user.Ctime,
		&user.Mtime,
	)
	//没有有效的数据
	if user.ID < 0 {
		return nil, nil
	}
	return user, nil

}

//查询用户信息
func CheckUserName(username string) (user *model.User, err error) {
	sqlStr := "select id, username, password, email, ctime,mtime from users where username = ?";
	row := utils.DB.QueryRow(sqlStr, username)

	user = &model.User{}
	row.Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.Email,
		&user.Ctime,
		&user.Mtime,
	)

	return user, nil
}

func SaveUser(username, password, email string) error {
	sqlStr := "insert into users(username, password, email) values (?,?,?)";
	_, err := utils.DB.Exec(sqlStr, username, password, email)
	if err != nil {
		return err
	}
	return nil
}

func GetUserById(id int) {
	sqlStr := "select * from users where id = ?";
	stmt, err := utils.DB.Prepare(sqlStr)
	defer stmt.Close()

	if err != nil {
		fmt.Println("err is :", err)
	}

	row := stmt.QueryRow(id)
	var user = &model.User{}
	row.Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.Email,
		&user.Ctime,
		&user.Mtime,
	)
	fmt.Println(user)
	os.Exit(1)

}
