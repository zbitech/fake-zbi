package rsc

import (
	"context"
	"github.com/zbitech/common/interfaces"
	"github.com/zbitech/common/pkg/model/config"
	"github.com/zbitech/common/pkg/model/entity"
	"github.com/zbitech/common/pkg/model/object"
	"github.com/zbitech/common/pkg/model/spec"
	"github.com/zbitech/common/pkg/model/ztypes"
	"go.mongodb.org/mongo-driver/bson"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

type FakeProjectResourceManager struct {
	FakeGetProjectResources            func(version string) (*config.VersionedResourceConfig, bool)
	FakeCreateProject                  func(ctx context.Context, request *object.ProjectRequest) (*entity.Project, error)
	FakeUpdateProject                  func(ctx context.Context, project *entity.Project, request *object.ProjectRequest) error
	FakeCreateProjectAssets            func(ctx context.Context, project *entity.Project) ([]spec.ResourceObject, error)
	FakeCreateProjectIngressAsset      func(ctx context.Context, appIngress *unstructured.Unstructured, project *entity.Project, action ztypes.EventAction) ([]spec.ResourceObject, error)
	FakeGetInstanceResources           func(iType ztypes.InstanceType, version string) (*config.VersionedResourceConfig, bool)
	FakeCreateIngressAsset             func(ctx context.Context, projIngress *unstructured.Unstructured, instance entity.InstanceIF, action ztypes.EventAction) (*spec.ResourceObject, error)
	FakeCreateInstanceRequest          func(ctx context.Context, iType ztypes.InstanceType, iRequest interface{}) (object.InstanceRequestIF, error)
	FakeCreateInstance                 func(ctx context.Context, project *entity.Project, request object.InstanceRequestIF) (*entity.Instance, error)
	FakeUpdateInstance                 func(ctx context.Context, project *entity.Project, instance entity.InstanceIF, request object.InstanceRequestIF) error
	FakeCreateInstanceAssets           func(ctx context.Context, action ztypes.ZBIManagerAction, project *entity.Project, instance entity.InstanceIF) ([]spec.ResourceObject, error)
	FakeCreateInstanceBackupAssets     func(ctx context.Context, instance entity.InstanceIF, request *object.SnapshotRequest) ([]spec.ResourceObject, error)
	FakeUnmarshalBSONInstance          func(ctx context.Context, data bson.Raw) (entity.InstanceIF, error)
	FakeUnmarshalBSONDetails           func(ctx context.Context, iType ztypes.InstanceType, value bson.Raw) (entity.InstanceIF, error)
	FakeCreateDeploymentResourceAssets func(ctx context.Context, instance entity.InstanceIF) ([]spec.ResourceObject, error)
	FakeCreateStartResourceAssets      func(ctx context.Context, instance entity.InstanceIF) ([]spec.ResourceObject, error)
	FakeCreateSnapshotAssets           func(ctx context.Context, instance entity.InstanceIF, volume string) ([]spec.ResourceObject, error)
	FakeCreateSnapshotScheduleAssets   func(ctx context.Context, instance entity.InstanceIF, volume string, schedule ztypes.ZBIBackupScheduleType) ([]spec.ResourceObject, error)
	FakeCreateRotationAssets           func(ctx context.Context, instance entity.InstanceIF) ([]spec.ResourceObject, error)
}

func NewFakeProjectResourceManager() interfaces.ProjectResourceManagerIF {
	return &FakeProjectResourceManager{}
}

func (f FakeProjectResourceManager) GetProjectResources(version string) (*config.VersionedResourceConfig, bool) {
	return f.FakeGetProjectResources(version)
}

func (f FakeProjectResourceManager) CreateProject(ctx context.Context, request *object.ProjectRequest) (*entity.Project, error) {
	return f.FakeCreateProject(ctx, request)
}

func (f FakeProjectResourceManager) UpdateProject(ctx context.Context, project *entity.Project, request *object.ProjectRequest) error {
	return f.FakeUpdateProject(ctx, project, request)
}

func (f FakeProjectResourceManager) CreateProjectAssets(ctx context.Context, project *entity.Project) ([]spec.ResourceObject, error) {
	return f.FakeCreateProjectAssets(ctx, project)
}

func (f FakeProjectResourceManager) CreateProjectIngressAsset(ctx context.Context, appIngress *unstructured.Unstructured, project *entity.Project, action ztypes.EventAction) ([]spec.ResourceObject, error) {
	return f.FakeCreateProjectIngressAsset(ctx, appIngress, project, action)
}

func (f FakeProjectResourceManager) GetInstanceResources(iType ztypes.InstanceType, version string) (*config.VersionedResourceConfig, bool) {
	return f.FakeGetInstanceResources(iType, version)
}

func (f FakeProjectResourceManager) CreateInstanceRequest(ctx context.Context, iType ztypes.InstanceType, iRequest interface{}) (object.InstanceRequestIF, error) {
	return f.FakeCreateInstanceRequest(ctx, iType, iRequest)
}

func (f FakeProjectResourceManager) CreateInstance(ctx context.Context, project *entity.Project, request object.InstanceRequestIF) (entity.InstanceIF, error) {
	return f.FakeCreateInstance(ctx, project, request)
}

func (f FakeProjectResourceManager) UpdateInstance(ctx context.Context, project *entity.Project, instance entity.InstanceIF, request object.InstanceRequestIF) error {
	return f.FakeUpdateInstance(ctx, project, instance, request)
}

func (f FakeProjectResourceManager) CreateInstanceAssets(ctx context.Context, action ztypes.ZBIManagerAction, project *entity.Project, instance entity.InstanceIF) ([]spec.ResourceObject, error) {
	return f.FakeCreateInstanceAssets(ctx, action, project, instance)
}

func (f FakeProjectResourceManager) CreateIngressAsset(ctx context.Context, projIngress *unstructured.Unstructured, instance entity.InstanceIF, action ztypes.EventAction) (*spec.ResourceObject, error) {
	return f.FakeCreateIngressAsset(ctx, projIngress, instance, action)
}

func (f FakeProjectResourceManager) CreateInstanceBackupAssets(ctx context.Context, instance *entity.Instance, request *object.SnapshotRequest) ([]spec.ResourceObject, error) {
	return f.FakeCreateInstanceBackupAssets(ctx, instance, request)
}

func (f FakeProjectResourceManager) CreateDeploymentResourceAssets(ctx context.Context, instance entity.InstanceIF) ([]spec.ResourceObject, error) {
	return f.FakeCreateDeploymentResourceAssets(ctx, instance)
}

func (f FakeProjectResourceManager) CreateStartResourceAssets(ctx context.Context, instance entity.InstanceIF) ([]spec.ResourceObject, error) {
	return f.FakeCreateStartResourceAssets(ctx, instance)
}

func (f FakeProjectResourceManager) CreateSnapshotAssets(ctx context.Context, instance entity.InstanceIF, volume string) ([]spec.ResourceObject, error) {
	return f.FakeCreateSnapshotAssets(ctx, instance, volume)
}

func (f FakeProjectResourceManager) CreateSnapshotScheduleAssets(ctx context.Context, instance entity.InstanceIF, volume string, schedule ztypes.ZBIBackupScheduleType) ([]spec.ResourceObject, error) {
	return f.FakeCreateSnapshotScheduleAssets(ctx, instance, volume, schedule)
}

func (f FakeProjectResourceManager) CreateRotationAssets(ctx context.Context, instance entity.InstanceIF) ([]spec.ResourceObject, error) {
	return f.FakeCreateRotationAssets(ctx, instance)
}

func (f FakeProjectResourceManager) UnmarshalBSONInstance(ctx context.Context, data bson.Raw) (entity.InstanceIF, error) {
	return f.FakeUnmarshalBSONInstance(ctx, data)
}

func (f FakeProjectResourceManager) UnmarshalBSONDetails(ctx context.Context, iType ztypes.InstanceType, value bson.Raw) (entity.InstanceIF, error) {
	return f.FakeUnmarshalBSONDetails(ctx, iType, value)
}
