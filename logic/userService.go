package logic

import "goSvrLib/network"

type UserService struct {
	server *network.Server
}

func (us *UserService) OnInitial(server *network.Server) error {
	us.server = server
	// todo 路由注册，提供静态页面服务
	return nil
}
func (us *UserService) OnRelease() {

}
