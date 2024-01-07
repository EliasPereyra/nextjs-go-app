package model

type User struct {
  Id		int	`json:"id"`
  Fullname	string	`json:"fullname"`
  Email		string	`json:"email"`
  Profile_img	string	`json:"profile_img"`
}