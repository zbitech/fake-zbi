package data

import (
	"encoding/json"
	"fmt"
	"github.com/zbitech/common/pkg/logger"
	"github.com/zbitech/common/pkg/model/spec"
	"github.com/zbitech/common/pkg/model/ztypes"
	"github.com/zbitech/common/pkg/rctx"
	"github.com/zbitech/common/pkg/utils"
	"io/ioutil"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

var (
	DATA_PATH = utils.GetEnv("DATA_PATH", "./")
)

func getArrayData(filePath string, obj []spec.ResourceObject) error {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	logger.Debugf(rctx.CTX, "File Data: %s", string(file))
	if err = json.Unmarshal(file, &obj); err != nil {
		return err
	}
	logger.Debugf(rctx.CTX, "READ Data: %s", utils.MarshalObject(obj))
	return nil
}

func getData(filePath string, obj interface{}) error {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	//	logger.Debugf(rctx.CTX, "File Data: %s", string(file))
	if err = json.Unmarshal(file, &obj); err != nil {
		return err
	}
	//	logger.Debugf(rctx.CTX, "READ Data: %s", utils.MarshalObject(obj))
	return nil
}

//func CreateDeploymentResource() unstructured.Unstructured {
//
//}

func GetInstanceResources(iType ztypes.InstanceType, rType ztypes.ResourceObjectType, size int) ([]spec.ResourceObject, error) {

	var filePath string
	switch rType {
	case ztypes.ResourcePersistentVolumeClaim:
		filePath = fmt.Sprintf("%s/data/manifest/%s/pvc.json", DATA_PATH, string(iType)) //"data/manifest/zcash/pvc.json"
	}

	var objects = make([]spec.ResourceObject, size)
	//	for index := 0; index < size; index++ {
	//		objects[index] = new(unstructured.Unstructured)
	//	}

	//	filePath = DATA_PATH + filePath
	if err := getArrayData(filePath, objects); err != nil {
		return nil, err
	}

	return objects, nil
}

func GetGenericResource(rType ztypes.ResourceObjectType) (*unstructured.Unstructured, error) {
	var obj = new(unstructured.Unstructured)
	var filePath string

	switch rType {
	case ztypes.ResourceNamespace:
		filePath = "data/manifest/generic/namespace.json"
	case ztypes.ResourceDeployment:
		filePath = "data/manifest/generic/deployment.json"
	case ztypes.ResourcePod:
		filePath = "data/manifest/generic/pod.json"
	case ztypes.ResourceConfigMap:
		filePath = "data/manifest/generic/configmap.json"
	case ztypes.ResourceSecret:
		filePath = "data/manifest/generic/secret.json"
	case ztypes.ResourceService:
		filePath = "data/manifest/generic/service.json"
	case ztypes.ResourcePersistentVolumeClaim:
		filePath = "data/manifest/generic/pvc.json"
	case ztypes.ResourceVolumeSnapshot:
		filePath = "data/manifest/generic/snapshot.json"
	case ztypes.ResourceSnapshotSchedule:
		filePath = "data/manifest/generic/schedule.json"
	case ztypes.ResourceHTTPProxy:
		filePath = "data/manifest/generic/ingress.json"
	}

	filePath = DATA_PATH + filePath
	if err := getData(filePath, obj); err != nil {
		return nil, err
	}
	return obj, nil
}
