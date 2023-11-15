package dto

type UserSearchQuery struct {
	Current  int    `query:"current" default:"1"`
	Size     int    `query:"size" default:"10"`
	Username string `query:"username"`
}
