package usersdto

type UserResponse struct {
	ID      int    `json:"id" gorm:"primaryKey"`
	Name    string `json:"fullName" form:"fullName"`
	Email   string `json:"email" form:"email"`
	Phone   string `json:"phone" form:"phone"`
	Gender  string `json:"gender" form:"gender"`
	Image   string `json:"image" from:"image"`
	Address string `json:"address" form:"address"`
}