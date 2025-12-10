package auth_repository

import (
	auth_dto "go_service/core/modules/auth/dto"
	app "go_service/core/util/error"
	"go_service/core/util/executor"
)

type AuthRepository struct {
	exec executor.Executor
}

func NewAuthRepoistory(executor executor.Executor) *AuthRepository {
	return &AuthRepository{
		exec: executor,
	}
}

func (r *AuthRepository) GetUserById(id string) (auth_dto.UserDto, error) {
	var user auth_dto.UserDto
	query := `SELECT id, email, password, created_at
              FROM "user" 
              WHERE id = $1`
	err := r.exec.QueryRow(query, id).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	if err != nil {
		return user, app.NotFound("usuario não encontrado com esse id")
	}
	return user, nil
}

func (r *AuthRepository) GetUserByEmail(email string) (auth_dto.UserDto, error) {
	var user auth_dto.UserDto
	query := `SELECT id, email, password, created_at
              FROM "user" 
              WHERE email = $1`
	err := r.exec.QueryRow(query, email).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	if err != nil {
		return user, app.NotFound("usuario não encontrado com esse email")
	}
	return user, nil
}
