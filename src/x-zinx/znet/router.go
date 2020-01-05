package znet

import "x-zinx/ziface"

type BaseRouter struct {
}

//在处理conn业务之前的钩子方法hook
func (b *BaseRouter) PreHandle(request ziface.IRequest) {

}

//在处理conn业务的主方法hook
func (b *BaseRouter) Handle(request ziface.IRequest) {

}

//在处理conn业务之后的钩子方法hook
func (b *BaseRouter) PostHandle(request ziface.IRequest) {

}
