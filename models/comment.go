package models

type Comment struct {
	baseModel
	UserID  uint
	PhotoID uint
	Message string
}
