package helper

import (
	cmdCommentResponse "HELLO-GO/cmd/response"
	"HELLO-GO/model/response"
)

func UpdateCommentResponse(commentResponse response.Comments) cmdCommentResponse.CommentResponse {
	commentResponseCmdObj := cmdCommentResponse.CommentResponse{}
	commentResponseCmdObj.StatusCode = 200
	commentResponseCmdObj.StatusMessage = "Please find the comment details"
	commentResponseCmdObj.CommentDetails = commentResponse
	return commentResponseCmdObj
}
