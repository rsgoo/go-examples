package dao

import (
	"xbookstore/model"
	"xbookstore/utils"
)

//向数据库中添加Session
func AddSession(sess *model.Session) error {
	sqlStr := "insert into sessions(session_id,username,user_id) values(?,?,?)"

	_, err := utils.DB.Exec(sqlStr, sess.SessionID, sess.UserName, sess.UserID)

	if err != nil {
		return err
	}

	return nil
}

//删除session
func DelSession(sessID string) error {
	sqlStr := "delete from sessions where session_id = ?"
	_, err := utils.DB.Exec(sqlStr, sessID)

	if err != nil {
		return err
	}

	return nil
}
