package models

import (
	//	"fmt"
	//	"github.com/astaxie/beego/orm"
	//	"time"
	"sort"
)

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
