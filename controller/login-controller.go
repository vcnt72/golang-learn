package controller

import (
	"github.com/vcnt72/golang-learn/dto"
	"github.com/vcnt72/golang-learn/service"
	"github.com/gin-gonic/gin"
)


type LoginController interface {
	Login(ctx *gin.Context) string
}

type loginController struct {
	loginService service.LoginService
	jwtService service.JWTService
}

func NewLoginController(loginService service.LoginService, jwtService service.JWTService) LoginController{
	return &loginController{
		loginService,
		jwtService,
	}
}

func (loginCtr *loginController) Login(ctx *gin.Context) string {
	var credentials dto.Credentials
	err := ctx.BindJSON(&credentials)
	if err != nil {
		return ""
	}

	isAuthenticated := loginCtr.loginService.Login(credentials.Email, credentials.Password)
	if isAuthenticated {
		return loginCtr.jwtService.GenerateToken(credentials.Email, true)
	}
	return ""
}