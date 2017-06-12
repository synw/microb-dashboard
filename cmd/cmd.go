package cmd

import (
	"github.com/synw/microb-dashboard/state"
	"github.com/synw/microb-dashboard/state/mutate"
	"github.com/synw/microb-dashboard/types"
	"github.com/synw/microb/libmicrob/datatypes"
	"github.com/synw/terr"
)

func Dispatch(cmd *datatypes.Command) *datatypes.Command {
	com := &datatypes.Command{}
	if cmd.Name == "start" {
		res := Start(cmd, state.DashboardServer)
		return res
	} else if cmd.Name == "stop" {
		return Stop(cmd, state.DashboardServer)
	}
	return com
}

func Start(cmd *datatypes.Command, server *types.DashboardServer) *datatypes.Command {
	tr := mutate.StartHttpServer(server)
	if tr != nil {
		cmd.Trace = tr
		cmd.Status = "error"
		terr.Debug("cmd err", tr)
		return cmd
	}
	var resp []interface{}
	resp = append(resp, "Dashboard server started")
	cmd.Status = "success"
	cmd.ReturnValues = resp
	return cmd
}

func Stop(cmd *datatypes.Command, server *types.DashboardServer) *datatypes.Command {
	tr := mutate.StopHttpServer(server)
	if tr != nil {
		cmd.Trace = tr
		cmd.Status = "error"
		return cmd
	}
	var resp []interface{}
	resp = append(resp, "Dashboard server stopped")
	cmd.Status = "success"
	cmd.ReturnValues = resp
	return cmd
}
