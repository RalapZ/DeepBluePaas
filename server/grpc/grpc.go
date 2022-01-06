package grpc

import (
	"github.com/RalapZ/DeepBluePaas/common/middleware"
	"github.com/go-kratos/kratos/contrib/registry/nacos/v2"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"

	"github.com/RalapZ/DeepBluePaas/server/config"
)

func Start() {
	//grpc.NewServer()
	//fmt.Printf("grpc:-----%##v\n",config.ServerInfo.Grpc)
	serverConfig := []constant.ServerConfig{
		*constant.NewServerConfig(config.NacosSerConfig.IpAddr, config.NacosSerConfig.Port),
	}
	clientConfig := config.NacosCliConfig

	//fmt.Printf("nacos客户端配置 %##v \n",clientConfig)
	//fmt.Printf("nacos 服务端配置 %##v\n",config.NacosSerConfig)

	client, err := clients.NewNamingClient(
		vo.NacosClientParam{
			&clientConfig,
			serverConfig,
		},
	)
	if err != nil {
		middleware.LogGlobal.Error(err.Error())
	}

	grpcSrv := grpc.NewServer(
		grpc.Address(config.IPAddr),
		//grpc.Address(":9000"),
		grpc.Middleware(
			recovery.Recovery(),
		),
	)
	//httpSrv := http.NewServer(http.(":8000"))  //注册http服务

	registry := nacos.New(client)

	app := kratos.New(
		kratos.Name(config.ServerInfo.Name),
		kratos.Version("1.4.0"),
		//kratos.
		kratos.Server(grpcSrv),
		kratos.Registrar(registry),
	)
	app.Run()
	if err != nil {
		panic(err)
	}

}
