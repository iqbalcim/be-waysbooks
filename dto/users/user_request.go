package usersdto

type UpdateUserRequest struct {
	Name     string `json:"fullName" form:"fullName"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Phone    string `json:"phone" form:"phone"`
	Gender   string `json:"gender" form:"gender"`
	Image    string `json:"image" from:"image"`
}