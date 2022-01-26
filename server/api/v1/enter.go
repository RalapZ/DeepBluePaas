package v1

import (
	"github.com/RalapZ/DeepBluePaas/server/api/v1/router"
)


type ApiGroup struct{
	router.DP_AUTH
	router.DP_CASBIN

}
