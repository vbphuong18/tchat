package controllers

import (
	"TChat/pkg/dto"
	"TChat/pkg/services"
	"TChat/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"
)

type AuthenHandler struct {
	authenService services.AuthenService
	v             *validator.Validate
}

func (a *AuthenHandler) Login(ctx *gin.Context) {
	var (
		response   dto.LoginResponse
		request    dto.LoginRequest
		httpStatus = http.StatusOK
	)
	defer func() {
		if httpStatus != http.StatusOK {
			response.ReturnCode = 2
			response.ReturnMessage = "Fail"
		} else {
			response.ReturnCode = 1
			response.ReturnMessage = "Ok"
		}
		ctx.JSON(httpStatus, response)
	}()
	err := ctx.ShouldBind(&request)
	if err != nil {
		httpStatus = http.StatusBadRequest
		return
	}
	err = a.v.Struct(request)
	if err != nil {
		httpStatus = http.StatusBadRequest
		return
	}

	p, err := a.authenService.Login(request.UserName)
	if err != nil {
		httpStatus = http.StatusInternalServerError
		return
	}
	err = utils.CheckPasswordHash(p, request.Password)
	if err != nil {
		httpStatus = http.StatusUnauthorized
		return
	}
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &utils.Claims{
		UserName: request.UserName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(utils.JwtKey)
	if err != nil {
		httpStatus = http.StatusInternalServerError
		return
	}
	ctx.SetCookie("token", tokenString, 5, "/", "", true, false)
}

func NewAuthenHandler(authenService services.AuthenService, v *validator.Validate) *AuthenHandler {
	return &AuthenHandler{authenService: authenService, v: v}
}
