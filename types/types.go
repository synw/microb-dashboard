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
}

type Page struct {
	Url     string
	Title   string
	Content string
}
