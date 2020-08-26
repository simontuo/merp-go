package jwt

import (
	"github.com/simontuo/merp-go/config"
	"time"
)

func TokenExpiresAt() int64 {
	return time.Now().Add(time.Minute * 120).Unix()
}

func TokenKey() []byte {

	return []byte(config.G_Token.Key)
}
