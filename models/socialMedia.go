package models

type SocialMedia struct {
	baseModel
	Name           string
	SocialMediaURL string
	UserID         uint
}
