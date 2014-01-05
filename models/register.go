package models

import (
	//	"Eva1/utils"
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
)

type RegisterForm struct {
	Email      string `form:"email" valid:"Required;Email;MaxSize(80)"`
	UserName   string `form:"name" valid:"Required;AlphaDash;MinSize(5);MaxSize(30)"`
	Password   string `form:"password" valid:"Required;MinSize(4);MaxSize(30)"`
	PasswordRe string `form:"passwordre" valid:"Required;MinSize(4);MaxSize(30)"`
}

func (form *RegisterForm) Valid(v *validation.Validation) {
	// Check if passwords of two times are same.
	if form.Password != form.PasswordRe {
		v.SetError("PasswordRe", "auth.repassword_not_match")
		return
	}

	e1, e2, _ := CanRegistered(form.UserName, form.Email)

	if !e1 {
		v.SetError("UserName", "用户名已经存在")
	}

	if !e2 {
		v.SetError("Email", "邮箱已经存在")
	}
}
func RegisterUser(user *User, form RegisterForm) error {
	//salt := GetUserSalt()
	//pwd := utils.EncodePassword(form.Password, salt)

	//user.User_name = form.UserName
	//user.User_email = form.Email

	//// save salt and encode password, use $ as split char
	//user.User_password = fmt.Sprintf("%s$%s", salt, pwd)
	//// save md5 email value for gravatar
	//user.User_email = utils.EncodeMd5(form.Email)

	user.User_password = form.Password
	user.User_name = form.UserName
	user.User_email = form.Email
	fmt.Println(form.UserName)
	fmt.Println(form.Password)
	fmt.Println(form.Email)
	fmt.Println("RegisterUser finished ")
	return user.Insert()
}
func CanRegistered(userName string, email string) (bool, bool, error) {
	cond := orm.NewCondition()
	cond = cond.Or("UserName", userName).Or("Email", email)

	var maps []orm.Params
	o := orm.NewOrm()
	n, err := o.QueryTable("user").SetCond(cond).Values(&maps, "UserName", "Email")
	if err != nil {
		return false, false, err
	}

	e1 := true
	e2 := true

	if n > 0 {
		for _, m := range maps {
			if e1 && orm.ToStr(m["UserName"]) == userName {
				e1 = false
			}
			if e2 && orm.ToStr(m["Email"]) == email {
				e2 = false
			}
		}
	}

	return e1, e2, nil
}
