package service

import (
	cmdResponse "HELLO-GO/cmd/response"
	"HELLO-GO/connector"
	"HELLO-GO/constant"
	"HELLO-GO/helper"
	"HELLO-GO/model/response"
)

func GetCommentResponse(header map[string]string, url string) cmdResponse.CommentResponse {
	commentAPiResponse := response.Comments{}
	connector.HttpGet(url, constant.CommentAPI, header, &commentAPiResponse)
	cmdCommentResponse := helper.UpdateCommentResponse(commentAPiResponse)
	return cmdCommentResponse
}
