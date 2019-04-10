package vpngate

import (
	"os"

	flag "github.com/saihon/flags"
)

type Flags struct {
	version bool
	request bool
	clean   bool
	option  string
}

var (
	flags Flags
)

func init() {
	flag.CommandLine.Init(os.Args[0], flag.ContinueOnError, false)
	flag.BoolVar(&flags.version, "version", 'v', false, "Output version information and exit")
	flag.BoolVar(&flags.request, "request", 'r', false, "Force HTTP request")
	flag.BoolVar(&flags.clean, "clean", 'c', false, "Remove cache and .ovpn files.")
	flag.StringVar(&flags.option, "option", 'o', "--auth-nocache", "Specify the options for openvpn command\n")
}
