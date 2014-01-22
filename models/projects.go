package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"sort"
	"time"
)

func init() {
	fmt.Println("init models file ")
	orm.RegisterModel(new(Project))
}

//使用json数据格式保存项目工程文件，或者xml,在User table 里面保存对应的文件路径就可以
type Project struct {
	Id              int64
	User            *User     `orm:"rel(fk)"`
	User_email      string    `orm:"size(30)"`
	Project_created time.Time `orm:"auto_now_add"`
	Project_update  time.Time `orm:"auto_now"`
	Star            int
	User_count      int
	Path            string `orm:"size(256)"`
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
	err := orm.NewOrm().Read(p, fileds...)
	return err
}

func (p *Project) Update(fileds ...string) error {
	_, err := orm.NewOrm().Update(p, fileds...)
	return err
}

func Projects() orm.QuerySeter {
	return orm.NewOrm().QueryTable("Projects").OrderBy("-Id")
}

type lessFunc func(p1, p2 *Project) bool

type ProjectSorter struct {
	projects []Project
	less     []lessFunc
}

func (ms *ProjectSorter) Sort(Projects []Project) {
	sort.Sort(ms)

}

func OrderedBy(projects []Project, less ...lessFunc) *ProjectSorter {
	return &ProjectSorter{
		projects: projects,
		less:     less,
	}
}

func (ms *ProjectSorter) Len() int {
	return len(ms.projects)
}

func (ms *ProjectSorter) Swap(i, j int) {
	ms.projects[i], ms.projects[j] = ms.projects[j], ms.projects[i]
}

func (ms *ProjectSorter) Less(i, j int) bool {
	p, q := &ms.projects[i], &ms.projects[j]
	var k int
	for k = 0; k < len(ms.less)-1; k++ {
		less := ms.less[k]
		switch {
		case less(p, q):
			return true
		case less(q, p):
			return false
		}
	}
	return ms.less[k](p, q)
}

// 文章列表按照日期和ID排序
func SortProject(projects []Project) (projectsaftersort []Project) {
	stars := func(p1, p2 *Project) bool {
		return p1.Star > p2.Star
	}
	update := func(p1, p2 *Project) bool {
		return p1.Project_update.After(p2.Project_update)
	}
	sort.Sort(OrderedBy(projects, update, stars))
	return projects
}
