package config

import (
	"github.com/RalapZ/DeepBluePaas/common/middleware"
	"github.com/RalapZ/DeepBluePaas/server/errmsg"
)

//type JWT struct {
//	SigningKey  string `mapstructure:"signing-key" json:"signingKey" yaml:"signing-key"`    // jwt签名
//	ExpiresTime int64  `mapstructure:"expires-time" json:"expiresTime" yaml:"expires-time"` // 过期时间
//	BufferTime  int64  `mapstructure:"buffer-time" json:"bufferTime" yaml:"buffer-time"`    // 缓冲时间
//	Issuer      string `mapstructure:"issuer" json:"issuer" yaml:"issuer"`                  // 签发者
//}


func JWTGetConfig()(errorcode int){
	middleware.JwtConfig.ExpiresTime = Config.GetInt64("auth.jwt.expires-time")
	middleware.JwtConfig.BufferTime = Config.GetInt64("auth.jwt.buffer-time")
	middleware.JwtConfig.Issuer = Config.GetString("auth.jwt.signing-key")
	return errmsg.GLOBAL_STATS_OK
}
