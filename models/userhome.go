package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

func Debug(infos ...interface{}) {
	if true {
		fmt.Printf("DEBUG: "+fmt.Sprintf("%s\n", infos[0]), infos[1:]...)
	}
}

func init() {
	fmt.Println("init models file ")
	orm.RegisterModel(new(Home))
}

type Home struct {
	Id              int64
	User            *User  `orm:"rel(fk)"`
	User_email      string `orm:"size(30)"`
	User_name       string `orm:"size(128)"`
	User_email      string `orm:"size(128);unique"`
	User_address    string `orm:"size(128)"`
	User_want_to_be string `orm:"size(128)"`
	User_really_is  string `orm:"size(128)"`
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

func Home() orm.QuerySeter {
	return orm.NewOrm().QueryTable("Home").OrderBy("-Id")
}
