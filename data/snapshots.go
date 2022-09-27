package data

import (
	"github.com/zbitech/common/pkg/model/entity"
	"time"
)

func CreateInstantSnapshot(instance entity.InstanceIF, props map[string]interface{}) *entity.SnapshotResource {

	return &entity.SnapshotResource{
		Name:     getProperty(props, "name", randomString(10)).(string),
		Project:  instance.GetProject(),
		Instance: instance.GetName(),
		Created:  getProperty(props, "created", time.Now()).(time.Time),
		Status:   getProperty(props, "status", randomValue(resourceStatus)).(string),
	}
}

func CreateScheduledSnapshot(instance entity.InstanceIF, props map[string]interface{}) *entity.SnapshotResource {

	return &entity.SnapshotResource{
		Name:     getProperty(props, "name", randomString(10)).(string),
		Project:  instance.GetProject(),
		Instance: instance.GetName(),
		Created:  getProperty(props, "created", time.Now()).(time.Time),
		Status:   getProperty(props, "status", randomValue(resourceStatus)).(string),
	}
}
