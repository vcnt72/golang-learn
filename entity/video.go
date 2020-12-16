package entity 

import "time"

type Person struct {
	ID uint64 `json:"id" gorm:"primary_key;auto_increment" json:"id"`
	FirstName string `json:"firstname" gorm:"type:varchar(100)"`
	LastName string `json:"lastname" gorm:"type:varchar(100)"`
	Age int8 `json:"age" binding:"gte=1,lte=130"`
	Email string `json:"email" validate:"required,email" gorm:"type:varchar(256);UNIQUE"`
}


type Video struct {
	ID uint64 `json:"id" gorm:"primary_key;auto_increment" `
	Title string `json:"title" binding:"min=2,max=10" gorm:"type:varchar(100)"`
	Description string `json:"description" binding:"max=20" gorm:"type:varchar(100)"`
	URL string `json:"url" binding:"required,url" gorm:"type:varchar(100)"`
	Author      Person    `json:"author" binding:"required" gorm:"foreignkey:PersonID"`
	PersonID    uint64    `json:"-"`
	CreatedAt   time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}