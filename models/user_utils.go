package models

import (
	"Eva1/utils"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/session"
	"strings"
)

func User_login(user *User, c *beego.Controller) bool {
	//login_user_conn := orm.NewOrm()
	//login_user_conn.Using("eva")
	cond := orm.NewCondition()
	//cond = cond.And("User_email", user.User_email).And("User_password", user.User_password)
	cond = cond.And("User_email", user.User_email).And("User_password", user.User_password)

	var maps []orm.Params
	login_user_conn := orm.NewOrm()
	n, err := login_user_conn.QueryTable("user").SetCond(cond).Values(&maps, "User_email")
	if err != nil {
		return false
	}
	fmt.Println(n)
	//if n == 1 {
	//	for _, m := range maps {
	//		if orm.ToStr(m["User_email"]) == user.User_email {
	//			fmt.Println("66666666666666666666")
	//			if orm.ToStr(m["User_password"]) == user.User_password {
	//				fmt.Println(orm.ToStr(m["User_password"]))
	//				fmt.Println("5555555555555555555555")
	//				return true
	//			}
	//		}

	//	}
	//}
	fmt.Println("5555555555555555555555")
	if n == 1 {
		err := login_user_conn.Raw("SELECT Id,User_name,User_address,User_password,User_created,User_update,User_company,User_want_to_be,User_really_is,User_project_json_path,User_email, name FROM user WHERE id = ?", user.User_email).QueryRow(user)
		fmt.Println(err)
		if err != nil {
			fmt.Println(user.Id)
		}
		return true
	}
	//err := login_user_conn.Read(user)
	//if err == orm.ErrNoRows {
	//	return false
	//} else if err == orm.ErrMissPK {
	//	return false
	//} else {
	//	fmt.Println(user.User_email, user.User_password)
	//}
	//fmt.Println("22222222222222222222222")

	return false
}

func User_logout(user *User, c *beego.Controller) {
	c.Redirect("/home", 302)
}

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
