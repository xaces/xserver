module xserver

go 1.16

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.7.2
	github.com/gorilla/websocket v1.4.2
	github.com/json-iterator/go v1.1.11
	github.com/kardianos/service v1.2.0
	github.com/mojocn/base64Captcha v1.3.5
	github.com/nats-io/nuid v1.0.1
	github.com/satori/go.uuid v1.2.0
	github.com/unrolled/secure v1.0.9
	github.com/wlgd/xutils v0.0.0-00010101000000-000000000000
)

replace github.com/wlgd/xutils => ../xutils
