package cli

import (
	"github.com/abiosoft/ishell"
	"github.com/synw/microb-cli/libmicrob/cmd/handler"
	command "github.com/synw/microb/libmicrob/cmd"
	"github.com/synw/terr"
)

func Cmds() *ishell.Cmd {
	command := &ishell.Cmd{
		Name: "dashboard",
		Help: "Commands for the Microb dashboard service: start, stop",
		Func: func(ctx *ishell.Context) {
			if len(ctx.Args) == 0 {
				err := terr.Err("A parameter is required: ex: dashboard start")
				ctx.Println(err.Error())
				return
			}
			if ctx.Args[0] == "start" {
				cmd := command.New("start", "dashboard", "cli", "")
				cmd, timeout, tr := handler.SendCmd(cmd, ctx)
				if tr != nil {
					tr = terr.Pass("cmd.cli.Grg", tr)
					msg := tr.Formatc()
					ctx.Println(msg)
				}
				if timeout == true {
					err := terr.Err("Timeout: server is not responding")
					ctx.Println(err.Error())
				}
			} else if ctx.Args[0] == "stop" {
				cmd := command.New("stop", "dashboard", "cli", "")
				cmd, timeout, tr := handler.SendCmd(cmd, ctx)
				if tr != nil {
					tr = terr.Pass("cli.Stop", tr)
					msg := tr.Formatc()
					ctx.Println(msg)
				}
				if timeout == true {
					err := terr.Err("Timeout: server is not responding")
					ctx.Println(err.Error())
				}
			}
			return
		},
	}
	return command
}
