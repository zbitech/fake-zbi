package data

import (
	"github.com/zbitech/common/pkg/model/ztypes"
	"math/rand"
)

var (
	letters           = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	boolean           = []interface{}{true, false}
	volumeTypes       = []interface{}{ztypes.EphemeralDataVolume, ztypes.PersistentDataVolume}
	sourceTypes       = []interface{}{ztypes.NewDataSource, ztypes.VolumeDataSource, ztypes.SnapshotDataSource}
	versions          = []interface{}{"v1"}
	networkTypes      = []interface{}{ztypes.NetworkTypeTest}
	instanceTypes     = []interface{}{ztypes.InstanceTypeZCASH, ztypes.InstanceTypeLWD}
	roleTypes         = []interface{}{ztypes.RoleUser, ztypes.RoleOwner}
	subscriptionTypes = []interface{}{ztypes.SubscriptionTeamMember, ztypes.SubscriptionBronzeLevel}
)

func randomString(n int) string {
	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func randomValue(array []interface{}) interface{} {
	return array[rand.Intn(len(array))]
}

func getProperty(props map[string]interface{}, key string, _default interface{}) interface{} {

	if props != nil {
		val, ok := props[key]
		if ok {
			return val
		}
	}

	return _default
}
