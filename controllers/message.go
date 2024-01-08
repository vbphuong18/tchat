package controllers

import (
	"TChat/domain"
	"TChat/dto"
	"TChat/helper"
	"TChat/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"time"
)

type MessageHandler struct {
	messageService services.MessageService
	v              validator.Validate
}

func (m *MessageHandler) CreateMessage(ctx *gin.Context) {
	var (
		response   dto.CreateMessageResponse
		request    dto.CreateMessageRequest
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
		return
	}
	err = m.v.Struct(request)
	if err != nil {
		httpStatus = http.StatusBadRequest
		return
	}
	err = m.messageService.CreateMessage(domain.ChatMessage{
		SendID:    request.SendID,
		ReceiveID: request.ReceiveID,
		Message:   request.Message,
		CreatedAt: time.Now(),
	})
	if err != nil {
		httpStatus = http.StatusInternalServerError
		return
	}
}

func (m *MessageHandler) ListMessage(ctx *gin.Context) {
	var (
		response   dto.ListMessageResponse
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
	sendID := ctx.Query("send_id")
	receiveID := ctx.Query("receive_id")
	startTime, err := helper.ConvertStringToTime(ctx.Query("start_time"))
	if err != nil {
		httpStatus = http.StatusBadRequest
		return
	}
	endTime, err := helper.ConvertStringToTime(ctx.Query("end_time"))
	if err != nil {
		httpStatus = http.StatusBadRequest
		return
	}
	msgDomain, err := m.messageService.ListMessage(sendID, receiveID, startTime, endTime)
	if err != nil {
		httpStatus = http.StatusInternalServerError
		return
	}
	var msgDto []dto.Message
	for i := 0; i < len(msgDomain); i++ {
		msgDto = append(msgDto, dto.Message{
			MessageID: msgDomain[i].MessageID,
			SendID:    msgDomain[i].SendID,
			ReceiveID: msgDomain[i].ReceiveID,
			Message:   msgDomain[i].Message,
			CreatedAt: msgDomain[i].CreatedAt,
		})
	} // convert msgDomain to msgDto
	response.Data = msgDto
	return
}

func (m *MessageHandler) DeleteMessage(ctx *gin.Context) {
	var (
		response   dto.DeleteMessageResponse
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
	messageID := ctx.Query("message_id")
	err := m.messageService.DeleteMessage(messageID)
	if err != nil {
		httpStatus = http.StatusInternalServerError
		return
	}
}

func NewMessageHandler(messageService services.MessageService) *MessageHandler {
	return &MessageHandler{messageService: messageService}
}
