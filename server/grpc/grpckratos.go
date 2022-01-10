package dbgrpc

import (
	"fmt"
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

func Startkratos() {
	serverConfig := []constant.ServerConfig{
		*constant.NewServerConfig(config.NacosSerConfig.IpAddr, config.NacosSerConfig.Port),
	}
	clientConfig := config.NacosCliConfig

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
	fmt.Printf("%##v",app)
	app.Run()
	if err != nil {
		panic(err)
	}

}
