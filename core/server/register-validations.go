package server

import (
	ssh_dto "go_service/core/modules/ssh/dto"
	"go_service/core/util/interceptor"
)

func RegisterValidations() {

	interceptor.RegisterValidation("branch_name", ssh_dto.ValidateEnumBranch, " é do tipo incorreto")

}
