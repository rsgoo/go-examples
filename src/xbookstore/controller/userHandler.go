package controller

import (
	"net/http"
	"xbookstore/dao"
	"html/template"
	"fmt"
	"os"
	"xbookstore/model"
	"xbookstore/utils"
)

//处理用户登录的函数
func Login(w http.ResponseWriter, r *http.Request) {

	//判断是否已经登录
	flag, _ := dao.IsLogin(r)
	if flag {
		//已经登录
		//去首页
		GetPageBooksByPrice(w, r)
	} else {
		//获取用户名和密码
		username := r.PostFormValue("username")
		password := r.PostFormValue("password")

		user, _ := dao.CheckUserNameAndPassword(username, password)
		if user.ID > 0 {

			//用户名和密码正确
			//创建一个Session
			sess := &model.Session{
				SessionID: utils.CreateUUID(),
				UserName:  user.Username,
				UserID:    user.ID,
			}

			//将Session保存到数据库中
			dao.AddSession(sess)

			//创建一个cookie, 让其与Session相关联
			cookie := http.Cookie{
				Name:     "user",
				Value:    sess.SessionID,
				HttpOnly: true,
			}

			http.SetCookie(w, &cookie)

			t := template.Must(template.ParseFiles("views/pages/user/login_success.html"))
			t.Execute(w, user)
		} else {
			t := template.Must(template.ParseFiles("views/pages/user/login.html"))
			t.Execute(w, "用户名或密码不正确！")
		}
	}

}

func Logout(w http.ResponseWriter, r *http.Request) {
	//获取cookie
	cookie, _ := r.Cookie("user")
	if cookie != nil {

		//获取cookie的value
		cookieValue := cookie.Value

		//删除本次会话
		dao.DelSession(cookieValue)

		//设置cookie失效
		cookie.MaxAge = -1
		http.SetCookie(w, cookie)
	}

	//去首页
	GetPageBooksByPrice(w, r)
}

func Register(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	email := r.PostFormValue("email")
	user, _ := dao.CheckUserName(username)
	if user.ID < 1 {
		err := dao.SaveUser(username, password, email)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		t := template.Must(template.ParseFiles("views/pages/user/register_success.html"))
		t.Execute(w, "")
	} else {
		t := template.Must(template.ParseFiles("views/pages/user/register.html"))
		t.Execute(w, "用户名已存在！")
	}
}

func CheckUserName(w http.ResponseWriter, r *http.Request) {
	//获取用户输入的用户名
	username := r.PostFormValue("username")

	if len(username) < 1 {
		w.Write([]byte("用户名不能为空！"))
		return
	}

	//调用user_dao中验证用户名和密码的方法
	user, _ := dao.CheckUserName(username)
	if user.ID > 0 {
		//用户名已存在
		w.Write([]byte("用户名已存在！"))
	} else {
		//用户名可用
		w.Write([]byte("<font style='color:green'>用户名可用！</font>"))
	}
}
