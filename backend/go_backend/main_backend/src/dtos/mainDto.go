package dtos

import "github.com/google/uuid"

// 기본적으로 백엔드 내에서 사용되는 dto
type Sub struct {
	Email     string `json:"email"`
	Name      string `json:"name"`
	User_type string `json:"user_type"`
}

type Payload struct {
	User_id uuid.UUID `json:"user_id"`
	Sub     Sub       `json:"sub"`
}
