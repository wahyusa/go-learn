package models

// Response
//
//	{
//		"id": 0,
//		"email": "string",
//		"username": "string",
//		"age": 0,
//		"profile_image_url": "string"
//	  }
type UserRegisterResponse struct {
	ID              int    `json:"id"`
	Email           string `json:"email"`
	Username        string `json:"username"`
	Age             int    `json:"age"`
	ProfileImageURL string `json:"profile_image_url"`
}
