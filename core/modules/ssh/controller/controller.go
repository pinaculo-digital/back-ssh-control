package ssh_controller

import (
	ssh_dto "go_service/core/modules/ssh/dto"
	ssh_service "go_service/core/modules/ssh/service"
	"go_service/core/util/interceptor"

	"github.com/gin-gonic/gin"
)

type SSHController struct {
	sshService *ssh_service.SSHService
}

func NewSSHController(service *ssh_service.SSHService) *SSHController {
	return &SSHController{
		sshService: service,
	}
}

func (controller *SSHController) Implement(ctx *gin.Context) {

	var data ssh_dto.ImplementDTO

	err := interceptor.ValidateAndExtract(ctx, &data)
	if err != nil {
		interceptor.AppError(ctx, err)
		return
	}

	result, err := controller.sshService.Implement(data)

	interceptor.AppSuccess(ctx, "Comando bem sucedido", gin.H{
		"stdout": result,
		"stderr": err.Error(),
	})
}

func (c *SSHController) Routes(g *gin.RouterGroup) *gin.RouterGroup {

	auth := g.Group("/ssh")
	{
		auth.PUT("/implement", c.Implement)

	}
	return auth
}
