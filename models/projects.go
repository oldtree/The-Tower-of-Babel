package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

func init() {
	fmt.Println("init models file ")
	orm.RegisterModel(new(Project))
}

//ʹ��json���ݸ�ʽ������Ŀ�����ļ�������xml,��User table ���汣���Ӧ���ļ�·���Ϳ���
type Project struct {
	Id int64
	//User_id         int64     `xorm:"index"`
	User            *User     `orm:"rel(fk)"`
	User_email      string    `orm:"size(30)"`
	Project_created time.Time `orm:"auto_now_add"`
	Project_update  time.Time `orm:"auto_now"`
}

func (p *Project) Insert() error {
	_, err := orm.NewOrm().Insert(p)
	return err
}

func (p *Project) Delete() error {
	_, err := orm.NewOrm().Delete(p)
	return err
}
func (p *Project) Read(fileds ...string) error {
	_, err := orm.NewOrm().Read(p, fileds...)
	return err
}

func (p *Project) Update(fileds ...string) error {
	err := orm.NewOrm().Update(u, fileds...)
	return err
}

func Project() orm.QuerySeter {
	return orm.NewOrm().QueryTable("Projects").OrderBy("-Id")
}
