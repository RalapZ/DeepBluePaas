package microservice

import (
	"fmt"

	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

type INaocsInt interface{

}


//注册服务
func Nacos_RegisterServiceInstance(client naming_client.INamingClient, param vo.RegisterInstanceParam, err error) {
	success, err := client.RegisterInstance(param)
	fmt.Printf("RegisterServiceInstance,param:%+v,result:%+v \n\n", param, success)
}

//反向注册服务
func Nacos_DeRegisterServiceInstance(client naming_client.INamingClient, param vo.DeregisterInstanceParam, err error) {
	success, err := client.DeregisterInstance(param)
	fmt.Printf("DeRegisterServiceInstance,param:%+v,result:%+v \n\n", param, success)
}

//获取服务信息
func Nacos_GetService(client naming_client.INamingClient, param vo.GetServiceParam, err error) {
	service, err := client.GetService(param)
	fmt.Printf("GetService,param:%+v, result:%+v \n\n", param, service)
}

//获取所有实例
func Nacos_SelectAllInstances(client naming_client.INamingClient, param vo.SelectAllInstancesParam, err error) {
	instances, err := client.SelectAllInstances(param)
	fmt.Printf("SelectAllInstance,param:%+v, result:%+v \n\n", param, instances)
}

//获取服务实例
func Nacos_SelectInstances(client naming_client.INamingClient, param vo.SelectInstancesParam, err error) {
	instances, err := client.SelectInstances(param)
	fmt.Printf("SelectInstances,param:%+v, result:%+v \n\n", param, instances)
}

//获取一个健康服务实例
func Nacos_SelectOneHealthyInstance(client naming_client.INamingClient, param vo.SelectOneHealthInstanceParam, err error) {
	instances, err := client.SelectOneHealthyInstance(param)
	fmt.Printf("SelectInstances,param:%+v, result:%+v \n\n", param, instances)
}

//订阅服务
func Nacos_Subscribe(client naming_client.INamingClient, param *vo.SubscribeParam)(err error) {
	err = client.Subscribe(param)
	return
}

//取消订阅服务
func Nacos_UnSubscribe(client naming_client.INamingClient, param *vo.SubscribeParam) (err error){
	client.Unsubscribe(param)
	return
}


//获取服务信息
func Nacos_GetAllService(client naming_client.INamingClient, param vo.GetAllServiceInfoParam, err error) {
	service, err := client.GetAllServicesInfo(param)
	fmt.Printf("GetAllService,param:%+v, result:%+v \n\n", param, service)
}
