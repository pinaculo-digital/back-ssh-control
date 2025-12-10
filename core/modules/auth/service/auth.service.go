package auth_service

import (
	"fmt"
	auth_dto "go_service/core/modules/auth/dto"
	auth_repository "go_service/core/modules/auth/repository"
	"os"

	redis_service "go_service/core/modules/persistence/redis"

	text "go_service/core/util/debug"
	app "go_service/core/util/error"
	"go_service/core/util/executor"
	guard_util "go_service/core/util/jwt"
	"go_service/core/util/transaction"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	authRepo        *auth_repository.AuthRepository
	GetRedisService *redis_service.RedisService
}

func NewAuthService(redisService *redis_service.RedisService) *AuthService {

	godotenv.Load()

	passAdmin := os.Getenv("PASS_ADMIN")
	emailAdmin := os.Getenv("EMAIL_ADMIN")

	service := &AuthService{
		authRepo:        auth_repository.NewAuthRepoistory(executor.NewDBExecutor(nil)),
		GetRedisService: redisService,
	}

	service.CreateUser(auth_dto.AuthDto{
		EmailDto: auth_dto.EmailDto{
			Email: emailAdmin,
		},
		Password: passAdmin,
	})

	return service
}

func (s *AuthService) GetUserById(id string) (user auth_dto.UserDto, err error) {
	user, err = s.authRepo.GetUserById(id)
	return user, err
}
func (s *AuthService) Login(login auth_dto.AuthDto) (user auth_dto.UserDto, token string, err error) {

	user, err = s.ValidateLogin(login)
	if err != nil {
		return user, token, err
	}
	token, err = guard_util.GenerateJwt(user, 360000)
	return user, token, err

}

func (s *AuthService) ValidateLogin(login auth_dto.AuthDto) (auth_dto.UserDto, error) {
	var user auth_dto.UserDto
	user, err := s.authRepo.GetUserByEmail(login.Email)
	if err != nil {
		return user, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password))
	if err != nil {
		return user, app.Forbidden("Senha incorreta")
	}
	return user, nil
}

func (s *AuthService) CreateUser(signIn auth_dto.AuthDto) (user auth_dto.UserDto, token string, err error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(signIn.Password), 12)
	if err != nil {
		return user, token, fmt.Errorf("erro ao gerar hash da senha: %w", err)
	}

	signIn.Password = string(hashedPassword)

	var id string
	err = transaction.RunInTx(func(tx *sqlx.Tx) error {
		authTxRepo := auth_repository.NewAuthRepoistory(executor.NewDBExecutor(tx))

		id, err = authTxRepo.InsertUser(signIn)
		if err != nil {
			return app.Conflict("conflito de cadastro no campo email")
		}

		return err
	})

	if err != nil {
		return user, token, err
	}
	user, _ = s.authRepo.GetUserById(id)

	token, err = guard_util.GenerateJwt(user, 360000)
	return user, token, err
}

// TEST DEVELOPER ONLY
func (s *AuthService) tokenbyemail(email string) (user auth_dto.UserDto, err error) {

	user, err = s.authRepo.GetUserByEmail(email)
	if err != nil {
		return user, err
	}
	token, err := guard_util.GenerateJwt(user, 360000)
	text.Println("token de testeS: ", token)
	return user, err
}
