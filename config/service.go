package config

import (
	"github.com/micro/go-micro/v2/web"
)

var WebService web.Service

//var (
//	UserClient srv_user.UserService
//)
//
//func InitWebService(service web.Service) (web.Service, error) {
//	err := service.Init(
//		web.Action(func(c *cli.Context) {
//			UserClient = srv_user.NewUserService("", service.Options().Service.Client())
//		}),
//	)
//
//	if err != nil {
//		return nil, err
//	}
//
//	return service, nil
//}
func InitWebService(service web.Service) {
	WebService = service
}
