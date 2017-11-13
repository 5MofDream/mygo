package models

import (
	"time"
)

//Users users
type Users struct {
	ID            int64     `json:"id" gorm:"column:id"`
	ParentID      int64     `json:"parent_id" gorm:"column:parent_id"`
	ForefatherIDs string    `json:"forefather_ids" gorm:"column:forefather_ids"`
	Parent        string    `json:"parent" gorm:"column:parent"`
	AccountID     int64     `json:"account_id" gorm:"column:account_id"`
	RoleID        int64     `json:"role_id" gorm:"column:role_id"`
	RoleIDs       string    `json:"role_ids" gorm:"column:role_ids"`
	Blocked       int64     `json:"blocked" gorm:"column:blocked"`
	ParentStr     string    `json:"parent_str" gorm:"column:parent_str"`
	Forefathers   string    `json:"forefathers" gorm:"column:forefathers"`
	Username      string    `json:"username" gorm:"column:username"`
	FundPassword  string    `json:"fund_password" gorm:"column:fund_password"`
	Nickname      string    `json:"nickname" gorm:"column:nickname"`
	Email         string    `json:"email" gorm:"column:email"`
	Password      string    `json:"password" gorm:"column:password"`
	IsAgent       int64     `json:"is_agent" gorm:"column:is_agent"`
	IsFromLink    int64     `json:"is_from_link" gorm:"column:is_from_link"`
	IsTester      int64     `json:"is_tester" gorm:"column:is_tester"`
	UserLevel     int64     `json:"user_level" gorm:"column:user_level"`
	PrizeGroup    string    `json:"prize_group" gorm:"column:prize_group"`
	LoginIP       string    `json:"login_ip" gorm:"column:login_ip"`
	RegisterIP    string    `json:"register_ip" gorm:"column:register_ip"`
	SigninAt      time.Time `json:"signin_at" gorm:"column:signin_at"`
	ActivatedAt   time.Time `json:"activated_at" gorm:"column:activated_at"`
	RememberToken string    `json:"remember_token" gorm:"column:remember_token"`
	RegisterAt    time.Time `json:"register_at" gorm:"column:register_at"`
	DeletedAt     time.Time `json:"deleted_at" gorm:"column:deleted_at"`
	CreatedAt     time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"column:updated_at"`
}

//TableName users
func (t *Users) TableName() string {
	return "users"
}
