package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

func init() {
	fmt.Println("init models file ")
	orm.RegisterModel(new(User))
}

type User struct {
	Id              int64
	User_name       string
	User_email      string
	User_address    string
	User_password   string
	User_created    time.Time
	User_update     time.Time
	User_company    string
	User_want_to_be string
	User_really_is  string
}

func (u *User) Insert() error {
	_, err := orm.NewOrm().Insert(u)
	return err
}

func (u *User) Delete() error {
	_, err := orm.NewOrm().Delete(u)
	return err
}
func (u *User) Read(fileds ...string) error {
	_, err := orm.NewOrm().Read(u, fileds...)
	return err
}

func (u *User) Update(fileds ...string) error {
	err := orm.NewOrm().Update(u, fileds...)
	return err
}

func Users() orm.QuerySeter {
	return orm.NewOrm().QueryTable("user").OrderBy("-Id")
}
