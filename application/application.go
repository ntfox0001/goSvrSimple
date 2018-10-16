package application

import (
	"goSvrLib/syncDBSystem"
	"goSvrLib/timerSystem"
	"goSvrLib/userSystem"
	"goSvrSimple/applicationConfig"
	"os"
)

type Application struct {
	c      chan os.Signal
	usrSev *userSystem.UserService
}

func NewApplication(configFile string) *Application {
	applicationConfig.InitApplicationConfig(configFile)
	return &Application{}
}

func (a *Application) Initial() error {
	timerSystem.Instance().Initial()
	syncDBSystem.Instance().Initial()
	a.usrSev = userSystem.NewUserService(applicationConfig.Config.UserService)

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
