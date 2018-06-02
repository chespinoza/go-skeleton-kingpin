package server

import (
	"log"
	"net"
	"net/http"
	"net/rpc"

	"go.uber.org/zap"
)

type CliSrv struct {
	log *zap.Logger
}

// Commands

// ReloadSomething command
func (cli *CliSrv) ReloadSomething(param *string, reply *string) error {
	cli.log.Info("Reload Something", zap.String("param", *param))
	*reply = "Reloaded!"
	return nil
}

// DoSomething command
func (cli *CliSrv) DoSomething(result *bool, reply *bool) error {
	cli.log.Info("Do Something")
	*result = true
	return nil
}

// StartRemoteCli starts remote CLI service
func StartRemoteCli(logger *zap.Logger) {
	go func() {
		cli := new(CliSrv)
		cli.log = logger
		rpc.Register(cli)
		rpc.HandleHTTP()
		l, err := net.Listen("tcp", ":1234")
		if err != nil {
			log.Fatal("Cli Listening error:", err)
		}
		if err := http.Serve(l, nil); err != nil {
			log.Fatal(err)
		}
	}()
}
