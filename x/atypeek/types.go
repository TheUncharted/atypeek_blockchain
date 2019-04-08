package atypeek

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type (
	Project struct {
		Id          string `json:"id"`
		CustomerId  string `json:"customerId"`
		Title       string `json:"title"`
		Description string `json:"description"`
		StartDate   string `json:"startDate"`
		EndDate     string `json:"endDate"`
	}

	Skill struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	}

	Course struct {
		Id string `json:"id"`
	}

	ProjectInfo struct {
		Id          string
		CustomerId  string
		Title       string
		Description string
		StartDate   string
		EndDate     string
	}
)

type IResume interface {
	AddProject(info ProjectInfo) (id string, err error)
	GetProject(id string) Project
	DeleteProject(id string) error
	UpdateProject(id string, info ProjectInfo) error
	ListProjects() []Project
}

type ISkill interface {
	AddSkill(idProject string, idSkill string) error
	RemoveSkill(idProject string, idSkill string) error
}

type ICourse interface {
	AddCourse(idSkill string, idCourse string) error
	RemoveCourse(idSkill string, idCourse string) error
}

func NewProject() Project {
	return Project{
		Id:          "",
		CustomerId:  "",
		Title:       "",
		Description: "",
		StartDate:   "",
		EndDate:     "",
	}
}

type Resume struct {
	Owner    sdk.AccAddress     `json:"owner"`
	Projects map[string]Project `json:"projects"`
	Skills   map[string]Skill   `json:"skills"`
	Courses  map[string]Course  `json:"courses"`
}

func NewResume() Resume {
	return Resume{
		Owner:    nil,
		Projects: make(map[string]Project),
		Skills:   make(map[string]Skill),
		Courses:  make(map[string]Course),
	}
}

func (r Resume) AddProject(i ProjectInfo) (id string, err error) {
	p := NewProject()
	p.Id = i.Id
	p.Title = i.Title
	p.Description = i.Description
	p.StartDate = i.StartDate
	p.EndDate = i.EndDate
	r.Projects[p.Id] = p
	return "", nil
}

func (r Resume) GetProject(id string) Project {
	return r.Projects[id]
}

func (r Resume) DeleteProject(id string) error {
	delete(r.Projects, id)
	return nil
}

func (r Resume) UpdateProject(id string, i ProjectInfo) error {
	p := NewProject()
	p.Id = i.Id
	p.Title = i.Title
	p.Description = i.Description
	p.StartDate = i.StartDate
	p.EndDate = i.EndDate
	r.Projects[id] = p
	return nil
}

func (r Resume) ListProjects() []Project {
	projects := []Project{}
	for _, v := range r.Projects {
		projects = append(projects, v)
	}
	return projects
}

func (r Resume) String() string {
	return fmt.Sprintf("%+v", r)
}
