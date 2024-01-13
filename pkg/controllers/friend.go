package controllers

import (
	"TChat/pkg/domain"
	"TChat/pkg/dto"
	"TChat/pkg/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type FriendHandler struct {
	userService   services.UserService
	friendService services.FriendService
	v             *validator.Validate
}

func (f *FriendHandler) CreateFriend(ctx *gin.Context) {
	var (
		response   dto.CreateFriendResponse
		request    dto.CreateFriendRequest
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
	// parse
	err := ctx.ShouldBind(&request)
	if err != nil {
		httpStatus = http.StatusBadRequest
		return
	}
	// validate
	err = f.v.Struct(request)
	if err != nil {
		httpStatus = http.StatusBadRequest
		return
	}
	if request.UserID1 == request.UserID2 {
		httpStatus = http.StatusBadRequest
		return
	}
	err = f.friendService.CreateFriend(domain.Friend{
		UserID1: request.UserID1,
		UserID2: request.UserID2,
	})
	if err != nil {
		httpStatus = http.StatusInternalServerError
		return
	}
}

func (f *FriendHandler) ListFriend(ctx *gin.Context) {
	var (
		response   dto.ListFriendResponse
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
	userID := ctx.Query("user_id")
	friendDomain, err := f.friendService.ListFriend(userID)
	if err != nil {
		httpStatus = http.StatusInternalServerError
		return
	}
	var friendDto []dto.Friend
	for i := 0; i < len(friendDomain); i++ {
		friendDto = append(friendDto, dto.Friend{
			UserID1: friendDomain[i].UserID1,
			UserID2: friendDomain[i].UserID2,
		})
	}
	response.Data = friendDto
	return
}

func (f *FriendHandler) DeleteFriend(ctx *gin.Context) {
	var (
		response   dto.DeleteFriendResponse
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
	userID1 := ctx.Query("user_id_1")
	userID2 := ctx.Query("user_id_2")
	_, err := f.userService.GetUserByUserID(userID1)
	if err != nil {
		httpStatus = http.StatusBadRequest
		return
	}
	_, err = f.userService.GetUserByUserID(userID2)
	if err != nil {
		httpStatus = http.StatusBadRequest
		return
	}
	err = f.friendService.DeleteFriend(userID1, userID2)
	if err != nil {
		httpStatus = http.StatusInternalServerError
		return
	}
}

func NewFriendHandler(friendService services.FriendService, userService services.UserService, v *validator.Validate) *FriendHandler {
	return &FriendHandler{friendService: friendService, userService: userService, v: v}
}
