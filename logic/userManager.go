package logic

import "goSvrLib/selectCase/selectCaseInterface"

type UserManager struct {
	helper selectCaseInterface.ISelectLoopHelper
}

func (um *UserManager) OnInitial(helper selectCaseInterface.ISelectLoopHelper) error {
	um.helper = helper
	// todo 初始化user管理逻辑
	return nil
}
func (um *UserManager) OnRelease() {

}
