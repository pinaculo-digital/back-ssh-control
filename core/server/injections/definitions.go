package container

import (
	auth_controller "go_service/core/modules/auth/controller"
	auth_middleware "go_service/core/modules/auth/middleware"
	auth_service "go_service/core/modules/auth/service"
	redis_service "go_service/core/modules/persistence/redis"
	ssh_controller "go_service/core/modules/ssh/controller"
	ssh_service "go_service/core/modules/ssh/service"

	"github.com/gin-gonic/gin"
)

type Container struct {
	Services    *Services
	Middlewares *Middlewares
	Controllers *Controllers
	Groups      *RouteGroups
}

type Services struct {
	Auth  *auth_service.AuthService
	Redis *redis_service.RedisService
	SSH   *ssh_service.SSHService
}

type Middlewares struct {
	Auth *auth_middleware.AuthMiddleware
}

type Controllers struct {
	Auth *auth_controller.AuthController
	SSH  *ssh_controller.SSHController
}

type RouteGroups struct {
	Public    *gin.RouterGroup
	Protected *gin.RouterGroup
}
