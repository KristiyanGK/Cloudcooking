package comments

import (
	"time"
)

type CommentResponseVm struct {
	ID string `json:"id"`
	Content string `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	User string `json:"user"`
}