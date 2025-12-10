package auth_dto

import (
	"strings"
	"time"
)

type UserDto struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	Password  string    `json:"-"`
}

type EmailDto struct {
	Email string `json:"email" validate:"required,email" example:"exemplo@mail.com"`
}

func (e EmailDto) GetName() string {
	if e.Email == "" {
		return ""
	}
	if at := strings.Index(e.Email, "@"); at != -1 {
		return e.Email[:at]
	}
	return e.Email
}

type AuthDto struct {
	EmailDto
	Password string `json:"password" validate:"required,min=8" example:"@Senha123"`
}

type MagicAuthDto struct {
	Email       string `json:"email" form:"email" example:"exemplo@mail.com"`
	RedirectUrl string `json:"redirectUrl" form:"redirectUrl" example:"localhost:5173/dashboard"`
}

// User struct matching the TypeScript DTO
type User struct {
	Name      string    `json:"name"`
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Active    bool      `json:"active"`
	AvatarID  string    `json:"avatarId"`
	ProfileID string    `json:"profileId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type QueryRedirect struct {
	User     string `form:"user" json:"user"`
	Redirect string `form:"redirect" json:"redirect"`
}

type TokenDto struct {
	Email      string `json:"email"`
	ExternalId string `json:"external_id"`
	Id         string `json:"id"`
	Name       string `json:"name"`
}
