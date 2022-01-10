package dbgrpc

import (
	"fmt"
	"github.com/RalapZ/DeepBluePaas/server/config"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"net"
	"time"
)


func db_grpc_RegisterServiceInstance(client naming_client.INamingClient, param vo.RegisterInstanceParam) {
	success, _ := client.RegisterInstance(param)
	fmt.Printf("RegisterServiceInstance,param:%+v,result:%+v \n\n", param, success)
}

func db_grpc_DeRegisterServiceInstance(client naming_client.INamingClient, param vo.DeregisterInstanceParam) {
	success, _ := client.DeregisterInstance(param)
	fmt.Printf("DeRegisterServiceInstance,param:%+v,result:%+v \n\n", param, success)
}

func db_grpc_GetService(client naming_client.INamingClient, param vo.GetServiceParam) {
	service, _ := client.GetService(param)
	fmt.Printf("GetService,param:%+v, result:%+v \n\n", param, service)
}

func db_grpc_SelectAllInstances(client naming_client.INamingClient, param vo.SelectAllInstancesParam) {
	instances, _ := client.SelectAllInstances(param)
	fmt.Printf("SelectAllInstance,param:%+v, result:%+v \n\n", param, instances)
}

func db_grpc_SelectInstances(client naming_client.INamingClient, param vo.SelectInstancesParam) {
	instances, _ := client.SelectInstances(param)
	fmt.Printf("SelectInstances,param:%+v, result:%+v \n\n", param, instances)
}

func db_grpc_SelectOneHealthyInstance(client naming_client.INamingClient, param vo.SelectOneHealthInstanceParam) {
	instances, _ := client.SelectOneHealthyInstance(param)
	fmt.Printf("SelectInstances,param:%+v, result:%+v \n\n", param, instances)
}

func db_grpc_Subscribe(client naming_client.INamingClient, param *vo.SubscribeParam) {
	client.Subscribe(param)
}

func db_grpc_UnSubscribe(client naming_client.INamingClient, param *vo.SubscribeParam) {
	client.Unsubscribe(param)
}

func db_grpc_GetAllService(client naming_client.INamingClient, param vo.GetAllServiceInfoParam) {
	service, _ := client.GetAllServicesInfo(param)
	fmt.Printf("GetAllService,param:%+v, result:%+v \n\n", param, service)
}



func GrpcRegister(){

	for {
		//if
		serverConfig := []constant.ServerConfig{
			{
				IpAddr: config.NacosSerConfig.IpAddr,
				Port: config.NacosSerConfig.Port,
			},
		}

		clientConfig := config.NacosCliConfig
		client, err := clients.NewNamingClient(
			vo.NacosClientParam{
				&clientConfig,
				serverConfig,
			},
		)
		if err != nil {
			panic(err)
		}

		//Register with default cluster and group
		//ClusterName=DEFAULT,GroupName=DEFAULT_GROUP
		db_grpc_RegisterServiceInstance(client, config.NacosInsConfig)

		time.Sleep(5)
	}
}

func Start() {
	server := grpc.NewServer(grpc.Address("0.0.0.0"),grpc.Middleware(recovery.Recovery()))
	go GrpcRegister()
	listen, _ := net.Listen("tcp", ":28080")
	fmt.Println("grpc start")
	err2 := server.Serve(listen)
	if err2 != nil{
		panic("grpc start faild")
	}
	//serverConfig := []constant.ServerConfig{
	//	*constant.NewServerConfig(config.NacosSerConfig.IpAddr, config.NacosSerConfig.Port),
	//}
	//clientConfig := config.NacosCliConfig
	//
	//client, err := clients.NewNamingClient(
	//	vo.NacosClientParam{
	//		&clientConfig,
	//		serverConfig,
	//	},
	//)
	//if err != nil {
	//	middleware.LogGlobal.Error(err.Error())
	//}
	//
	//grpcSrv := grpc.NewServer(
	//	grpc.Address(config.IPAddr),
	//	grpc.Middleware(
	//		recovery.Recovery(),
	//	),
	//)
	////httpSrv := http.NewServer(http.(":8000"))  //注册http服务
	//
	//registry := nacos.New(client)
	//
	//app := kratos.New(
	//	kratos.Name(config.ServerInfo.Name),
	//	kratos.Version("1.4.0"),
	//	//kratos.
	//	kratos.Server(grpcSrv),
	//	kratos.Registrar(registry),
	//)
	//fmt.Printf("%##v",app)
	//app.Run()
	//if err != nil {
	//	panic(err)
	//}

}