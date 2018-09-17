package application

import (
	"goSvrLib/userSystem"
	"os"
)

type Application struct {
	c         chan os.Signal
	usrSystem userSystem.UserSystem
}

func NewApplication() *Application {
	return &Application{}
}

func (a *Application) Initial() error {
	a.usrSystem = userSystem.NewUserSystem()
	
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
