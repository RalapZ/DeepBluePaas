package dbgrpc

import (
	"fmt"
	"github.com/RalapZ/DeepBluePaas/server/config"
	"testing"
)

//func TestStart2(t *testing.T) {
//	Start()
//}


func TestGrpcRegister(t *testing.T){
	config.ParserConfig()
	fmt.Printf("%##v",config.NacosInsConfig)
	GrpcRegister()
}