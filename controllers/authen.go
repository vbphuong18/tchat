package controllers

import (
	"TChat/dto"
	"TChat/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
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
	err = a.authenService.Login(request.UserName, request.Password)
	if err != nil {
		httpStatus = http.StatusInternalServerError
		return
	}
}

func NewAuthenHandler(authenService services.AuthenService, v *validator.Validate) *AuthenHandler {
	return &AuthenHandler{authenService: authenService, v: v}
}
