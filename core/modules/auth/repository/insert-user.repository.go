package auth_repository

import (
	auth_dto "go_service/core/modules/auth/dto"

	"github.com/segmentio/ksuid"
)

func (r *AuthRepository) InsertUser(signIn auth_dto.AuthDto) (id string, err error) {

	userID := ksuid.New().String()

	userQuery :=
		`INSERT INTO "user" 
        (id, password, email, created_at) 
        VALUES 
        ($1, $2, $3, CURRENT_TIMESTAMP)`
	_, err = r.exec.Exec(
		userQuery,
		userID,
		signIn.Password,
		signIn.Email,
	)

	return userID, err

}

func (r *AuthRepository) InsertUserInfo(id string, name string) (err error) {
	userInfoQuery :=
		`INSERT INTO "user_info" 
		(id, user_id,name) 
		VALUES 
		($1, $2,$3)`
	_, err = r.exec.Exec(
		userInfoQuery,
		ksuid.New().String(),
		id,
		name,
	)
	return err

}
