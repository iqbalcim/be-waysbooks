package authdto

type LoginResponse struct {
	ID    int    `gorm:"type:varchar(255)" json:"id"`
	Email string `gorm:"type:varchar(255)" json:"email"`
	Name  string `gorm:"type:varchar(255)" json:"fullName"`
	Role  string `gorm:"type:varchar(255)" json:"role"`
	Token string `gorm:"type:varchar(255)" json:"token"`
}

type CheckAuthResponse struct {
	Id    int    `gorm:"type: int" json:"id"`
	Email string `gorm:"type:varchar(255)" json:"email"`
	Name  string `gorm:"type:varchar(255)" json:"fullName"`
	Role  string `gorm:"type:varchar(255)" json:"role"`
}