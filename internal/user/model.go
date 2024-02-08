package user

import "github.com/rvraagusan/chat-app/internal/common/models/id"

type User struct {
	ID       id.ID  `json:"id"`
	Username string `json:"user_name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}
