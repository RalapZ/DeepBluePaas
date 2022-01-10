package config

import "github.com/RalapZ/DeepBluePaas/server/errmsg"

func NacosGetInstanceConfig() (errorcode int) {
	NacosInsConfig.Ip = LocalAddr
	NacosInsConfig.Port = ServerInfo.Grpc
	//NacosInsConfig.ClusterName =
	//NacosInsConfig.GroupName =
	NacosInsConfig.ServiceName = Config.GetString("registry.nacos.instance.servicename")
	NacosInsConfig.Weight = Config.GetFloat64("registry.nacos.instance.weight")
	NacosInsConfig.Enable = Config.GetBool("registry.nacos.instance.enable")
	NacosInsConfig.Healthy = Config.GetBool("registry.nacos.instance.healthy")
	NacosInsConfig.Ephemeral = Config.GetBool("registry.nacos.instance.ephemeral")
	NacosInsConfig.Metadata = Config.GetStringMapString("registry.nacos.instance.metadata")
	return errmsg.GLOBAL_STATS_OK
}


func NacosGetConfig()(errorcode int){
	if ok:= Config.IsSet("registry.nacos.client.namespaceid");!ok{
		return errmsg.REGISTRY_CONFIG_NOT_ERROR
	}

	NacosSerConfig.IpAddr = Config.GetString("registry.nacos.server.addr")
	NacosSerConfig.Port = Config.GetUint64("registry.nacos.server.port")

	NacosCliConfig.NamespaceId = Config.GetString("registry.nacos.client.namespaceid")
	NacosCliConfig.TimeoutMs = Config.GetUint64("registry.nacos.client.timeoutms")
	NacosCliConfig.NotLoadCacheAtStart = Config.GetBool("registry.nacos.client.notLoadcacheatstart")
	NacosCliConfig.LogDir = Config.GetString("registry.nacos.client.logdir")
	NacosCliConfig.CacheDir = Config.GetString("registry.nacos.client.cachedir")
	NacosCliConfig.RotateTime = Config.GetString("registry.nacos.client.rotatetime")
	NacosCliConfig.MaxAge = Config.GetInt64("registry.nacos.client.maxage")
	NacosCliConfig.LogLevel =  Config.GetString("registry.nacos.client.loglevel")
	return errmsg.GLOBAL_STATS_OK
}
