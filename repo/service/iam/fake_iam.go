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
	FakeCreateProfileEvent        func(ctx context.Context, userId, actor string, action ztypes.EventAction, evtError error) error
	FakeGetProfileEvents          func(ctx context.Context, userId string) ([]entity.ProfileEvent, error)
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
	FakeGetTeamMembers            func(ctx context.Context, teamId string) ([]entity.TeamMember, error)
	FakeAddTeamMember             func(ctx context.Context, teamId string, member entity.TeamMember) error
	FakeRemoveTeamMember          func(ctx context.Context, teamId string, key string) error
	FakeUpdateTeamMember          func(ctx context.Context, member *entity.TeamMember) error
	FakeGetAllMemberships         func(ctx context.Context) ([]entity.TeamMember, error)
	FakeGetTeamMemberships        func(ctx context.Context, email string) ([]entity.TeamMember, error)
	FakeGetTeamMembership         func(ctx context.Context, key string) (*entity.TeamMember, error)
	FakeGetTeamMembershipByEmail  func(ctx context.Context, team, email string) (*entity.TeamMember, error)
	FakeCreateTeamEvent           func(ctx context.Context, teamId, userId string, action ztypes.EventAction, evtError error) error
	FakeGetTeamEvents             func(ctx context.Context, teamId string) ([]entity.TeamEvent, error)
	FakeGetResourceStats          func(ctx context.Context, owner string) *entity.ResourceSummary
	FakeUpdateTeamMembershipEmail func(ctx context.Context, oldEmail, newEmail string) error
	FakeGetTeamMembershipByKey    func(ctx context.Context, key string) (*entity.TeamMember, error)
	FakeGetTeamMembershipsByEmail func(ctx context.Context, email string) ([]entity.TeamMembership, error)
	FakeGetTeamMembershipByUser   func(ctx context.Context, team, userId string) (*entity.TeamMember, error)
	FakeGetTeamMembershipsByUser  func(ctx context.Context, userId string) ([]entity.TeamMembership, error)
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

func (f FakeIAMService) CreateProfileEvent(ctx context.Context, userId, actor string, action ztypes.EventAction, evtError error) error {
	return f.FakeCreateProfileEvent(ctx, userId, actor, action, evtError)
}

func (f FakeIAMService) GetProfileEvents(ctx context.Context, userId string) ([]entity.ProfileEvent, error) {
	return f.FakeGetProfileEvents(ctx, userId)
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

func (f FakeIAMService) GetTeamMembers(ctx context.Context, teamId string) ([]entity.TeamMember, error) {
	return f.FakeGetTeamMembers(ctx, teamId)
}

func (f FakeIAMService) AddTeamMember(ctx context.Context, teamId string, member entity.TeamMember) error {
	return f.FakeAddTeamMember(ctx, teamId, member)
}

func (f FakeIAMService) RemoveTeamMember(ctx context.Context, teamId string, key string) error {
	return f.FakeRemoveTeamMember(ctx, teamId, key)
}

func (f FakeIAMService) UpdateTeamMember(ctx context.Context, member *entity.TeamMember) error {
	return f.FakeUpdateTeamMember(ctx, member)
}

func (f FakeIAMService) GetAllMemberships(ctx context.Context) ([]entity.TeamMember, error) {
	return f.FakeGetAllMemberships(ctx)
}

func (f FakeIAMService) GetTeamMemberships(ctx context.Context, email string) ([]entity.TeamMember, error) {
	return f.FakeGetTeamMemberships(ctx, email)
}

func (f FakeIAMService) GetTeamMembership(ctx context.Context, key string) (*entity.TeamMember, error) {
	return f.FakeGetTeamMembership(ctx, key)
}

func (f FakeIAMService) GetTeamMembershipByEmail(ctx context.Context, team, email string) (*entity.TeamMember, error) {
	return f.FakeGetTeamMembershipByEmail(ctx, team, email)
}

func (f FakeIAMService) GetTeamsByOwner(ctx context.Context, owner string) ([]entity.Team, error) {
	return f.FakeGetTeamsByOwner(ctx, owner)
}

func (f FakeIAMService) GetTeamsByMembership(ctx context.Context, email string) ([]entity.Team, error) {
	return f.FakeGetTeamsByMembership(ctx, email)
}

func (f FakeIAMService) UpdateTeamMembershipEmail(ctx context.Context, oldEmail, newEmail string) error {
	return f.FakeUpdateTeamMembershipEmail(ctx, oldEmail, newEmail)
}

func (f FakeIAMService) GetTeamMembershipByKey(ctx context.Context, key string) (*entity.TeamMember, error) {
	return f.FakeGetTeamMembershipByKey(ctx, key)
}

func (f FakeIAMService) GetTeamMembershipsByEmail(ctx context.Context, email string) ([]entity.TeamMembership, error) {
	return f.FakeGetTeamMembershipsByEmail(ctx, email)
}

func (f FakeIAMService) GetTeamMembershipByUser(ctx context.Context, team, userId string) (*entity.TeamMember, error) {
	return f.FakeGetTeamMembershipByUser(ctx, team, userId)
}

func (f FakeIAMService) GetTeamMembershipsByUser(ctx context.Context, userId string) ([]entity.TeamMembership, error) {
	return f.FakeGetTeamMembershipsByUser(ctx, userId)
}

func (f FakeIAMService) CreateTeamEvent(ctx context.Context, teamId, actor string, action ztypes.EventAction, evtError error, target ...string) error {
	return f.FakeCreateTeamEvent(ctx, teamId, actor, action, evtError)
}

func (f FakeIAMService) GetTeamEvents(ctx context.Context, teamId string) ([]entity.TeamEvent, error) {
	return f.FakeGetTeamEvents(ctx, teamId)
}

func (f FakeIAMService) GetResourceStats(ctx context.Context, owner string) *entity.ResourceSummary {
	return f.FakeGetResourceStats(ctx, owner)
}
