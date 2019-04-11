package atypeek

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type (
	Profile struct {
		Owner        sdk.AccAddress `json:"owner"`
		Projects     []string       `json:"projects"`
		Skills       []string       `json:"skills"`
		Courses      []string       `json:"courses"`
		Endorsements []string       `json:"endorsements"`
	}

	IProfile interface {
		AddProject(info ProjectInfo) (id string, err error)
		ListProjects() []Project
	}

	Project struct {
		Id          string         `json:"id"`
		Owner       sdk.AccAddress `json:"owner"`
		CustomerId  string         `json:"customerId"`
		Title       string         `json:"title"`
		Description string         `json:"description"`
		StartDate   string         `json:"startDate"`
		EndDate     string         `json:"endDate"`
		Skills      []string       `json:"skills"`
	}

	ProjectInfo struct {
		Id          string
		Owner       sdk.AccAddress
		CustomerId  string
		Title       string
		Description string
		StartDate   string
		EndDate     string
	}

	Skill struct {
		Owner        sdk.AccAddress `json:"owner"`
		Id           string         `json:"id"`
		Name         string         `json:"name"`
		Courses      []string       `json:"courses"`
		Score        int64          `json:"score"`
		Endorsements []string       `json:"endorsements"`
	}

	ISkill interface {
		AddSkill(idSkill string) error
		RemoveSkill(idSkill string) error
	}

	Course struct {
		Owner sdk.AccAddress `json:"owner"`
		Id    string         `json:"id"`
		Name  string         `json:"name"`
	}

	ICourse interface {
		AddCourse(idSkill string, idCourse string) error
		RemoveCourse(idSkill string, idCourse string) error
	}

	Endorsement struct {
		Id          string         `json:"id"`
		IdSkill     string         `json:"idSkill"`
		Contributor sdk.AccAddress `json:"contributor"`
		Time        string         `json:"time"`
		Receiver    sdk.AccAddress `json:"receiver"`
		Vote        int            `json:"vote"`
	}
)

func NewProfile() Profile {
	return Profile{
		Owner:    nil,
		Projects: []string{},
		Skills:   []string{},
		Courses:  []string{},
	}
}

func (p Profile) AddProject(i ProjectInfo) (id string, err error) {

	p.Projects = append(p.Projects, i.Id)
	return "", nil
}

func (p Profile) ListProjects() []string {

	return p.Projects
}

func (p Profile) String() string {
	return fmt.Sprintf("%+v", p)
}

func NewProject() Project {
	return Project{
		Id:          "",
		Owner:       nil,
		CustomerId:  "",
		Title:       "",
		Description: "",
		StartDate:   "",
		EndDate:     "",
		Skills:      []string{},
	}
}

func NewProjectWithProjectInfo(i ProjectInfo) Project {
	return Project{
		Id:          i.Id,
		Owner:       i.Owner,
		CustomerId:  i.CustomerId,
		Title:       i.Title,
		Description: i.Description,
		StartDate:   i.StartDate,
		EndDate:     i.EndDate,
		Skills:      []string{},
	}
}

func (p Project) String() string {
	return fmt.Sprintf("%+v", p)
}

func NewSkill() Skill {
	return Skill{
		Owner:   nil,
		Id:      "",
		Name:    "",
		Courses: nil,
	}
}

func (s Skill) String() string {
	return fmt.Sprintf("%+v", s)
}

func NewCourse() Course {
	return Course{
		Id: "",
	}
}

func (c Course) String() string {
	return fmt.Sprintf("%+v", c)
}

func (e Endorsement) String() string {
	return fmt.Sprintf("%+v", e)
}

func NewEndorsement() Endorsement {
	return Endorsement{
		Id:          "",
		IdSkill:     "",
		Contributor: nil,
		Time:        "",
		Receiver:    nil,
		Vote:        0,
	}
}

type SkillScore struct {
	IdSkill string `json:"id"`
	Score   int    `json:"score"`
}

func (s SkillScore) String() string {
	return fmt.Sprintf("%+v", s)
}
