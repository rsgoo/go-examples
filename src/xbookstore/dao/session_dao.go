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

//获取session
func GetSession(sessID string) (*model.Session, error) {
	sqlStr := "select session_id, user_id, username from sessions where session_id = ?"
	row := utils.DB.QueryRow(sqlStr, sessID)

	var session = &model.Session{}
	row.Scan(&session.SessionID, &session.UserID, &session.UserName)
	return session, nil
}
