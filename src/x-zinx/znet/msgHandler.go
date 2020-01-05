package znet

import (
	"x-zinx/ziface"
	"fmt"
)

// 消息处理模块的实现
type MsgHandler struct {
	//存放每个msgId 所对应的处理方法
	ApiList map[uint32]ziface.IRouter
}

func NewMsgHandler() *MsgHandler {
	return &MsgHandler{
		ApiList: make(map[uint32]ziface.IRouter),
	}
}

//调度 | 执行对应的Router消息处理方法
func (mh *MsgHandler) DoMsgHandler(request ziface.IRequest) {
	//从request中找到msgID
	msgID := request.GetMsgID()
	routerHandler, ok := mh.ApiList[msgID]
	if !ok {
		fmt.Println("api msgID = ", request.GetMsgID(), " is not found! need register")
	}
	//根据 msgID, 调度对应的router业务集合
	routerHandler.PreHandle(request)
	routerHandler.Handle(request)
	routerHandler.PreHandle(request)

}

//为消息添加具体的处理逻辑
func (mh *MsgHandler) AddRouter(msgID uint32, router ziface.IRouter) {
	//判断 当前msgID绑定的api处理方法是否已经存在
	if _, ok := mh.ApiList[msgID]; ok {
		//已经注册
		panic("repeat api, msgID = " + string(msgID))
	}

	//添加 msgID与api的绑定关系
	mh.ApiList[msgID] = router
	fmt.Println("success add api msgID = ", msgID)

}
