package controllers

import (
	"TChat/pkg/dto"
	"TChat/pkg/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type GroupHandler struct {
	groupService services.GroupService
	userService  services.UserService
	v            *validator.Validate
}

func (g *GroupHandler) CreateGroup(ctx *gin.Context) {
	var (
		response   dto.CreateGroupResponse
		request    dto.CreateGroupRequest
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
	err = g.v.Struct(request)
	if err != nil {
		httpStatus = http.StatusBadRequest
		return
	}
	groupDomain, err := g.groupService.CreateGroup(request.ListUserID, request.Name, request.Avt)
	if err != nil {
		httpStatus = http.StatusInternalServerError
		return
	}
	response.Group = dto.Group{
		GroupID:    groupDomain.GroupID,
		Name:       groupDomain.Name,
		Avt:        groupDomain.Avt,
		ListUserID: groupDomain.ListUserID,
	}
	return
}

func (g *GroupHandler) ListGroupByUserID(ctx *gin.Context) {
	var (
		response   dto.ListGroupByUserIDResponse
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
	// get from request param
	userID := ctx.Query("user_id")
	groupDomain, err := g.groupService.ListGroupByUserID(userID)
	if err != nil {
		httpStatus = http.StatusInternalServerError
		return
	}
	var groupDto []dto.Group
	for i := 0; i < len(groupDomain); i++ {
		groupDto = append(groupDto, dto.Group{
			GroupID: groupDomain[i].GroupID,
			Name:    groupDomain[i].Name,
			Avt:     groupDomain[i].Avt,
		})
	}
	response.Group = groupDto
	return
}

func (g *GroupHandler) GetGroupByGroupID(ctx *gin.Context) {
	var (
		response   dto.GetGroupByGroupIDResponse
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
	groupID := ctx.Query("group_id")
	groupDomain, err := g.groupService.GetGroupByGroupID(groupID)
	if err != nil {
		httpStatus = http.StatusInternalServerError
		return
	}
	response.Group = dto.Group{
		GroupID:    groupDomain.GroupID,
		Name:       groupDomain.Name,
		Avt:        groupDomain.Avt,
		ListUserID: groupDomain.ListUserID,
	}
	return
}

func (g *GroupHandler) AddMember(ctx *gin.Context) {
	var (
		response   dto.AddMemberResponse
		request    dto.AddMemberRequest
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
		fmt.Println(err)
		return
	}
	err = g.v.Struct(request)
	if err != nil {
		httpStatus = http.StatusBadRequest
		return
	}
	err = g.groupService.AddMember(request.GroupID, request.ListUserID)
	if err != nil {
		httpStatus = http.StatusInternalServerError
		return
	}
}

func NewGroupHandler(groupService services.GroupService, userService services.UserService, v *validator.Validate) *GroupHandler {
	return &GroupHandler{groupService: groupService, userService: userService, v: v}
}
