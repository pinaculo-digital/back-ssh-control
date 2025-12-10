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

func NewContainer() *Container {
	return &Container{
		Services:    &Services{},
		Middlewares: &Middlewares{},
		Controllers: &Controllers{},
		Groups:      &RouteGroups{},
	}
}

func (c *Container) Init() {
	c.initCoreServices()
	c.initModuleServices()
	c.initMiddlewares()
	c.initControllers()
}

func (c *Container) initCoreServices() {
	var err error

	c.Services.Redis, err = redis_service.NewRedisService()
	if err != nil {
		panic(err)
	}
}

func (c *Container) initModuleServices() {
	c.Services.Auth = auth_service.NewAuthService(c.Services.Redis)
	c.Services.SSH = ssh_service.NewSSHService()
}

func (c *Container) initMiddlewares() {
	c.Middlewares.Auth = auth_middleware.NewAuthMiddleware(c.Services.Auth)
}

func (c *Container) initControllers() {
	c.Controllers.Auth = auth_controller.NewController(c.Services.Auth)
	c.Controllers.SSH = ssh_controller.NewSSHController(c.Services.SSH)
}

func (c *Container) RegisterRoutes(engine *gin.Engine) {
	c.Groups.Public = engine.Group("/")
	c.Groups.Protected = engine.Group("/", c.Middlewares.Auth.JwtGuard)

	// Rotas públicas
	c.Controllers.Auth.Routes(c.Groups.Public)
	c.Controllers.SSH.Routes(c.Groups.Protected)
}
