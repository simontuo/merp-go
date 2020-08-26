package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/registry"
	"github.com/simontuo/merp-go/helper/response"
	reg "github.com/simontuo/merp-go/lib/registry"
)

type ServiceHandler struct {
	Registry registry.Registry
}

func (s *ServiceHandler) List() gin.HandlerFunc {
	return func(c *gin.Context) {
		s.Registry = reg.Etcd()
		services, err := s.Registry.ListServices()
		if err != nil {
			response.Err(c, http.StatusInternalServerError, err.Error())
			return
		}
		fmt.Println(services)
		var liveServices []*registry.Service

		for _, v := range services {

			liveServices = append(liveServices, v)
		}

		data := make(map[string]interface{})
		data["items"] = liveServices
		data["total"] = len(liveServices)

		response.OK(c, data)
	}
}

func (s *ServiceHandler) Info() gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Param("name")

		s.Registry = reg.Etcd()
		service, err := s.Registry.GetService(name)
		if err != nil {
			response.Err(c, http.StatusInternalServerError, err.Error())
			return
		}

		response.OK(c, service[0])
	}
}
