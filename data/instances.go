package data

import (
	"github.com/zbitech/common/pkg/model/entity"
	"github.com/zbitech/common/pkg/model/ztypes"
	"time"
)

var (
	Instance1 = &entity.ZcashInstance{
		Instance: entity.Instance{Project: Project1.Name, Name: "instance1", Version: "v1", Network: ztypes.NetworkTypeTest, Description: "Instance One", Owner: Owner1.UserId, Status: "New", Timestamp: time.Now(), InstanceType: ztypes.InstanceTypeZCASH, DataSourceType: ztypes.NoDataSource},
		ZcashDetails: entity.ZcashDetails{
			TransactionIndex: false,
			Miner:            false,
			Peers:            []string{},
			//			DataVolume:       entity.DataVolume{Name: "zcash-data-i1", Volume: "data", Size: 10},
			//			ParamsVolume:     entity.DataVolume{Name: "zcash-params-i1", Volume: "params", Size: 3},
		},
	}
	Instance2 = &entity.ZcashInstance{
		Instance: entity.Instance{Project: Project1.Name, Name: "instance2", Version: "v1", Network: ztypes.NetworkTypeTest, Description: "Instance Two", Owner: Owner1.UserId, Status: "New", Timestamp: time.Now(), InstanceType: ztypes.InstanceTypeZCASH, DataSourceType: ztypes.NoDataSource},
		ZcashDetails: entity.ZcashDetails{
			TransactionIndex: false,
			Miner:            false,
			Peers:            []string{},
			//			DataVolume:       entity.DataVolume{Name: "zcash-data-i2", Volume: "data", Size: 10},
			//			ParamsVolume:     entity.DataVolume{Name: "zcash-params-i2", Volume: "params", Size: 3},
		},
	}
	Instance3 = &entity.ZcashInstance{
		Instance: entity.Instance{Project: Project2.Name, Name: "instance3", Version: "v1", Network: ztypes.NetworkTypeTest, Description: "Instance Three", Owner: Owner1.UserId, Status: "New", Timestamp: time.Now(), InstanceType: ztypes.InstanceTypeZCASH, DataSourceType: ztypes.NoDataSource},
		ZcashDetails: entity.ZcashDetails{
			TransactionIndex: false,
			Miner:            false,
			Peers:            []string{},
			//			DataVolume:       entity.DataVolume{Name: "zcash-data-i3", Volume: "data", Size: 10},
			//			ParamsVolume:     entity.DataVolume{Name: "zcash-params-i3", Volume: "params", Size: 3},
		},
	}
	Instance4 = &entity.ZcashInstance{
		Instance: entity.Instance{Project: Project2.Name, Name: "instance4", Version: "v1", Network: ztypes.NetworkTypeTest, Description: "Instance Four", Owner: Owner1.UserId, Status: "New", Timestamp: time.Now(), InstanceType: ztypes.InstanceTypeZCASH, DataSourceType: ztypes.NoDataSource},
		ZcashDetails: entity.ZcashDetails{
			TransactionIndex: false,
			Miner:            false,
			Peers:            []string{},
			//			DataVolume:       entity.DataVolume{Name: "zcash-data-i4", Volume: "data", Size: 10},
			//			ParamsVolume:     entity.DataVolume{Name: "zcash-params-i4", Volume: "params", Size: 3},
		},
	}
	Instance5 = &entity.ZcashInstance{
		Instance: entity.Instance{Project: Project3.Name, Name: "instance5", Version: "v1", Network: ztypes.NetworkTypeTest, Description: "Instance Five", Owner: Owner2.UserId, Status: "New", Timestamp: time.Now(), InstanceType: ztypes.InstanceTypeZCASH, DataSourceType: ztypes.NoDataSource},
		ZcashDetails: entity.ZcashDetails{
			TransactionIndex: false,
			Miner:            false,
			Peers:            []string{},
			//			DataVolume:       entity.DataVolume{Name: "zcash-data-i5", Volume: "data", Size: 10},
			//			ParamsVolume:     entity.DataVolume{Name: "zcash-params-i5", Volume: "params", Size: 3},
		},
	}
	Instance6 = &entity.ZcashInstance{
		Instance: entity.Instance{Project: Project3.Name, Name: "instance6", Version: "v1", Network: ztypes.NetworkTypeTest, Description: "Instance Six", Owner: Owner2.UserId, Status: "New", Timestamp: time.Now(), InstanceType: ztypes.InstanceTypeZCASH, DataSourceType: ztypes.NoDataSource},
		ZcashDetails: entity.ZcashDetails{
			TransactionIndex: false,
			Miner:            false,
			Peers:            []string{},
			//			DataVolume:       entity.DataVolume{Name: "zcash-data-i6", Volume: "data", Size: 10},
			//			ParamsVolume:     entity.DataVolume{Name: "zcash-params-i6", Volume: "params", Size: 3},
		},
	}
	Instance7 = &entity.ZcashInstance{
		Instance: entity.Instance{Project: Project4.Name, Name: "instance7", Version: "v1", Network: ztypes.NetworkTypeTest, Description: "Instance Seven", Owner: Owner2.UserId, Status: "New", Timestamp: time.Now(), InstanceType: ztypes.InstanceTypeZCASH, DataSourceType: ztypes.NoDataSource},
		ZcashDetails: entity.ZcashDetails{
			TransactionIndex: false,
			Miner:            false,
			Peers:            []string{},
			//			DataVolume:       entity.DataVolume{Name: "zcash-data-i7", Volume: "data", Size: 10},
			//			ParamsVolume:     entity.DataVolume{Name: "zcash-params-i7", Volume: "params", Size: 3},
		},
	}
	Instance8 = &entity.ZcashInstance{
		Instance: entity.Instance{Project: Project4.Name, Name: "instance8", Version: "v1", Network: ztypes.NetworkTypeTest, Description: "Instance Eight", Owner: Owner2.UserId, Status: "New", Timestamp: time.Now(), InstanceType: ztypes.InstanceTypeZCASH, DataSourceType: ztypes.NoDataSource},
		ZcashDetails: entity.ZcashDetails{
			TransactionIndex: false,
			Miner:            false,
			Peers:            []string{},
			//			DataVolume:       entity.DataVolume{Name: "zcash-data-i8", Volume: "data", Size: 10},
			//			ParamsVolume:     entity.DataVolume{Name: "zcash-params-i8", Volume: "params", Size: 3},
		},
	}
	LwdInstance1 = &entity.LWDInstance{
		Instance: entity.Instance{Project: Project1.Name, Name: "instance10", Version: "v1", Network: ztypes.NetworkTypeTest, Description: "Instance Ten", Owner: Owner1.UserId, Status: "New", Timestamp: time.Now(), InstanceType: ztypes.InstanceTypeLWD, DataSourceType: ztypes.NoDataSource},
		LWDDetails: entity.LWDDetails{
			//			DataVolume:    entity.DataVolume{Name: "lwd-data-i1", Volume: "data", Size: 10},
			ZcashInstance: "zcash-svc-" + Instance1.Name + ".svc.cluster.local",
		},
	}
	LwdInstance2 = &entity.LWDInstance{
		Instance: entity.Instance{Project: Project1.Name, Name: "instance11", Version: "v1", Network: ztypes.NetworkTypeTest, Description: "Instance Eleven", Owner: Owner1.UserId, Status: "New", Timestamp: time.Now(), InstanceType: ztypes.InstanceTypeLWD, DataSourceType: ztypes.NoDataSource},
		LWDDetails: entity.LWDDetails{
			//			DataVolume:    entity.DataVolume{Name: "lwd-data-i2", Volume: "data", Size: 10},
			ZcashInstance: "zcash-svc-" + Instance2.Name + ".svc.cluster.local",
		},
	}

	//Instance1 = entity.Instance{Project: Project1.Name, Name: "instance1", Version: "v1", Network: ztypes.NetworkTypeTest, Description: "Instance One", Owner: Owner1.UserId, Object: "New", Timestamp: time.Time{}, InstanceType: ztypes.InstanceTypeZCASH, InstanceDetail: &entity.ZcashInstanceDetail{TransactionIndex: false, Miner: false, Peers: []string{}}}
	//Instance2 = entity.Instance{Project: Project1.Name, Name: "instance2", Version: "v1", Network: ztypes.NetworkTypeTest, Description: "Instance Two", Owner: Owner1.UserId, Object: "New", Timestamp: time.Time{}, InstanceType: ztypes.InstanceTypeZCASH, InstanceDetail: &entity.ZcashInstanceDetail{TransactionIndex: false, Miner: false, Peers: []string{}}}
	//Instance3 = entity.Instance{Project: Project2.Name, Name: "instance3", Version: "v1", Network: ztypes.NetworkTypeTest, Description: "Instance Three", Owner: Owner1.UserId, Object: "New", Timestamp: time.Time{}, InstanceType: ztypes.InstanceTypeZCASH, InstanceDetail: &entity.ZcashInstanceDetail{TransactionIndex: false, Miner: false, Peers: []string{}}}
	//Instance4 = entity.Instance{Project: Project2.Name, Name: "instance4", Version: "v1", Network: ztypes.NetworkTypeTest, Description: "Instance Four", Owner: Owner1.UserId, Object: "New", Timestamp: time.Time{}, InstanceType: ztypes.InstanceTypeZCASH, InstanceDetail: &entity.ZcashInstanceDetail{TransactionIndex: false, Miner: false, Peers: []string{}}}
	//Instance5 = entity.Instance{Project: Project3.Name, Name: "instance5", Version: "v1", Network: ztypes.NetworkTypeTest, Description: "Instance Five", Owner: Owner2.UserId, Object: "New", Timestamp: time.Time{}, InstanceType: ztypes.InstanceTypeZCASH, InstanceDetail: &entity.ZcashInstanceDetail{TransactionIndex: false, Miner: false, Peers: []string{}}}
	//Instance6 = entity.Instance{Project: Project3.Name, Name: "instance6", Version: "v1", Network: ztypes.NetworkTypeTest, Description: "Instance Six", Owner: Owner2.UserId, Object: "New", Timestamp: time.Time{}, InstanceType: ztypes.InstanceTypeZCASH, InstanceDetail: &entity.ZcashInstanceDetail{TransactionIndex: false, Miner: false, Peers: []string{}}}
	//Instance7 = entity.Instance{Project: Project4.Name, Name: "instance7", Version: "v1", Network: ztypes.NetworkTypeTest, Description: "Instance Seven", Owner: Owner2.UserId, Object: "New", Timestamp: time.Time{}, InstanceType: ztypes.InstanceTypeZCASH, InstanceDetail: &entity.ZcashInstanceDetail{TransactionIndex: false, Miner: false, Peers: []string{}}}
	//Instance8 = entity.Instance{Project: Project4.Name, Name: "instance8", Version: "v1", Network: ztypes.NetworkTypeTest, Description: "Instance Eight", Owner: Owner2.UserId, Object: "New", Timestamp: time.Time{}, InstanceType: ztypes.InstanceTypeZCASH, InstanceDetail: &entity.ZcashInstanceDetail{TransactionIndex: false, Miner: false, Peers: []string{}}}

	Instances       = []entity.InstanceIF{Instance1, Instance2, Instance3, Instance4, Instance5, Instance6, Instance7, Instance8}
	Owner1Instances = []entity.InstanceIF{Instance1, Instance2, Instance3, Instance4}
	Owner2Instances = []entity.InstanceIF{Instance5, Instance6, Instance7, Instance8}

	Project1Instances = []entity.InstanceIF{Instance1, Instance2}
	Project2Instances = []entity.InstanceIF{Instance3, Instance4}
	Project3Instances = []entity.InstanceIF{Instance6, Instance6}
	Project4Instances = []entity.InstanceIF{Instance7, Instance8}
)

func AppendInstances(instances []entity.InstanceIF, _instances ...entity.InstanceIF) []entity.InstanceIF {
	return append(instances, _instances...)
}

func CreateInstances(project *entity.Project, count int, iType ztypes.InstanceType, props map[string]interface{}) []entity.InstanceIF {
	var instances = make([]entity.InstanceIF, count)
	for index := range instances {
		instances[index] = CreateInstance(project, iType, props)
	}
	return instances
}

func CreateInstance(project *entity.Project, iType ztypes.InstanceType, props map[string]interface{}) entity.InstanceIF {

	switch iType {
	case ztypes.InstanceTypeZCASH:
		return CreateZcashInstance(project, props)
	case ztypes.InstanceTypeLWD:
		return CreateLWDInstance(project, props)
	}

	return nil
}

func createInstance(project *entity.Project, iType ztypes.InstanceType, props map[string]interface{}) entity.Instance {
	return entity.Instance{
		Project:        project.GetName(),
		Name:           getProperty(props, "name", randomString(10)).(string),
		Version:        getProperty(props, "version", randomValue(versions)).(string),
		Network:        project.GetNetwork(),
		Description:    getProperty(props, "description", randomString(30)).(string),
		Owner:          project.GetOwner(),
		Status:         getProperty(props, "status", ztypes.StatusActive).(ztypes.StatusType),
		Timestamp:      getProperty(props, "timeStamp", time.Now()).(time.Time),
		DataVolumeType: getProperty(props, "dataVolumeType", randomValue(volumeTypes)).(ztypes.DataVolumeType),
		DataSourceType: getProperty(props, "dataSourceType", randomValue(sourceTypes)).(ztypes.DataSourceType),
		DataSource:     getProperty(props, "dataSource", randomString(30)).(string),
		InstanceType:   iType,
	}
}

func CreateZcashInstance(project *entity.Project, props map[string]interface{}) *entity.ZcashInstance {

	return &entity.ZcashInstance{
		Instance: createInstance(project, ztypes.InstanceTypeZCASH, props),
		ZcashDetails: entity.ZcashDetails{
			TransactionIndex: getProperty(props, "transactionIndex", randomValue(boolean)).(bool),
			Miner:            getProperty(props, "miner", randomValue(boolean)).(bool),
		},
	}
}

func CreateLWDInstance(project *entity.Project, props map[string]interface{}) *entity.LWDInstance {

	return &entity.LWDInstance{
		Instance: createInstance(project, ztypes.InstanceTypeLWD, props),
		LWDDetails: entity.LWDDetails{
			ZcashInstance: getProperty(props, "zcashInstance", "").(string),
		},
	}
}
