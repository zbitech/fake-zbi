package repo

import (
	"context"
	"github.com/zbitech/common/interfaces"
	"github.com/zbitech/common/pkg/model/entity"
	"github.com/zbitech/common/pkg/model/ztypes"
	"time"
)

type FakeProjectService struct {
	FakeGetProjects                func(ctx context.Context) ([]entity.Project, error)
	FakeCreateProject              func(ctx context.Context, project *entity.Project) error
	FakeUpdateProject              func(ctx context.Context, project *entity.Project) error
	FakeGetProject                 func(ctx context.Context, name string) (*entity.Project, error)
	FakeGetProjectsByOwner         func(ctx context.Context, owner string) ([]entity.Project, error)
	FakeGetProjectsByTeam          func(ctx context.Context, team string) ([]entity.Project, error)
	FakeUpdateProjectStatus        func(ctx context.Context, project string, status ztypes.StatusType, timestamp time.Time) error
	FakeCreateProjectEvent         func(ctx context.Context, project, userId string, action ztypes.EventAction, evtErr error) error
	FakeGetProjectEvents           func(ctx context.Context, project string) ([]entity.ProjectEvent, error)
	FakeGetInstances               func(ctx context.Context) ([]entity.InstanceIF, error)
	FakeCreateInstance             func(ctx context.Context, instance entity.InstanceIF) error
	FakeUpdateInstance             func(ctx context.Context, instance entity.InstanceIF) error
	FakeGetInstance                func(ctx context.Context, project, name string) (entity.InstanceIF, error)
	FakeGetInstancesByProject      func(ctx context.Context, project string) ([]entity.InstanceIF, error)
	FakeGetInstancesByOwner        func(ctx context.Context, owner string) ([]entity.InstanceIF, error)
	FakeUpdateInstanceStatus       func(ctx context.Context, project, name string, status ztypes.StatusType, timestamp time.Time) error
	FakeCreateInstanceEvent        func(ctx context.Context, project, instance, userId string, action ztypes.EventAction, evtErr error) error
	FakeGetInstanceEvents          func(ctx context.Context, project, instance string) ([]entity.InstanceEvent, error)
	FakeCreateProjectResources     func(ctx context.Context, project string, resources ...entity.KubernetesResource) error
	FakeUpdateProjectResources     func(ctx context.Context, project string, resources ...entity.KubernetesResource) error
	FakeGetProjectResources        func(ctx context.Context, project string) ([]entity.KubernetesResource, error)
	FakeGetProjectResourcesByType  func(ctx context.Context, project string, rType ztypes.ResourceObjectType) ([]entity.KubernetesResource, error)
	FakeGetProjectResourceByType   func(ctx context.Context, project string, rType ztypes.ResourceObjectType, name string) (*entity.KubernetesResource, error)
	FakeCreateInstanceResources    func(ctx context.Context, project, instance string, resources ...entity.KubernetesResource) error
	FakeUpdateInstanceResources    func(ctx context.Context, project, instance string, resources ...entity.KubernetesResource) error
	FakeGetInstanceResourcesByType func(ctx context.Context, project, instance string, rType ztypes.ResourceObjectType) ([]entity.KubernetesResource, error)
	FakeGetInstanceResources       func(ctx context.Context, project, instance string) ([]entity.KubernetesResource, error)
	FakeGetInstanceResourceByType  func(ctx context.Context, project, instance string, rType ztypes.ResourceObjectType, name string) (*entity.KubernetesResource, error)
	FakeGetUserSummary             func(ctx context.Context, userId string) (*entity.ResourceSummary, error)
	FakeGetProjectStats            func(ctx context.Context, project string) *entity.ProjectSummary
	FakeGetOwnerStats              func(ctx context.Context, owner string) *entity.ProjectSummary
}

func NewFakeProjectService() interfaces.ProjectRepositoryIF {
	return &FakeProjectService{}
}

func (f FakeProjectService) GetProjects(ctx context.Context) ([]entity.Project, error) {
	return f.FakeGetProjects(ctx)
}

func (f FakeProjectService) CreateProject(ctx context.Context, project *entity.Project) error {
	return f.FakeCreateProject(ctx, project)
}

func (f FakeProjectService) UpdateProject(ctx context.Context, project *entity.Project) error {
	return f.FakeUpdateProject(ctx, project)
}

func (f FakeProjectService) GetProject(ctx context.Context, name string) (*entity.Project, error) {
	return f.FakeGetProject(ctx, name)
}

func (f FakeProjectService) GetProjectsByOwner(ctx context.Context, owner string) ([]entity.Project, error) {
	return f.FakeGetProjectsByOwner(ctx, owner)
}

func (f FakeProjectService) GetProjectsByTeam(ctx context.Context, team string) ([]entity.Project, error) {
	return f.FakeGetProjectsByTeam(ctx, team)
}

func (f FakeProjectService) UpdateProjectStatus(ctx context.Context, project string, status ztypes.StatusType, timestamp time.Time) error {
	return f.FakeUpdateProjectStatus(ctx, project, status, timestamp)
}

func (f FakeProjectService) CreateProjectEvent(ctx context.Context, project, actor string, action ztypes.EventAction, evtErr error) error {
	return f.FakeCreateProjectEvent(ctx, project, actor, action, evtErr)
}

func (f FakeProjectService) GetProjectEvents(ctx context.Context, project string) ([]entity.ProjectEvent, error) {
	return f.FakeGetProjectEvents(ctx, project)
}

func (f FakeProjectService) GetInstances(ctx context.Context) ([]entity.InstanceIF, error) {
	return f.FakeGetInstances(ctx)
}

func (f FakeProjectService) CreateInstance(ctx context.Context, instance entity.InstanceIF) error {
	return f.FakeCreateInstance(ctx, instance)
}

func (f FakeProjectService) UpdateInstance(ctx context.Context, instance entity.InstanceIF) error {
	return f.FakeUpdateInstance(ctx, instance)
}

func (f FakeProjectService) GetInstance(ctx context.Context, project, name string) (entity.InstanceIF, error) {
	return f.FakeGetInstance(ctx, project, name)
}

func (f FakeProjectService) GetInstancesByProject(ctx context.Context, project string) ([]entity.InstanceIF, error) {
	return f.FakeGetInstancesByProject(ctx, project)
}

func (f FakeProjectService) GetInstancesByOwner(ctx context.Context, owner string) ([]entity.InstanceIF, error) {
	return f.FakeGetInstancesByOwner(ctx, owner)
}

func (f FakeProjectService) UpdateInstanceStatus(ctx context.Context, project, name string, status ztypes.StatusType, timestamp time.Time) error {
	return f.FakeUpdateInstanceStatus(ctx, project, name, status, timestamp)
}

func (f FakeProjectService) CreateInstanceEvent(ctx context.Context, project, instance, actor string, action ztypes.EventAction, evtErr error) error {
	return f.FakeCreateInstanceEvent(ctx, project, instance, actor, action, evtErr)
}

func (f FakeProjectService) GetInstanceEvents(ctx context.Context, project, instance string) ([]entity.InstanceEvent, error) {
	return f.FakeGetInstanceEvents(ctx, project, instance)
}

func (f FakeProjectService) CreateProjectResources(ctx context.Context, project string, resources ...entity.KubernetesResource) error {
	return f.FakeCreateProjectResources(ctx, project, resources...)
}

func (f FakeProjectService) UpdateProjectResources(ctx context.Context, project string, resources ...entity.KubernetesResource) error {
	return f.FakeUpdateProjectResources(ctx, project, resources...)
}

func (f FakeProjectService) GetProjectResources(ctx context.Context, project string) ([]entity.KubernetesResource, error) {
	return f.FakeGetProjectResources(ctx, project)
}

func (f FakeProjectService) GetProjectResourcesByType(ctx context.Context, project string, rType ztypes.ResourceObjectType) ([]entity.KubernetesResource, error) {
	return f.FakeGetProjectResourcesByType(ctx, project, rType)
}

func (f FakeProjectService) GetProjectResourceByType(ctx context.Context, project string, rType ztypes.ResourceObjectType, name string) (*entity.KubernetesResource, error) {
	return f.FakeGetProjectResourceByType(ctx, project, rType, name)
}

func (f FakeProjectService) CreateInstanceResources(ctx context.Context, project, instance string, resources ...entity.KubernetesResource) error {
	return f.FakeCreateInstanceResources(ctx, project, instance, resources...)
}

func (f FakeProjectService) UpdateInstanceResources(ctx context.Context, project, instance string, resources ...entity.KubernetesResource) error {
	return f.FakeUpdateInstanceResources(ctx, project, instance, resources...)
}

func (f FakeProjectService) GetInstanceResources(ctx context.Context, project, instance string) ([]entity.KubernetesResource, error) {
	return f.FakeGetInstanceResources(ctx, project, instance)
}

func (f FakeProjectService) GetInstanceResourcesByType(ctx context.Context, project, instance string, rType ztypes.ResourceObjectType) ([]entity.KubernetesResource, error) {
	return f.FakeGetInstanceResourcesByType(ctx, project, instance, rType)
}

func (f FakeProjectService) GetInstanceResourceByType(ctx context.Context, project, instance string, rType ztypes.ResourceObjectType, name string) (*entity.KubernetesResource, error) {
	return f.FakeGetInstanceResourceByType(ctx, project, instance, rType, name)
}

func (f FakeProjectService) GetUserSummary(ctx context.Context, userId string) (*entity.ResourceSummary, error) {
	return f.FakeGetUserSummary(ctx, userId)
}

func (f FakeProjectService) GetProjectStats(ctx context.Context, project string) *entity.ProjectSummary {
	return f.FakeGetProjectStats(ctx, project)
}

func (f FakeProjectService) GetOwnerStats(ctx context.Context, owner string) *entity.ProjectSummary {
	return f.FakeGetOwnerStats(ctx, owner)
}
