package ssh_service

import (
	"fmt"
	ssh_dto "go_service/core/modules/ssh/dto"
)

type SSHService struct {
}

func NewSSHService() *SSHService {
	return &SSHService{}
}

func (service *SSHService) Implement(data ssh_dto.ImplementDTO) (string, error) {

	branch := data.Branch

	script := fmt.Sprintf("cd scripts & ./implement_sfu.sh %s", branch)

	return service.sendCommandAndPrintResult(script)

}
