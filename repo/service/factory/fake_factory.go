package factory

import (
	"context"
	"github.com/zbitech/common/interfaces"
	"github.com/zbitech/fake/repo/service/iam"
	"github.com/zbitech/fake/repo/service/jwtsvr"
	"github.com/zbitech/fake/repo/service/repo"
)

type FakeServiceFactory struct {
	jwtServer   interfaces.JwtServerIF
	iamService  interfaces.IAMServiceIF
	projService interfaces.ProjectRepositoryIF
}

func NewFakeServiceFactory() interfaces.ServiceFactoryIF {
	return &FakeServiceFactory{}
}

func (f *FakeServiceFactory) Init(ctx context.Context) error {
	f.jwtServer = jwtsvr.NewZBIJwtServer()
	f.iamService = iam.NewFakeIAMService()
	f.projService = repo.NewFakeProjectService()
	return nil
}

func (f *FakeServiceFactory) GetJwtServer() interfaces.JwtServerIF {
	return f.jwtServer
}

func (f *FakeServiceFactory) GetIAMService() interfaces.IAMServiceIF {
	return f.iamService
}

func (f *FakeServiceFactory) GetProjectService() interfaces.ProjectRepositoryIF {
	return f.projService
}

func (f *FakeServiceFactory) OpenConnection(ctx context.Context) error {
	return nil
}

func (f *FakeServiceFactory) CloseConnection(ctx context.Context) error {
	return nil
}
