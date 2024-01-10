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
	err = f.friendService.CreateFriend(domain.Friend{
		UserID1: request.UserID1,
		UserID2: request.UserID2,
	})
	if err != nil {
		httpStatus = http.StatusInternalServerError
		return
	}
}

func NewFriendHandler(friendService services.FriendService, v *validator.Validate) *FriendHandler {
	return &FriendHandler{friendService: friendService, v: v}
}
