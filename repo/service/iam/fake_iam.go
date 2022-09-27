package iam

import (
	"context"
	"github.com/golang-jwt/jwt"
	"github.com/zbitech/common/interfaces"
	"github.com/zbitech/common/pkg/model/entity"
	"github.com/zbitech/common/pkg/model/ztypes"
	"time"
)

type FakeIAMService struct {
	FakeRegisterUser              func(ctx context.Context, user *entity.User, invite *entity.RegistrationInvite, pass string) error
	FakeDeactivateUser            func(ctx context.Context, userid string) error
	FakeReactivateUser            func(ctx context.Context, userid string) error
	FakeGetUsers                  func(ctx context.Context) []entity.User
	FakeGetUsersByRole            func(ctx context.Context, role ztypes.Role) ([]entity.User, error)
	FakeGetUser                   func(ctx context.Context, userId string) (*entity.User, error)
	FakeGetUserByEmail            func(ctx context.Context, email string) (*entity.User, error)
	FakeUpdateUser                func(ctx context.Context, user *entity.User) error
	FakeChangePassword            func(ctx context.Context, userid, password string) error
	FakeAuthenticateUser          func(ctx context.Context, userId, password string) (*string, error)
	FakeValidateAuthToken         func(ctx context.Context, tokenString string) (jwt.Claims, *entity.User, error)
	FakeGetAPIKeys                func(ctx context.Context, userId string) ([]entity.APIKey, error)
	FakeGetAPIKey                 func(ctx context.Context, apiKey string) (*entity.APIKey, error)
	FakeCreateAPIKey              func(ctx context.Context, userid string) (*entity.APIKey, error)
	FakeDeleteAPIKey              func(ctx context.Context, apiKey string) error
	FakeStoreUserPolicy           func(ctx context.Context, p entity.UserPolicy) error
	FakeGetUserPolicy             func(ctx context.Context, userId string) (*entity.UserPolicy, error)
	FakeStoreAPIKeyPolicy         func(ctx context.Context, p entity.APIKeyPolicy) error
	FakeGetAPIKeyPolicy           func(ctx context.Context, key string) (*entity.APIKeyPolicy, error)
	FakeStoreInstancePolicy       func(ctx context.Context, p entity.InstancePolicy) error
	FakeStoreInstancePolicies     func(ctx context.Context, p []entity.InstancePolicy) error
	FakeGetInstancePolicy         func(ctx context.Context, project, instance string) (*entity.InstancePolicy, error)
	FakeGetInstanceMethodPolicy   func(ctx context.Context, project, instance, methodName string) (*entity.MethodPolicy, error)
	FakeGetInstanceMethodPolicies func(ctx context.Context, project, instance, methodCategory string) ([]entity.MethodPolicy, error)
	FakeCreateRegistrationInvite  func(ctx context.Context, invite entity.RegistrationInvite) error
	FakeGetRegistrationInvite     func(ctx context.Context, key string) (*entity.RegistrationInvite, error)
	FakeGetRegistrationInvites    func(ctx context.Context) ([]entity.RegistrationInvite, error)
	FakeUpdateRegistrationInvite  func(ctx context.Context, invite *entity.RegistrationInvite) error
	FakeGetExpiringInvitations    func(ctx context.Context, date time.Time) ([]entity.TeamMember, error)
	FakePurgeExpiredInvitations   func(ctx context.Context) (int64, error)
	FakeCreateTeam                func(ctx context.Context, team entity.Team) error
	FakeGetTeams                  func(ctx context.Context) ([]entity.Team, error)
	FakeGetTeamsByOwner           func(ctx context.Context, owner string) ([]entity.Team, error)
	FakeGetTeamsByMembership      func(ctx context.Context, email string) ([]entity.Team, error)
	FakeGetTeam                   func(ctx context.Context, teamId string) (*entity.Team, error)
	FakeUpdateTeam                func(ctx context.Context, team *entity.Team) error
	FakeDeleteTeam                func(ctx context.Context, teamId string) error
	FakeGetTeamByOwner            func(ctx context.Context, owner string) (*entity.Team, error)
	FakeGetResourceStats          func(ctx context.Context, owner string) *entity.ResourceSummary
	FakeDeleteUser                func(ctx context.Context, user *entity.User) error
	FakeAddTeamMember             func(ctx context.Context, teamId, key, email, userid string, admin bool, status ztypes.InvitationStatus) error
	FakeUpdateTeamMember          func(ctx context.Context, teamId, key, email string, admin bool) error
	FakeUpdateTeamMemberEmail     func(ctx context.Context, teamId, key, email string) error
	FakeUpdateTeamMemberRole      func(ctx context.Context, teamId, key string, admin bool) error
	FakeUpdateTeamMemberStatus    func(ctx context.Context, teamId, key string, status ztypes.InvitationStatus) error
	FakeDeleteTeamMember          func(ctx context.Context, teamId, key string) error
}

func NewFakeIAMService() interfaces.IAMServiceIF {
	return &FakeIAMService{}
}

func (f FakeIAMService) RegisterUser(ctx context.Context, user *entity.User, invite *entity.RegistrationInvite, pass string) error {
	return f.FakeRegisterUser(ctx, user, invite, pass)
}

func (f FakeIAMService) DeactivateUser(ctx context.Context, userid string) error {
	return f.FakeDeactivateUser(ctx, userid)
}

func (f FakeIAMService) ReactivateUser(ctx context.Context, userid string) error {
	return f.FakeReactivateUser(ctx, userid)
}

func (f FakeIAMService) GetUsers(ctx context.Context) []entity.User {
	return f.FakeGetUsers(ctx)
}

func (f FakeIAMService) GetUsersByRole(ctx context.Context, role ztypes.Role) ([]entity.User, error) {
	return f.FakeGetUsersByRole(ctx, role)
}

func (f FakeIAMService) GetUser(ctx context.Context, userId string) (*entity.User, error) {
	return f.FakeGetUser(ctx, userId)
}

func (f FakeIAMService) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	return f.FakeGetUserByEmail(ctx, email)
}

func (f FakeIAMService) UpdateUser(ctx context.Context, user *entity.User) error {
	return f.FakeUpdateUser(ctx, user)
}

func (f FakeIAMService) ChangePassword(ctx context.Context, userid, password string) error {
	return f.FakeChangePassword(ctx, userid, password)
}

func (f FakeIAMService) AuthenticateUser(ctx context.Context, userId, password string) (*string, error) {
	return f.FakeAuthenticateUser(ctx, userId, password)
}

func (f FakeIAMService) ValidateAuthToken(ctx context.Context, tokenString string) (jwt.Claims, *entity.User, error) {
	return f.FakeValidateAuthToken(ctx, tokenString)
}

func (f FakeIAMService) GetAPIKeys(ctx context.Context, userId string) ([]entity.APIKey, error) {
	return f.FakeGetAPIKeys(ctx, userId)
}

func (f FakeIAMService) GetAPIKey(ctx context.Context, apiKey string) (*entity.APIKey, error) {
	return f.FakeGetAPIKey(ctx, apiKey)
}

func (f FakeIAMService) CreateAPIKey(ctx context.Context, userid string) (*entity.APIKey, error) {
	return f.FakeCreateAPIKey(ctx, userid)
}

func (f FakeIAMService) DeleteAPIKey(ctx context.Context, apiKey string) error {
	return f.FakeDeleteAPIKey(ctx, apiKey)
}

func (f FakeIAMService) StoreUserPolicy(ctx context.Context, p entity.UserPolicy) error {
	return f.FakeStoreUserPolicy(ctx, p)
}

func (f FakeIAMService) GetUserPolicy(ctx context.Context, userId string) (*entity.UserPolicy, error) {
	return f.FakeGetUserPolicy(ctx, userId)
}

func (f FakeIAMService) StoreAPIKeyPolicy(ctx context.Context, p entity.APIKeyPolicy) error {
	return f.FakeStoreAPIKeyPolicy(ctx, p)
}

func (f FakeIAMService) GetAPIKeyPolicy(ctx context.Context, key string) (*entity.APIKeyPolicy, error) {
	return f.FakeGetAPIKeyPolicy(ctx, key)
}

func (f FakeIAMService) StoreInstancePolicy(ctx context.Context, p entity.InstancePolicy) error {
	return f.FakeStoreInstancePolicy(ctx, p)
}

func (f FakeIAMService) StoreInstancePolicies(ctx context.Context, p []entity.InstancePolicy) error {
	return f.FakeStoreInstancePolicies(ctx, p)
}

func (f FakeIAMService) GetInstancePolicy(ctx context.Context, project, instance string) (*entity.InstancePolicy, error) {
	return f.FakeGetInstancePolicy(ctx, project, instance)
}

func (f FakeIAMService) GetInstanceMethodPolicy(ctx context.Context, project, instance, methodName string) (*entity.MethodPolicy, error) {
	return f.FakeGetInstanceMethodPolicy(ctx, project, instance, methodName)
}

func (f FakeIAMService) GetInstanceMethodPolicies(ctx context.Context, project, instance, methodCategory string) ([]entity.MethodPolicy, error) {
	return f.FakeGetInstanceMethodPolicies(ctx, project, instance, methodCategory)
}

func (f FakeIAMService) CreateRegistrationInvite(ctx context.Context, invite entity.RegistrationInvite) error {
	return f.FakeCreateRegistrationInvite(ctx, invite)
}

func (f FakeIAMService) GetRegistrationInvite(ctx context.Context, key string) (*entity.RegistrationInvite, error) {
	return f.FakeGetRegistrationInvite(ctx, key)
}

func (f FakeIAMService) GetRegistrationInvites(ctx context.Context) ([]entity.RegistrationInvite, error) {
	return f.FakeGetRegistrationInvites(ctx)
}

func (f FakeIAMService) UpdateRegistrationInvite(ctx context.Context, invite *entity.RegistrationInvite) error {
	return f.FakeUpdateRegistrationInvite(ctx, invite)
}

func (f FakeIAMService) GetExpiringInvitations(ctx context.Context, date time.Time) ([]entity.TeamMember, error) {
	return f.FakeGetExpiringInvitations(ctx, date)
}

func (f FakeIAMService) PurgeExpiredInvitations(ctx context.Context) (int64, error) {
	return f.FakePurgeExpiredInvitations(ctx)
}

func (f FakeIAMService) CreateTeam(ctx context.Context, team entity.Team) error {
	return f.FakeCreateTeam(ctx, team)
}

func (f FakeIAMService) GetTeams(ctx context.Context) ([]entity.Team, error) {
	return f.FakeGetTeams(ctx)
}

func (f FakeIAMService) GetTeam(ctx context.Context, teamId string) (*entity.Team, error) {
	return f.FakeGetTeam(ctx, teamId)
}

func (f FakeIAMService) UpdateTeam(ctx context.Context, team *entity.Team) error {
	return f.FakeUpdateTeam(ctx, team)
}

func (f FakeIAMService) DeleteTeam(ctx context.Context, teamId string) error {
	return f.FakeDeleteTeam(ctx, teamId)
}

func (f FakeIAMService) GetTeamByOwner(ctx context.Context, owner string) (*entity.Team, error) {
	return f.FakeGetTeamByOwner(ctx, owner)
}

func (f FakeIAMService) GetTeamsByOwner(ctx context.Context, owner string) ([]entity.Team, error) {
	return f.FakeGetTeamsByOwner(ctx, owner)
}

func (f FakeIAMService) GetResourceStats(ctx context.Context, owner string) *entity.ResourceSummary {
	return f.FakeGetResourceStats(ctx, owner)
}

func (f FakeIAMService) DeleteUser(ctx context.Context, user *entity.User) error {
	return f.FakeDeleteUser(ctx, user)
}

func (f FakeIAMService) AddTeamMember(ctx context.Context, teamId, key, email, userid string, admin bool, status ztypes.InvitationStatus) error {
	return f.FakeAddTeamMember(ctx, teamId, key, email, userid, admin, status)
}

func (f FakeIAMService) UpdateTeamMember(ctx context.Context, teamId, key, email string, admin bool) error {
	return f.FakeUpdateTeamMember(ctx, teamId, key, email, admin)
}

func (f FakeIAMService) UpdateTeamMemberEmail(ctx context.Context, teamId, key, email string) error {
	return f.FakeUpdateTeamMemberEmail(ctx, teamId, key, email)
}

func (f FakeIAMService) UpdateTeamMemberRole(ctx context.Context, teamId, key string, admin bool) error {
	return f.FakeUpdateTeamMemberRole(ctx, teamId, key, admin)
}

func (f FakeIAMService) UpdateTeamMemberStatus(ctx context.Context, teamId, key string, status ztypes.InvitationStatus) error {
	return f.FakeUpdateTeamMemberStatus(ctx, teamId, key, status)
}

func (f FakeIAMService) DeleteTeamMember(ctx context.Context, teamId, key string) error {
	return f.FakeDeleteTeamMember(ctx, teamId, key)
}
