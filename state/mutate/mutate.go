package mutate

import (
	"errors"
	"github.com/synw/microb-dashboard/httpServer"
	"github.com/synw/microb-dashboard/types"
	"github.com/synw/terr"
)

func StartHttpServer(server *types.DashboardServer) *terr.Trace {
	if server.Running == true {
		err := errors.New("Dashboard server is already running")
		tr := terr.New("state.mutate.StartHttpServer", err)
		return tr
	}
	go httpServer.Run(server)
	return nil
}

func StopHttpServer(server *types.DashboardServer) *terr.Trace {
	if server.Running == false {
		err := errors.New("Dashboard server is not running")
		tr := terr.New("state.mutate.StopHttpServer", err)
		return tr
	}
	tr := httpServer.Stop(server)
	if tr != nil {
		return tr
	}
	return nil
}
