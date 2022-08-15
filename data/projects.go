package data

import (
	"github.com/zbitech/common/pkg/model/entity"
	"github.com/zbitech/common/pkg/model/object"
	"github.com/zbitech/common/pkg/model/ztypes"
	"time"
)

var (
	Project1 = entity.Project{Name: "project1", Network: ztypes.NetworkTypeTest, Description: "Project One", Version: "v1", Owner: Owner1.UserId, TeamId: Owner1Team.TeamId, Status: "Active", Timestamp: time.Now()}
	Project2 = entity.Project{Name: "project2", Network: ztypes.NetworkTypeTest, Description: "Project Two", Version: "v1", Owner: Owner1.UserId, TeamId: Owner1Team.TeamId, Status: "Active", Timestamp: time.Now()}
	Project3 = entity.Project{Name: "project3", Network: ztypes.NetworkTypeTest, Description: "Project Three", Version: "v1", Owner: Owner2.UserId, TeamId: Owner2Team.TeamId, Status: "Active", Timestamp: time.Now()}
	Project4 = entity.Project{Name: "project4", Network: ztypes.NetworkTypeTest, Description: "Project Four", Version: "v1", Owner: Owner2.UserId, TeamId: Owner2Team.TeamId, Status: "Active", Timestamp: time.Now()}

	Projects       = []entity.Project{Project1, Project2, Project3, Project4}
	Owner1Projects = []entity.Project{Project1, Project2}
	Owner2Projects = []entity.Project{Project3, Project4}
)

func AppendProjects(projects []entity.Project, _projects ...entity.Project) []entity.Project {
	return append(projects, _projects...)
}

func CreateProjects(count int, props map[string]interface{}) []entity.Project {

	var projects = make([]entity.Project, count)
	for index := range projects {
		projects[index] = *CreateProject(props)
	}

	return projects
}

func CreateProjectRequest(props map[string]interface{}) *object.ProjectRequest {
	return &object.ProjectRequest{
		Network:     getProperty(props, "network", randomValue(networkTypes)).(ztypes.NetworkType),
		Version:     getProperty(props, "version", randomValue(versions)).(string),
		Name:        getProperty(props, "name", randomString(10)).(string),
		Team:        getProperty(props, "team", randomString(50)).(string),
		Description: getProperty(props, "description", randomString(30)).(string),
	}
}

func CreateProject(props map[string]interface{}) *entity.Project {
	return &entity.Project{
		Name:        getProperty(props, "name", randomString(10)).(string),
		Description: getProperty(props, "description", randomString(20)).(string),
		Version:     getProperty(props, "version", randomValue(versions)).(string),
		Network:     getProperty(props, "network", randomValue(networkTypes)).(ztypes.NetworkType),
		Owner:       getProperty(props, "owner", randomString(10)).(string),
		Status:      getProperty(props, "status", ztypes.StatusActive).(ztypes.StatusType),
		Timestamp:   getProperty(props, "timestamp", time.Now()).(time.Time),
		TeamId:      getProperty(props, "team", randomString(50)).(string),
	}
}
