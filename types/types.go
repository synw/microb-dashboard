package types

import (
	"net/http"
)

type DashboardServer struct {
	Domain   string
	Addr     string
	Instance *http.Server
	Running  bool
}

type Conf struct {
	Domain string
	Addr   string
	WsAddr string
	WsKey  string
}

type Conn struct {
	Addr      string
	Timestamp string
	User      string
	Token     string
}

type Page struct {
	Url     string
	Title   string
	Content string
	Conn    *Conn
}
