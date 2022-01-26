package model

import (
	"fmt"
	"golang.org/x/text/message"
	"github.com/toolkits/pkg/str"
	"strings"
	"time"
)


var p *message.Printer


const (
	LOGIN_TYPE_SMS      = "sms"
	LOGIN_TYPE_EMAIL    = "email"
	LOGIN_TYPE_PWD      = "password"
	LOGIN_TYPE_LDAP     = "ldap"
	LOGIN_EXPIRES_IN = 300
)
const (
	USER_STATE_ACTIVE = iota
	USER_STATE_INACTIVE
	USER_STATE_LOCKED
	USER_STATE_FROZEN
	USER_STATE_WRITEN_OFF
)
const (
	USER_T_NATIVE = iota
	USER_T_TEMP
)

type User struct {
	Id           int64     `json:"id"`
	UUID         string    `json:"uuid" xorm:"'uuid'"`
	Username     string    `json:"username"`
	Password     string    `json:"-"`
	Passwords    string    `json:"-"`
	Dispname     string    `json:"dispname"`
	Phone        string    `json:"phone"`
	Email        string    `json:"email"`
	Im           string    `json:"im"`
	Portrait     string    `json:"portrait"`
	Intro        string    `json:"intro"`
	Organization string    `json:"organization"`
	Type         int       `json:"type" xorm:"'typ'" description:"0: long-term account; 1: temporary account"`
	Status       int       `json:"status" description:"0: active, 1: inactive, 2: locked, 3: frozen, 4: writen-off"`
	IsRoot       int       `json:"is_root"`
	LeaderId     int64     `json:"leader_id"`
	LeaderName   string    `json:"leader_name"`
	LoginErrNum  int       `json:"login_err_num"`
	ActiveBegin  int64     `json:"active_begin" description:"for temporary account"`
	ActiveEnd    int64     `json:"active_end" description:"for temporary account"`
	LockedAt     int64     `json:"locked_at" description:"locked time"`
	UpdatedAt    int64     `json:"updated_at" description:"user info change time"`
	PwdUpdatedAt int64     `json:"pwd_updated_at" description:"password change time"`
	PwdExpiresAt int64     `xorm:"-" json:"pwd_expires_at" description:"password expires time"`
	LoggedAt     int64     `json:"logged_at" description:"last logged time"`
	CreateAt     time.Time `json:"create_at" xorm:"<-"`
}



func _e(format string, a ...interface{}) error{
	return fmt.Errorf(format,a...)
}

func _s(format string, a ...interface{})string{
	return p.Sprintf(format, a...)
}


func (u *User) Validate() error {
	u.Username = strings.TrimSpace(u.Username)
	if u.Username == "" {
		return _e("username is blank")
	}



	if str.Dangerous(u.Username) {
		return _e("%s %s format error", _s("username"), u.Username)
	}

	if str.Dangerous(u.Dispname) {
		return _e("%s %s format error", _s("dispname"), u.Dispname)
	}

	if u.Phone != "" && !str.IsPhone(u.Phone) {
		return _e("%s %s format error", _s("phone"), u.Phone)
	}

	if u.Email != "" && !str.IsMail(u.Email) {
		return _e("%s %s format error", _s("email"), u.Email)
	}

	if len(u.Username) > 32 {
		return _e("username too long (max:%d)", 32)
	}

	if len(u.Dispname) > 32 {
		return _e("dispname too long (max:%d)", 32)
	}

	if strings.ContainsAny(u.Im, "%'") {
		return _e("%s %s format error", "im", u.Im)
	}

	cnt, _ := DB["rdb"].Where("((email <> '' and email=?) or (phone <> '' and phone=?)) and username=?",
		u.Email, u.Phone, u.Username).Count(u)
	if cnt > 0 {
		return _e("email %s or phone %s is exists", u.Email, u.Phone)
	}
	return nil
}
