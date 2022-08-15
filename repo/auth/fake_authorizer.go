package auth

import (
	"context"
	"github.com/zbitech/common/interfaces"
	"github.com/zbitech/common/pkg/model/ztypes"
)

type FakeAccessAuthorizer struct {
	FakeValidateProjectAction              func(ctx context.Context, project string, action ztypes.PlatformAction) error
	FakeValidateInstanceAction             func(ctx context.Context, project, instance string, action ztypes.PlatformAction) error
	FakeValidateTeamAction                 func(ctx context.Context, team string, action ztypes.PlatformAction) error
	FakeValidateUserInstanceMethodAccess   func(ctx context.Context, project, instance, method string) (ztypes.SubscriptionLevel, error)
	FakeValidateAPIKeyInstanceMethodAccess func(ctx context.Context, project, instance, method, apikey string) (ztypes.SubscriptionLevel, error)
}

func NewFakeAccessAuthorizer() interfaces.AccessAuthorizerIF {
	return &FakeAccessAuthorizer{}
}

func (f FakeAccessAuthorizer) ValidateProjectAction(ctx context.Context, project string, action ztypes.PlatformAction) error {
	return f.FakeValidateProjectAction(ctx, project, action)
}

func (f FakeAccessAuthorizer) ValidateInstanceAction(ctx context.Context, project, instance string, action ztypes.PlatformAction) error {
	return f.FakeValidateInstanceAction(ctx, project, instance, action)
}

func (f FakeAccessAuthorizer) ValidateTeamAction(ctx context.Context, team string, action ztypes.PlatformAction) error {
	return f.FakeValidateTeamAction(ctx, team, action)
}

func (f FakeAccessAuthorizer) ValidateUserInstanceMethodAccess(ctx context.Context, project, instance, method string) (ztypes.SubscriptionLevel, error) {
	return f.FakeValidateUserInstanceMethodAccess(ctx, project, instance, method)
}

func (f FakeAccessAuthorizer) ValidateAPIKeyInstanceMethodAccess(ctx context.Context, project, instance, method, apikey string) (ztypes.SubscriptionLevel, error) {
	return f.FakeValidateAPIKeyInstanceMethodAccess(ctx, project, instance, method, apikey)
}
