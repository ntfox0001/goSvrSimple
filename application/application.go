package application

import (
	"goSvrLib/userSystem"
	"goSvrSimple/applicationConfig"
	"os"
)

type Application struct {
	c      chan os.Signal
	usrSev userSystem.UserService
}

func NewApplication(configFile string) *Application {
	applicationConfig.InitApplicationConfig(configFile)
	return &Application{}
}

func (a *Application) Initial() error {
	params := userSystem.UserServiceParams{
		Listenip: applicationConfig.Config.UserService.Listenip,
		Port:     applicationConfig.Config.UserService.Port,
		CertFile: applicationConfig.Config.UserService.CertFile,
		KeyFile:  applicationConfig.Config.UserService.KeyFile,
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
