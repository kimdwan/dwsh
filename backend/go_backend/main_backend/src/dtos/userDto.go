package dtos

type UserSignUpDto struct {
	Email         string  `json:"email" validate:"required,email"`
	Password      string  `json:"password" validate:"required,min=4,max=16"`
	User_name     string  `json:"user_name" validate:"required,max=10"`
	User_type     string  `json:"user_type" validate:"required,oneof=DS GUEST"`
	Secret_key    *string `json:"secret_key,omitempty" validate:"omitempty,min=1"`
	Our_first_day *string `json:"our_first_day,omitempty" validate:"omitempty,datetime=2006-01-02"`
	Term_agree_1  bool    `json:"term_agree_1" validate:"boolean"`
	Term_agree_2  bool    `json:"term_agree_2" validate:"boolean"`
	Term_agree_3  bool    `json:"term_agree_3" validate:"boolean"`
}
