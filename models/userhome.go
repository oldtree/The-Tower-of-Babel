package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type Home struct {
	Id              int64
	User            *User  `orm:"rel(fk)"`
	User_name       string `orm:"size(128)"`
	User_email      string `orm:"size(128);unique"`
	User_address    string `orm:"size(128)"`
	User_want_to_be string `orm:"size(128)"`
	User_really_is  string `orm:"size(128)"`
}

func init() {
	fmt.Println("init models file ")
	orm.RegisterModel(new(Home))
}

func (u *Home) Insert() error {
	_, err := orm.NewOrm().Insert(u)
	return err
}

func (u *Home) Delete() error {
	_, err := orm.NewOrm().Delete(u)
	return err
}
func (u *Home) Read(fileds ...string) error {
	err := orm.NewOrm().Read(u, fileds...)
	return err
}

func (u *Home) Update(fileds ...string) error {
	_, err := orm.NewOrm().Update(u, fileds...)
	return err
}

func SelfHome() orm.QuerySeter {
	return orm.NewOrm().QueryTable("Home").OrderBy("-Id")
}
