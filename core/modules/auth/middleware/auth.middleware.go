package auth_middleware

import (
	auth_service "go_service/core/modules/auth/service"
	"os"
	"strconv"

	text "go_service/core/util/debug"
	"go_service/core/util/interceptor"
	guard_util "go_service/core/util/jwt"

	"github.com/9ssi7/turnstile"
	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	authService *auth_service.AuthService
}

func NewAuthMiddleware(authService *auth_service.AuthService) *AuthMiddleware {
	return &AuthMiddleware{
		authService: authService,
	}
}

func (c *AuthMiddleware) JwtGuard(ctx *gin.Context) {

	bearerToken := ctx.Request.Header.Get("Authorization")

	token, err := guard_util.GetJwtInfo(bearerToken)

	if err != nil {
		interceptor.AppUnauthorized(ctx, "Token invalido")
		return
	}

	userId, exist := token["id"].(string)
	if !exist {
		interceptor.AppUnauthorized(ctx, "Token invalido")
		return
	}

	user, err := c.authService.GetUserById(userId)
	if err != nil {
		interceptor.AppUnauthorized(ctx, "Token valido, porém usuário não existe")
		return
	}

	ctx.Set("userId", userId)
	ctx.Set("user", user)
}

func (c *AuthMiddleware) SfuApiKey(ctx *gin.Context) {

	key := ctx.Request.Header.Get("SFU_API_KEY")

	envApiKey := os.Getenv("SFU_API_KEY")
	if key != envApiKey {
		interceptor.AppUnauthorized(ctx, "Api key invalida")
	}

}

func (c *AuthMiddleware) IpValidation(ctx *gin.Context) {
	ip := ctx.ClientIP()

	text.Println("Tentativa de acesso por: ", ip)
	limit, err := strconv.Atoi(os.Getenv("LIMIT_TRY_ACCESS"))
	if err != nil {
		limit = 10
	}

	tryAcessCount, err := c.authService.GetRedisService.GetIpTryAccess(ip)
	if err == nil && tryAcessCount > limit {
		interceptor.AppForbidden(ctx, "seu ip foi bloqueado por exceder o limite de tentativas de acesso")
		return
	}
	c.authService.GetRedisService.AddIpTryAccess(ip)

}

func (c *AuthMiddleware) CaptchaVerify(ctx *gin.Context) {
	ip := ctx.ClientIP()

	tokenCaptcha := ctx.GetHeader("TOKEN_CAPTCHA")

	if tokenCaptcha == "" {
		interceptor.AppForbidden(ctx, "Para acessar a autenticação, resolva o captcha e envie no cabeçalho TOKEN_CAPTCHA")
		return
	}

	captchaKey := os.Getenv("CAPTCHA_KEY")

	srv := turnstile.New(turnstile.Config{
		Secret: captchaKey,
	})

	ok, err := srv.Verify(ctx, tokenCaptcha, ip)

	if err != nil {
		interceptor.AppInternalServerError(ctx, err.Error())
		return
	}

	if !ok {
		interceptor.AppForbidden(ctx, "token do captcha incorreto ou expirado")
		return
	}

}
