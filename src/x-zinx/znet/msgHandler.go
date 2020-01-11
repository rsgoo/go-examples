package znet

import (
	"x-zinx/ziface"
	"fmt"
	"x-zinx/utils"
)

// 消息处理模块的实现
type MsgHandler struct {
	//存放每个msgId 所对应的处理方法
	ApiList map[uint32]ziface.IRouter

	//Worker负责 读取任务的消息队列
	TaskQueue []chan ziface.IRequest

	//业务工作Worker池的数量
	WorkerPoolSize uint32
}

func NewMsgHandler() *MsgHandler {
	return &MsgHandler{
		ApiList:        make(map[uint32]ziface.IRouter),
		TaskQueue:      make([]chan ziface.IRequest, utils.GlobalObject.WorkerPoolSize),
		WorkerPoolSize: utils.GlobalObject.WorkerPoolSize, //从全局配置中获取
	}
}

//调度 | 执行对应的Router消息处理方法
func (mh *MsgHandler) DoMsgHandler(request ziface.IRequest) {
	//从request中找到msgID
	msgID := request.GetMsgID()
	routerHandler, ok := mh.ApiList[msgID]
	if !ok {
		fmt.Println("api msgID = ", request.GetMsgID(), " is not found! need register")
		return
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

//启动一个Worker流程
func (mh *MsgHandler) StartOneWorker(workerID int, taskQueue chan ziface.IRequest) {
	fmt.Println("WorkerID = ", workerID, " is start")
	//不断的阻塞获取消息
	for {
		select {
		//如果有消息过来，出队列是一个客户端的Request, 执行当前Request所绑定的业务
		case request := <-taskQueue:
			mh.DoMsgHandler(request)
		}
	}
}

//将消息交给taskQueue, 由worker进行处理
func (mh *MsgHandler) SendMsgToTaskQueue(request ziface.IRequest) {
	connID := request.GetConnection().GetConnID()
	workerID := connID % mh.WorkerPoolSize
	fmt.Println(
		"Add connID = ", connID,
		"request MsgID = ", request.GetMsgID(),
		"workerID = ", workerID)
	mh.TaskQueue[workerID] <- request

}

//启动一个工作池（开启工作池的动作只发生一次）
func (mh *MsgHandler) StartWorkerPool() {
	//更具workerPoolSize 分别开启Worker, 每个worker用一个Go来承载
	for i := 0; i < int(mh.WorkerPoolSize); i++ {
		//一个worker被启动
		//给当前的worker对应的channel消息队列开辟空间
		mh.TaskQueue[i] = make(chan ziface.IRequest, utils.GlobalObject.MaxWorkerTaskLen)

		//启动当前的worker， 阻塞等待消息从channel中传递过来
		go mh.StartOneWorker(i, mh.TaskQueue[i])
	}
}
