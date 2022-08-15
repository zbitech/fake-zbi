package data

import (
	"github.com/zbitech/common/pkg/id"
	"github.com/zbitech/common/pkg/jwtutil"
	"github.com/zbitech/common/pkg/model/entity"
	"github.com/zbitech/common/pkg/model/ztypes"
	"github.com/zbitech/common/pkg/utils"
	"github.com/zbitech/common/pkg/vars"
)

var (
	AdminUser = entity.NewUser(vars.ADMIN_USER, "ZBI", "Administrator", vars.ADMIN_EMAIL, ztypes.RoleAdmin, ztypes.SubscriptionNone)
	Owner1    = entity.NewUser("owner1", "ZBI", "owner1", "owner1@zbitech.local", ztypes.RoleOwner, ztypes.SubscriptionBronzeLevel)
	Owner2    = entity.NewUser("owner2", "ZBI", "owner2", "owner2@zbitech.local", ztypes.RoleOwner, ztypes.SubscriptionBronzeLevel)

	Owner1TeamAdmin = entity.NewUser("owner1admin", "Owner1", "Admin", "owner1admin@zbitech.local", ztypes.RoleUser, ztypes.SubscriptionTeamMember)
	Owner1TeamUser  = entity.NewUser("owner1user", "Owner1", "user", "owner1user@zbitech.local", ztypes.RoleUser, ztypes.SubscriptionTeamMember)
	Owner2TeamAdmin = entity.NewUser("owner2admin", "Owner2", "Admin", "owner2admin@zbitech.local", ztypes.RoleUser, ztypes.SubscriptionTeamMember)
	Owner2TeamUser  = entity.NewUser("owner2user", "Owner2", "User", "owner2user@zbitech.local", ztypes.RoleUser, ztypes.SubscriptionTeamMember)

	AllUsers = []entity.User{*AdminUser, *Owner1, *Owner2, *Owner1TeamAdmin, *Owner1TeamUser, *Owner2TeamAdmin, *Owner2TeamUser}

	Owner1Password          = "owner1password"
	Owner2Password          = "owner2password"
	Owner1TeamAdminPassword = "owner1adminpassword"
	Owner1TeamUserPassword  = "owner1userpassword"
	Owner2TeamAdminPassword = "owner2adminpassword"
	Owner2TeamUserPassword  = "owner2userpassword"

	AdminToken, _           = jwtutil.GenerateJwtToken(*AdminUser)
	Owner1Token, _          = jwtutil.GenerateJwtToken(*Owner1)
	Owner1TeamAdminToken, _ = jwtutil.GenerateJwtToken(*Owner1TeamAdmin)
	Owner1TeamUserToken, _  = jwtutil.GenerateJwtToken(*Owner1TeamUser)
	Owner2Token, _          = jwtutil.GenerateJwtToken(*Owner2)
	Owner2TeamAdminToken, _ = jwtutil.GenerateJwtToken(*Owner2TeamAdmin)
	Owner2TeamUserToken, _  = jwtutil.GenerateJwtToken(*Owner2TeamUser)
	InvalidToken            = *Owner1Token + "FAKE"

	AdminBasicCreds  = "Basic " + utils.Base64EncodeString(AdminUser.UserId+":"+vars.ADMIN_PASSWORD)
	Owner1BasicCreds = "Basic " + utils.Base64EncodeString(Owner1.UserId+":"+Owner1Password)
	Owner2BasicCreds = "Basic " + utils.Base64EncodeString(Owner2.UserId+":"+Owner2Password)
)

func AppendUsers(users []entity.User, _users ...entity.User) []entity.User {
	return append(users, _users...)
}

func CreateBasicCredentials(users []entity.User, passwords []string) []string {
	var credentials = make([]string, len(users))
	for index := range credentials {
		credentials[index] = "Basic " + utils.Base64EncodeString(users[index].UserId+":"+passwords[index])
	}
	return credentials
}

func CreatePasswords(count int) []string {
	var passwords = make([]string, count)
	for index := range passwords {
		passwords[index] = id.GenerateSecurePassword()
	}

	return passwords
}

func CreateUsers(count int, props map[string]interface{}) []entity.User {
	var users = make([]entity.User, count)
	for index := range users {
		users[index] = *CreateUser(props)
	}
	return users
}

func CreateUser(props map[string]interface{}) *entity.User {
	return entity.NewUser(vars.ADMIN_USER,
		getProperty(props, "firstName", randomString(15)).(string),
		getProperty(props, "lastName", randomString(15)).(string),
		getProperty(props, "email", randomString(15)+"@zbitech.local").(string),
		getProperty(props, "role", randomValue(roleTypes)).(ztypes.Role),
		getProperty(props, "subscription", randomValue(subscriptionTypes)).(ztypes.SubscriptionLevel))
}
