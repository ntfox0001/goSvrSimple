package logic

import (
	"goSvrLib/selectCase/selectCaseInterface"
)

type User struct {
	helper selectCaseInterface.ISelectLoopHelper
}

func (u *User) Initial(helper selectCaseInterface.ISelectLoopHelper) error {
	u.helper = helper
	// todo 初始化user
	return nil
}
func (u *User) Release() {

}
