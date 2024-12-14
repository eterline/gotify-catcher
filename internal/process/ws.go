package process

import "fmt"

type TargetServer struct {
	prot string
	host string
	token string

	_ struct {}
}

func SetServerWS(host, token string) *TargetServer {
	return &TargetServer {
		prot: "ws",
		host: host,
		token: token,
	}
}

func SetServerWSS(host, token string) *TargetServer {
	return &TargetServer {
		prot: "wss",
		host: host,
		token: token,
	}
}

func (srv *TargetServer) ToString() string {
	return fmt.Sprintf("%s://%s/stream", srv.prot, srv.host)
}