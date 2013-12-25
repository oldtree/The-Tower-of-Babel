package models

import (
	"Eva1/utils"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/session"
	"strings"
)

func LoginUser(user *User, c *beego.Controller) {
	c.CruSession = beego.GlobalSessions.SessionRegenerateId(c.Ctx.ResponseWriter, c.Ctx.Request)
	c.Ctx.Input.CruSession = c.CruSession

	sess := c.StartSession()
	sess.Set("auth_user_id", user.Id)
}

func LogoutUser(c *beego.Controller) {
	//c.CruSession.Delete("auth_user_id")
	c.DelSession("auth_user_id")
	c.DestroySession()
}

func GetUserSalt() string {
	return utils.GetRandomString(10)
}

func GetUserFromSession(user *User, sess session.SessionStore) bool {
	if id, ok := sess.Get("auth_user_id").(int64); ok && id > 0 {
		user.Id = id
		if user.Read("Id") == nil {
			return true
		}
	}
	return false
}

// check if exist user by username or email
func HasUser(user *User, username string) bool {
	var err error
	qs := orm.NewOrm()
	if strings.IndexRune(username, '@') == -1 {
		user.User_name = username
		err = qs.Read(user, "UserName")
	} else {
		user.User_email = username
		err = qs.Read(user, "Email")
	}
	if err == nil {
		return true
	}
	return false
}

// verify username/email and password
func VerifyUser(user *User, username, password string) (success bool) {
	// search user by username or email
	if HasUser(user, username) == false {
		return
	}

	if VerifyPassword(password, user.User_password) {
		// success
		success = true
	}
	return
}

// compare raw password and encoded password
func VerifyPassword(rawPwd, encodedPwd string) bool {

	// split
	var salt, encoded string
	if len(encodedPwd) > 11 {
		salt = encodedPwd[:10]
		encoded = encodedPwd[11:]
	}

	return utils.EncodePassword(rawPwd, salt) == encoded
}
