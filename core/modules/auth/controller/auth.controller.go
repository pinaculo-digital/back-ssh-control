package auth_controller

import (
	auth_dto "go_service/core/modules/auth/dto"
	auth_middleware "go_service/core/modules/auth/middleware"
	auth_service "go_service/core/modules/auth/service"
	"go_service/core/util/interceptor"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type AuthController struct {
	authService *auth_service.AuthService
	authMid     *auth_middleware.AuthMiddleware
}

func NewController(authService *auth_service.AuthService) *AuthController {
	return &AuthController{
		authService: authService,
		authMid:     auth_middleware.NewAuthMiddleware(authService),
	}
}

func (controller *AuthController) Login(ctx *gin.Context) {

	var body auth_dto.AuthDto
	err := interceptor.ValidateAndExtract(ctx, &body)
	if err != nil {
		interceptor.AppError(ctx, err)
		return
	}

	_, token, err := controller.authService.Login(body)
	if err != nil {
		interceptor.AppError(ctx, err)
		return
	}

	interceptor.AppSuccess(ctx, "login bem sucedido", gin.H{
		"token": token,
	})

}

var X = godotenv.Load()

func (c *AuthController) Routes(g *gin.RouterGroup) *gin.RouterGroup {

	auth := g.Group("/auth", c.authMid.IpValidation)
	{
		auth.POST("/login", c.Login)

	}
	return auth
}
