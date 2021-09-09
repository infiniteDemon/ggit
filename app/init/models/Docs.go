package models

type DocsRequestLogin struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}
