package auth

import (
	"context"
	"github.com/zbitech/common/interfaces"
)

type FakeAuthorizerFactory struct {
	authorizer interfaces.AccessAuthorizerIF
}

func NewFakeAuthorizerFactory() interfaces.AccessAuthorizerFactoryIF {
	return &FakeAuthorizerFactory{}
}

func (f *FakeAuthorizerFactory) Init(ctx context.Context) error {
	f.authorizer = NewFakeAccessAuthorizer()
	return nil
}

func (f *FakeAuthorizerFactory) GetAccessAuthorizer() interfaces.AccessAuthorizerIF {
	return f.authorizer
}
