package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	FirstName     string    `json:"first_name"`
	LastName      string    `json:"last_name"`
	AccountNumber string    `json:"account_number"`
	UserEmail     string    `json:"user_email"`
	UserMobile    string    `json:"user_mobile"`
	UserPhone     string    `json:"user_phone"`
	Username      string    `json:"username"`
	RoleID        string    `json:"role_id"`
	Password      string    `json:"password"`
	AdminApproval string    `json:"admin_approval"`
	CreatedOn     time.Time `json:"created_on"`
}

type GetUserResp struct {
	Message string `json:"message"`
	User    *User  `json:"user"`
}

type CreateUserReq struct {
	User User `json:"user"`
}

type CreateUserResp struct {
	Message string `json:"message"`
	User    *User  `json:"user"`
}

type DeleteUserResp struct {
	Message string `json:"message"`
	ID      string `json:"id"`
}

type UpdateUserReq struct {
	User User `json:"user"`
}

type UpdateUserResp struct {
	Message string `json:"message"`
	User    *User  `json:"user"`
}

type GetAllUserResp struct {
	Message string  `json:"message"`
	User    []*User `json:"user"`
}

type Authentication struct {
	gorm.Model
	AuthToken    string `json:"auth_token"`
	UserId       string `json:"user_id"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	AuthState    string `json:"auth_state"`
}

type Role struct {
	gorm.Model
	RoleName        string `json:"role_name"`
	AccessPrivilage string `json:"cccess_privilage"`
	UserID          []User
}
