package test

import (
	"context"
	"github.com/golang-jwt/jwt"
	"github.com/zbitech/common/factory"
	"github.com/zbitech/common/pkg/jwtutil"
	"github.com/zbitech/common/pkg/model/entity"
	"github.com/zbitech/common/pkg/model/object"
	"github.com/zbitech/fake/klient"
	"github.com/zbitech/fake/mgr"
	"github.com/zbitech/fake/repo"
)

func InitTest(ctx context.Context) {
	factory.InitConfig(ctx)
	factory.InitProjectResourceConfig(ctx)

	factory.SetServiceFactory(ctx, repo.NewFakeServiceFactory())
	factory.SetAccessAuthorizerFactory(ctx, repo.NewFakeAccessAuthorizerFactory())
	factory.SetKubernetesFactory(ctx, klient.NewFakeKlientFactory(), "")
	factory.SetManagerFactory(ctx, mgr.NewFakeResourceManagerFactory())
}

func ParseJwtClaims(tokenString string, user *entity.User, err error) (jwt.Claims, *entity.User, error) {
	token, _ := jwtutil.ParseWithClaims(tokenString, &object.ZBIBasicClaims{}, func(token *jwt.Token) error {
		return err
	})

	if err != nil {
		return nil, nil, err
	}
	return token.Claims.(*object.ZBIBasicClaims), user, err
}
