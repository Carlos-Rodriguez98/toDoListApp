package models

type Usuario struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	UserName string `json:"userName"`
	Password string `json:"-"` //This information isn't returned.
	Image    string `json:"image"`
}
