package models

type Photo struct {
	baseModel
	Title    string
	Caption  string
	PhotoURL string
	UserID   uint
	Comment  []Comment
}
