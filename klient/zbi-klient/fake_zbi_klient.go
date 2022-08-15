package zklient

import (
	"context"
	"github.com/zbitech/common/interfaces"
	"github.com/zbitech/common/pkg/model/entity"
	"github.com/zbitech/common/pkg/model/k8s"
	"github.com/zbitech/common/pkg/model/object"
)

type FakeZBIClient struct {
	client interfaces.KlientIF

	FakeRunInformer                func(ctx context.Context)
	FakeCreateIngress              func(ctx context.Context, project *entity.Project, instance *entity.Instance) ([]entity.KubernetesResource, error)
	FakeUpdateControllerIngress    func(ctx context.Context, project *entity.Project, action string) ([]entity.KubernetesResource, error)
	FakeUpdateProjectIngress       func(ctx context.Context, project *entity.Project, instance entity.InstanceIF, action string) ([]entity.KubernetesResource, error)
	FakeValidateProjectRequest     func(ctx context.Context, request *object.ProjectRequest) (bool, map[string]string)
	FakeCreateProject              func(ctx context.Context, request *object.ProjectRequest) (*entity.Project, error)
	FakeCreateProjectResources     func(ctx context.Context, project *entity.Project) ([]entity.KubernetesResource, error)
	FakeUpdateProject              func(ctx context.Context, project *entity.Project, request *object.ProjectRequest) error
	FakeDeleteProject              func(ctx context.Context, name string) error
	FakeGetAllProjects             func(ctx context.Context) ([]entity.Project, error)
	FakeGetProject                 func(ctx context.Context, name string) (*entity.Project, error)
	FakeGetProjectsByOwner         func(ctx context.Context, owner string) ([]entity.Project, error)
	FakeGetProjectsByTeam          func(ctx context.Context, team string) ([]entity.Project, error)
	FakeGetProjectResources        func(ctx context.Context, project *entity.Project) ([]entity.KubernetesResource, error)
	FakeValidateInstanceRequest    func(ctx context.Context, project *entity.Project, request object.InstanceRequestIF) (bool, map[string]string)
	FakeCreateInstance             func(ctx context.Context, project *entity.Project, request object.InstanceRequestIF) (entity.InstanceIF, error)
	FakeUpdateInstance             func(ctx context.Context, project *entity.Project, instance entity.InstanceIF, request object.InstanceRequestIF) error
	FakeDeleteInstance             func(ctx context.Context, project *entity.Project, instance entity.InstanceIF) error
	FakeCreateInstanceResources    func(ctx context.Context, project *entity.Project, instance entity.InstanceIF) ([]entity.KubernetesResource, error)
	FakeStopInstance               func(ctx context.Context, project *entity.Project, instance entity.InstanceIF) error
	FakeStartInstance              func(ctx context.Context, project *entity.Project, instance entity.InstanceIF) ([]entity.KubernetesResource, error)
	FakeBackupInstance             func(ctx context.Context, instance entity.InstanceIF, request *object.SnapshotRequest) ([]entity.KubernetesResource, error)
	FakeRotateInstanceCredentials  func(ctx context.Context, project *entity.Project, instance entity.InstanceIF) ([]entity.KubernetesResource, error)
	FakeDeleteInstanceVolumes      func(ctx context.Context, project *entity.Project, instance entity.InstanceIF) error
	FakeGetAllInstances            func(ctx context.Context) ([]entity.InstanceIF, error)
	FakeGetInstancesByProject      func(ctx context.Context, project string) ([]entity.InstanceIF, error)
	FakeGetInstancesByOwner        func(ctx context.Context, owner string) ([]entity.InstanceIF, error)
	FakeGetInstanceByName          func(ctx context.Context, project, instance string) (entity.InstanceIF, error)
	FakeGetInstanceResources       func(ctx context.Context, project *entity.Project, instance entity.InstanceIF) ([]entity.KubernetesResource, error)
	FakeGetStorageClasses          func(ctx context.Context) []k8s.StorageClass
	FakeGetSnapshotClasses         func(ctx context.Context) []k8s.SnapshotClass
	FakeCreateInstanceBackup       func(ctx context.Context, instance entity.InstanceIF, req *object.SnapshotRequest) (*entity.KubernetesResource, error)
	FakeScheduleInstanceBackup     func(ctx context.Context, instance entity.InstanceIF, req *object.SnapshotScheduleRequest) (*entity.KubernetesResource, error)
	FakeGetInstanceVolumes         func(ctx context.Context, instance entity.InstanceIF) []k8s.InstanceVolume
	FakeGetInstanceSnapshots       func(ctx context.Context, instance entity.InstanceIF) []k8s.Snapshot
	FakeGetInstanceSchedules       func(ctx context.Context, instance entity.InstanceIF) []k8s.SnapshotSchedule
	FakeGetResourceSummary         func(ctx context.Context) map[string]interface{}
	FakeGetProjectResourceSummary  func(ctx context.Context, project *entity.Project, extraLabels map[string]string) map[string]interface{}
	FakeGetInstanceResourceSummary func(ctx context.Context, instance entity.InstanceIF, extraLabels map[string]string) map[string]interface{}
	FakeDeleteInstanceVolume       func(ctx context.Context, instance entity.InstanceIF, name string) error
	FakeDeleteInstanceSnapshot     func(ctx context.Context, instance entity.InstanceIF, name string) error
	FakeDeleteInstanceSchedule     func(ctx context.Context, instance entity.InstanceIF, name string) error
}

func NewFakeZBIClient(client interfaces.KlientIF) (interfaces.ZBIClientIF, error) {

	return &FakeZBIClient{
		client: client,
	}, nil
}

func (f FakeZBIClient) RunInformer(ctx context.Context) {
	//	f.FakeRunInformer(ctx)
}

func (f FakeZBIClient) CreateIngress(ctx context.Context, project *entity.Project, instance entity.InstanceIF) ([]entity.KubernetesResource, error) {
	return f.FakeCreateInstanceResources(ctx, project, instance)
}

func (f FakeZBIClient) UpdateControllerIngress(ctx context.Context, project *entity.Project, action string) ([]entity.KubernetesResource, error) {
	return f.FakeUpdateControllerIngress(ctx, project, action)
}

func (f FakeZBIClient) UpdateProjectIngress(ctx context.Context, project *entity.Project, instance entity.InstanceIF, action string) ([]entity.KubernetesResource, error) {
	return f.FakeUpdateProjectIngress(ctx, project, instance, action)
}

func (f FakeZBIClient) ValidateProjectRequest(ctx context.Context, request *object.ProjectRequest) (bool, map[string]string) {
	return f.FakeValidateProjectRequest(ctx, request)
}

func (f FakeZBIClient) CreateProject(ctx context.Context, request *object.ProjectRequest) (*entity.Project, error) {
	return f.FakeCreateProject(ctx, request)
}

func (f FakeZBIClient) CreateProjectResources(ctx context.Context, project *entity.Project) ([]entity.KubernetesResource, error) {
	return f.FakeCreateProjectResources(ctx, project)
}

func (f FakeZBIClient) UpdateProject(ctx context.Context, project *entity.Project, request *object.ProjectRequest) error {
	return f.FakeUpdateProject(ctx, project, request)
}

func (f FakeZBIClient) DeleteProject(ctx context.Context, name string) error {
	return f.FakeDeleteProject(ctx, name)
}

func (f FakeZBIClient) GetAllProjects(ctx context.Context) ([]entity.Project, error) {
	return f.FakeGetAllProjects(ctx)
}

func (f FakeZBIClient) GetProject(ctx context.Context, name string) (*entity.Project, error) {
	return f.FakeGetProject(ctx, name)
}

func (f FakeZBIClient) GetProjectsByOwner(ctx context.Context, owner string) ([]entity.Project, error) {
	return f.FakeGetProjectsByOwner(ctx, owner)
}

func (f FakeZBIClient) GetProjectsByTeam(ctx context.Context, team string) ([]entity.Project, error) {
	return f.FakeGetProjectsByTeam(ctx, team)
}

func (f FakeZBIClient) GetProjectResources(ctx context.Context, project *entity.Project) ([]entity.KubernetesResource, error) {
	return f.FakeGetProjectResources(ctx, project)
}

func (f FakeZBIClient) ValidateInstanceRequest(ctx context.Context, project *entity.Project, request object.InstanceRequestIF) (bool, map[string]string) {
	return f.FakeValidateInstanceRequest(ctx, project, request)
}

func (f FakeZBIClient) CreateInstance(ctx context.Context, project *entity.Project, request object.InstanceRequestIF) (entity.InstanceIF, error) {
	return f.FakeCreateInstance(ctx, project, request)
}

func (f FakeZBIClient) UpdateInstance(ctx context.Context, project *entity.Project, instance entity.InstanceIF, request object.InstanceRequestIF) error {
	return f.FakeUpdateInstance(ctx, project, instance, request)
}

func (f FakeZBIClient) DeleteInstance(ctx context.Context, project *entity.Project, instance entity.InstanceIF) error {
	return f.FakeDeleteInstance(ctx, project, instance)
}

func (f FakeZBIClient) CreateInstanceResources(ctx context.Context, project *entity.Project, instance entity.InstanceIF) ([]entity.KubernetesResource, error) {
	return f.FakeCreateInstanceResources(ctx, project, instance)
}

func (f FakeZBIClient) StopInstance(ctx context.Context, project *entity.Project, instance entity.InstanceIF) error {
	return f.FakeStopInstance(ctx, project, instance)
}

func (f FakeZBIClient) StartInstance(ctx context.Context, project *entity.Project, instance entity.InstanceIF) ([]entity.KubernetesResource, error) {
	return f.FakeStartInstance(ctx, project, instance)
}

func (f FakeZBIClient) BackupInstance(ctx context.Context, instance *entity.Instance, request *object.SnapshotRequest) ([]entity.KubernetesResource, error) {
	return f.FakeBackupInstance(ctx, instance, request)
}

func (f FakeZBIClient) RotateInstanceCredentials(ctx context.Context, project *entity.Project, instance entity.InstanceIF) ([]entity.KubernetesResource, error) {
	return f.FakeRotateInstanceCredentials(ctx, project, instance)
}

func (f FakeZBIClient) DeleteInstanceVolumes(ctx context.Context, project *entity.Project, instance entity.InstanceIF) error {
	return f.FakeDeleteInstanceVolumes(ctx, project, instance)
}

func (f FakeZBIClient) DeleteInstanceVolume(ctx context.Context, instance entity.InstanceIF, name string) error {
	return f.FakeDeleteInstanceVolume(ctx, instance, name)
}

func (f FakeZBIClient) DeleteInstanceSnapshot(ctx context.Context, instance entity.InstanceIF, name string) error {
	return f.FakeDeleteInstanceSnapshot(ctx, instance, name)
}

func (f FakeZBIClient) DeleteInstanceSchedule(ctx context.Context, instance entity.InstanceIF, name string) error {
	return f.FakeDeleteInstanceSchedule(ctx, instance, name)
}

func (f FakeZBIClient) GetAllInstances(ctx context.Context) ([]entity.InstanceIF, error) {
	return f.FakeGetAllInstances(ctx)
}

func (f FakeZBIClient) GetInstancesByProject(ctx context.Context, project string) ([]entity.InstanceIF, error) {
	return f.FakeGetInstancesByProject(ctx, project)
}

func (f FakeZBIClient) GetInstancesByOwner(ctx context.Context, owner string) ([]entity.InstanceIF, error) {
	return f.FakeGetInstancesByOwner(ctx, owner)
}

func (f FakeZBIClient) GetInstanceByName(ctx context.Context, project, instance string) (entity.InstanceIF, error) {
	return f.FakeGetInstanceByName(ctx, project, instance)
}

func (f FakeZBIClient) GetInstanceResources(ctx context.Context, project *entity.Project, instance entity.InstanceIF) ([]entity.KubernetesResource, error) {
	return f.FakeGetInstanceResources(ctx, project, instance)
}

func (f FakeZBIClient) GetStorageClasses(ctx context.Context) []k8s.StorageClass {
	return f.FakeGetStorageClasses(ctx)
}

func (f FakeZBIClient) GetSnapshotClasses(ctx context.Context) []k8s.SnapshotClass {
	return f.FakeGetSnapshotClasses(ctx)
}

func (f FakeZBIClient) CreateInstanceBackup(ctx context.Context, instance entity.InstanceIF, req *object.SnapshotRequest) (*entity.KubernetesResource, error) {
	return f.FakeCreateInstanceBackup(ctx, instance, req)
}

func (f FakeZBIClient) ScheduleInstanceBackup(ctx context.Context, instance entity.InstanceIF, req *object.SnapshotScheduleRequest) (*entity.KubernetesResource, error) {
	return f.FakeScheduleInstanceBackup(ctx, instance, req)
}

func (f FakeZBIClient) GetInstanceVolumes(ctx context.Context, instance entity.InstanceIF) []k8s.InstanceVolume {
	return f.FakeGetInstanceVolumes(ctx, instance)
}

func (f FakeZBIClient) GetInstanceSnapshots(ctx context.Context, instance entity.InstanceIF) []k8s.Snapshot {
	return f.FakeGetInstanceSnapshots(ctx, instance)
}

func (f FakeZBIClient) GetInstanceSchedules(ctx context.Context, instance entity.InstanceIF) []k8s.SnapshotSchedule {
	return f.FakeGetInstanceSchedules(ctx, instance)
}

func (f FakeZBIClient) GetResourceSummary(ctx context.Context) map[string]interface{} {
	return f.FakeGetResourceSummary(ctx)
}

func (f FakeZBIClient) GetProjectResourceSummary(ctx context.Context, project *entity.Project, extraLabels map[string]string) map[string]interface{} {
	return f.FakeGetProjectResourceSummary(ctx, project, extraLabels)
}

func (f FakeZBIClient) GetInstanceResourceSummary(ctx context.Context, instance entity.InstanceIF, extraLabels map[string]string) map[string]interface{} {
	return f.FakeGetInstanceResourceSummary(ctx, instance, extraLabels)
}
