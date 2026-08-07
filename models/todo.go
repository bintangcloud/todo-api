package models

type Todo struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Status      string `json:"status" binding:"required"`
	UserID      uint   `json:"user_id"`
}
