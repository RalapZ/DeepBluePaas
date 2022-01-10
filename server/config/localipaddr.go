package config

import (
	"github.com/RalapZ/DeepBluePaas/server/errmsg"
	"net"
	"strings"
)

func getLocalIpaddr()(errorcode int){
	conn, err := net.Dial("udp", "8.8.8.8:80")
	defer conn.Close()
	if err != nil {
		return
	}
	LocalAddr = strings.Split(conn.LocalAddr().String(),":")[0]

	return  errmsg.GLOBAL_STATS_OK
}
