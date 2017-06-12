package state

import (
	"github.com/synw/microb-dashboard/conf"
	"github.com/synw/microb-dashboard/httpServer"
	"github.com/synw/microb-dashboard/types"
	"github.com/synw/terr"
	"net/http"
)

var DashboardServer = &types.DashboardServer{}
var Conf *types.Conf

func InitState(dev bool, verbosity int) *terr.Trace {
	Conf, tr := conf.GetConf(dev)
	if tr != nil {
		return tr
	}
	instance := &http.Server{}
	running := false
	DashboardServer = &types.DashboardServer{Conf.Domain, Conf.Addr, instance, running}
	httpServer.InitHttpServer(DashboardServer, false)
	return nil
}
