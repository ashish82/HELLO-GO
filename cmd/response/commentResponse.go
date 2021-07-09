package response

import "HELLO-GO/model/response"

type CommentResponse struct {
	StatusCode     int
	StatusMessage  string
	CommentDetails response.Comments `json:"commentDetails"`
}
