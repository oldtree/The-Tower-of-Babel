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
	Id                     int64
	User_name              string    `orm:"size(30)"`
	User_email             string    `orm:"size(30);unique"`
	User_address           string    `orm:"size(128)"`
	User_password          string    `orm:"size(30);unique"`
	User_created           time.Time `orm:"auto_now_add"`
	User_update            time.Time `orm:"auto_now"`
	User_company           string    `orm:"size(128)"`
	User_want_to_be        string    `orm:"size(128)"`
	User_really_is         string    `orm:"size(128)"`
	User_project_json_path string    `orm:"size(128)"`
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
