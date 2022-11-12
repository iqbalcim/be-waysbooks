package models

type Book struct {
	ID              int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Title           string `json:"title" form:"title" gorm:"type: varchar(255)"`
	PublicationDate string `json:"publication_date" form:"publication_date" gorm:"type: varchar(255)"`
	Pages           int    `json:"pages" form:"pages" gorm:"type: int"`
	ISBN            string `json:"isbn" form:"isbn" gorm:"type: varchar(255)"`
	Author          string `json:"author" form:"author" gorm:"type: varchar(255)"`
	Price           int    `json:"price" form:"price" gorm:"type: int"`
	Description     string `json:"description" form:"description" gorm:"type: text"`
	BookAttachment  string `json:"book_attachment" form:"book_attachment" gorm:"type: varchar(255)"`
	Thumbnail       string `json:"thumbnail" form:"thumbnail" gorm:"type: varchar(255)"`
}

type BookResponse struct {
	ID              int    `json:"id"`
	Title           string `json:"title"`
	PublicationDate string `json:"publicationDate"`
	Pages           int    `json:"pages"`
	ISBN            int    `json:"ISBN"`
	Author          string `json:"author"`
	Price           int    `json:"price"`
	Description     string `json:"description" gorm:"type: text"`
	BookAttachment  string `json:"book_attachment"`
	Thumbnail       string `json:"thumbnail"`
}

func (BookResponse) TableName() string {
	return "books"
}