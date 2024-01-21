package controllers

import (
	"TChat/pkg/domain"
	"TChat/pkg/dto"
	"TChat/pkg/services"
	"TChat/pkg/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type UserHandler struct {
	userService services.UserService
	v           *validator.Validate
}

func (u *UserHandler) CreateUser(ctx *gin.Context) {
	var (
		response   dto.CreateUserResponse
		request    dto.CreateUserRequest
		httpStatus = http.StatusOK
	)
	defer func() {
		if httpStatus != http.StatusOK {
			response.ReturnCode = 2
			response.ReturnMessage = "Fail"
		} else {
			response.ReturnCode = 1
			response.ReturnMessage = "OK"
		}
		ctx.JSON(httpStatus, response)
	}()
	err := ctx.ShouldBind(&request)
	if err != nil {
		httpStatus = http.StatusBadRequest
		fmt.Println(err)
		return
	}
	err = u.v.Struct(request)
	if err != nil {
		httpStatus = http.StatusBadRequest
		return
	}
	passwordHash, err := utils.Hash(request.Password)
	if err != nil {
		httpStatus = http.StatusBadRequest
		return
	}
	err = u.userService.CreateUser(domain.User{
		PhoneNumber: request.PhoneNumber,
		DateOfBirth: request.DateOfBirth,
		Name:        request.Name,
		Email:       request.Email,
		Gender:      domain.GenderType(request.Gender),
		UserName:    request.UserName,
		Password:    passwordHash,
		AvtImg:      request.AvtImg,
		CoverImg:    request.CoverImg,
	})
	if err != nil {
		httpStatus = http.StatusInternalServerError
		return
	}
}

func (u *UserHandler) ListUser(ctx *gin.Context) {
	var (
		response   dto.ListUserResponse
		httpStatus = http.StatusOK
	)
	defer func() {
		if httpStatus != http.StatusOK {
			response.ReturnCode = 2
			response.ReturnMessage = "Fail"
		} else {
			response.ReturnCode = 1
			response.ReturnMessage = "OK"
		}
		ctx.JSON(httpStatus, response)
	}()
	usDomain, err := u.userService.ListUser()
	if err != nil {
		httpStatus = http.StatusInternalServerError
		return
	}
	var usDto []dto.User
	for i := 0; i < len(usDomain); i++ {
		usDto = append(usDto, dto.User{
			UserID:      usDomain[i].UserID,
			DateOfBirth: usDomain[i].DateOfBirth,
			Name:        usDomain[i].Name,
			Gender:      dto.GenderType(usDomain[i].Gender),
			AvtImg:      usDomain[i].AvtImg,
			CoverImg:    usDomain[i].CoverImg,
		})
	} // convert usDomain to usDto
	response.Data = usDto
	return
}

func (u *UserHandler) SearchUser(ctx *gin.Context) {
	var (
		response   dto.SearchUserResponse
		httpStatus = http.StatusOK
	)
	defer func() {
		if httpStatus != http.StatusOK {
			response.ReturnCode = 2
			response.ReturnMessage = "Fail"
		} else {
			response.ReturnCode = 1
			response.ReturnMessage = "OK"
		}
		ctx.JSON(httpStatus, response)
	}()
	name := ctx.Query("name")
	phoneNumber := ctx.Query("phone_number")
	fmt.Println(phoneNumber)
	if name == "" && phoneNumber == "" {
		httpStatus = http.StatusBadRequest
		return
	}
	srchDomain, err := u.userService.SearchUser(name, phoneNumber)
	if err != nil {
		httpStatus = http.StatusInternalServerError
		return
	}
	var srchDto []dto.User
	for i := 0; i < len(srchDomain); i++ {
		srchDto = append(srchDto, dto.User{
			UserID:      srchDomain[i].UserID,
			DateOfBirth: srchDomain[i].DateOfBirth,
			Name:        srchDomain[i].Name,
			Gender:      dto.GenderType(srchDomain[i].Gender),
			AvtImg:      srchDomain[i].AvtImg,
			CoverImg:    srchDomain[i].CoverImg,
		})
	} // convert srchDomain to srchDto
	response.Data = srchDto
	return
}

func (u *UserHandler) GetUserByUserID(ctx *gin.Context) {
	var (
		response   dto.GetUserByUserIDResponse
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
	userID := ctx.Query("user_id")
	userDomain, err := u.userService.GetUserByUserID(userID)
	if err != nil {
		httpStatus = http.StatusInternalServerError
		return
	}
	response.Data = dto.User{
		UserID:      userDomain.UserID,
		DateOfBirth: userDomain.DateOfBirth,
		Name:        userDomain.Name,
		Gender:      dto.GenderType(userDomain.Gender),
		AvtImg:      userDomain.AvtImg,
		CoverImg:    userDomain.CoverImg,
	}
	return
}

func (u *UserHandler) DeleteUser(ctx *gin.Context) {
	var (
		response   dto.DeleteUserResponse
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
	userID := ctx.Query("user_id")
	err := u.userService.DeleteUser(userID)
	if err != nil {
		httpStatus = http.StatusInternalServerError
		return
	}
}

func NewUserHandler(userService services.UserService, v *validator.Validate) *UserHandler {
	return &UserHandler{userService: userService, v: v}
}
