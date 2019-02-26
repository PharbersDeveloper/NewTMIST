package BmFactory

import (
	"github.com/PharbersDeveloper/NewTMIST/BmDaemons/BmMongodb"
	"github.com/PharbersDeveloper/NewTMIST/BmDaemons/BmRedis"
	"github.com/PharbersDeveloper/NewTMIST/BmHandler"
	"github.com/PharbersDeveloper/NewTMIST/BmMiddleware"
)

var BLACKMIRROR_MODEL_FACTORY = map[string]interface{}{
}

var BLACKMIRROR_RESOURCE_FACTORY = map[string]interface{}{
}

var BLACKMIRROR_STORAGE_FACTORY = map[string]interface{}{
}

var BLACKMIRROR_DAEMON_FACTORY = map[string]interface{}{
	"BmMongodbDaemon": BmMongodb.BmMongodb{},
	"BmRedisDaemon":   BmRedis.BmRedis{},
}

var BLACKMIRROR_FUNCTION_FACTORY = map[string]interface{}{
	"BmUploadToOssHandler": BmHandler.UploadToOssHandler{},
}

var BLACKMIRROR_MIDDLEWARE_FACTORY = map[string]interface{}{
	"BmCheckTokenMiddleware": BmMiddleware.CheckTokenMiddleware{},
}

func GetModelByName(name string) interface{} {
	return BLACKMIRROR_MODEL_FACTORY[name]
}

func GetResourceByName(name string) interface{} {
	return BLACKMIRROR_RESOURCE_FACTORY[name]
}

func GetStorageByName(name string) interface{} {
	return BLACKMIRROR_STORAGE_FACTORY[name]
}

func GetDaemonByName(name string) interface{} {
	return BLACKMIRROR_DAEMON_FACTORY[name]
}

func GetFunctionByName(name string) interface{} {
	return BLACKMIRROR_FUNCTION_FACTORY[name]
}

func GetMiddlewareByName(name string) interface{} {
	return BLACKMIRROR_MIDDLEWARE_FACTORY[name]
}
