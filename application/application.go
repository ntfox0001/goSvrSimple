package application

import (
	"goSvrLib/userSystem"
	"os"
)

type Application struct {
	c      chan os.Signal
	usrSev userSystem.UserService
}

func NewApplication() *Application {
	return &Application{}
}

func (a *Application) Initial() error {
	params := userSystem.UserServiceParams{
		Listenip:
	}
	a.usrSev = userSystem.NewUserService()

	return nil
}

func (a *Application) Run() {
runable:
	for {
		select {
		case <-a.c:
			break runable
		}
	}
}
func (a *Application) Release() {

}
