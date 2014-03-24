package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	//	"time"
)

func New_user(user_name string, user_email string, user_passworld string) *User {
	new_user := User{}
	new_user.User_name = user_name
	new_user.User_email = user_email
	new_user.User_password = user_passworld
	if new_user.Insert() != nil {
		fmt.Println("User insert error")
		return nil
	}
	return &new_user
}

func Login_user(user_email string, user_passworld string) *User {
	new_user := User{}
	new_user.User_email = user_email
	new_user.User_password = user_passworld
	if new_user.Insert() != nil {
		fmt.Println("User insert error")
		return nil
	}
	return &new_user

}

func Get_user_byname(username string) (users []User, err error) {
	con := orm.NewCondition()
	con.Or("UserName", username)

	var results_map []orm.ParamsList
	o := orm.NewOrm()
	n, err := o.QueryTable("User").SetCond(con).ValuesList(&results_map, "UserName")
	fmt.Println("Get_user_byname error")
	if err != nil {
		return nil, err
	}
	if n == 0 {
		return nil, nil
	}
	for u := range results_map {
		fmt.Println(u)

	}
	return users, nil
}
