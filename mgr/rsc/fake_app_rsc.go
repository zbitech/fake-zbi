package rsc

import (
	"context"
	"github.com/zbitech/common/interfaces"
	"github.com/zbitech/common/pkg/model/config"
	"github.com/zbitech/common/pkg/model/entity"
	"github.com/zbitech/common/pkg/model/object"
	"github.com/zbitech/common/pkg/model/spec"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

type FakeAppResourceManager struct {
	FakeCreateProjectIngressAsset    func(ctx context.Context, obj *unstructured.Unstructured, project *entity.Project, instance entity.InstanceIF, action string) (*spec.ResourceObject, error)
	FakeCreateControllerIngressAsset func(ctx context.Context, obj *unstructured.Unstructured, project *entity.Project, action string) (*spec.ResourceObject, error)
	FakeGetAppResources              func(version string) (*config.VersionedResourceConfig, bool)
	FakeCreateSnapshotAsset          func(ctx context.Context, req *object.SnapshotRequest) ([]spec.ResourceObject, error)
	FakeCreateSnapshotScheduleAsset  func(ctx context.Context, req *object.SnapshotScheduleRequest) ([]spec.ResourceObject, error)
	FakeCreateVolumeAsset            func(ctx context.Context, volumes ...spec.VolumeSpec) ([]spec.ResourceObject, error)
}

func NewFakeIngressResourceManager() interfaces.AppResourceManagerIF {
	return &FakeAppResourceManager{}
}

func (f FakeAppResourceManager) GetAppResources(version string) (*config.VersionedResourceConfig, bool) {
	return f.FakeGetAppResources(version)
}

func (f FakeAppResourceManager) CreateProjectIngressAsset(ctx context.Context, obj *unstructured.Unstructured, project *entity.Project, instance entity.InstanceIF, action string) (*spec.ResourceObject, error) {
	return f.FakeCreateProjectIngressAsset(ctx, obj, project, instance, action)
}

func (f FakeAppResourceManager) CreateControllerIngressAsset(ctx context.Context, obj *unstructured.Unstructured, project *entity.Project, action string) (*spec.ResourceObject, error) {
	return f.FakeCreateControllerIngressAsset(ctx, obj, project, action)
}

func (f FakeAppResourceManager) CreateSnapshotAsset(ctx context.Context, req *object.SnapshotRequest) ([]spec.ResourceObject, error) {
	return f.FakeCreateSnapshotAsset(ctx, req)
}

func (f FakeAppResourceManager) CreateSnapshotScheduleAsset(ctx context.Context, req *object.SnapshotScheduleRequest) ([]spec.ResourceObject, error) {
	return f.FakeCreateSnapshotScheduleAsset(ctx, req)
}

func (f FakeAppResourceManager) CreateVolumeAsset(ctx context.Context, volumes ...spec.VolumeSpec) ([]spec.ResourceObject, error) {
	return f.FakeCreateVolumeAsset(ctx, volumes...)
}
