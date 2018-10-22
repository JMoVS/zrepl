package client

import (
	"errors"
	"github.com/zrepl/zrepl/cli"
	"github.com/zrepl/zrepl/config"
	"github.com/zrepl/zrepl/daemon"
	"log"
	"os"
)

var pprofArgs struct {
	daemon.PprofServerControlMsg
}

var PprofCmd = &cli.Subcommand{
	Use:   "pprof off | [on TCP_LISTEN_ADDRESS]",
	Short: "start a http server exposing go-tool-compatible profiling endpoints at TCP_LISTEN_ADDRESS",
	Run: func(subcommand *cli.Subcommand, args []string) error {
		if len(args) < 1 {
			goto enargs
		}
		switch args[0] {
		case "on":
			pprofArgs.Run = true
			if len(args) != 2 {
				return errors.New("must specify TCP_LISTEN_ADDRESS as second positional argument")
			}
			pprofArgs.HttpListenAddress = args[1]
		case "off":
			if len(args) != 1 {
				goto enargs
			}
			pprofArgs.Run = false
		}

		RunPProf(subcommand.Config())
		return nil
	enargs:
		return errors.New("invalid number of positional arguments")

	},
}

func RunPProf(conf *config.Config) {
	log := log.New(os.Stderr, "", 0)

	die := func() {
		log.Printf("exiting after error")
		os.Exit(1)
	}

	log.Printf("connecting to zrepl daemon")

	httpc, err := controlHttpClient(conf.Global.Control.SockPath)
	if err != nil {
		log.Printf("error creating http client: %s", err)
		die()
	}
	err = jsonRequestResponse(httpc, daemon.ControlJobEndpointPProf, pprofArgs.PprofServerControlMsg, struct{}{})
	if err != nil {
		log.Printf("error sending control message: %s", err)
		die()
	}
	log.Printf("finished")
}
