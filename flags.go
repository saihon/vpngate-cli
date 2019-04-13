package vpngate

import (
	"fmt"
	"os"
	"strings"

	flag "github.com/saihon/flags"
)

type Flags struct {
	request bool
	clean   bool
	option  string
}

var (
	flags Flags
)

func init() {
	flag.CommandLine.Init(os.Args[0], flag.ExitOnError, false)

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(),
			"\nUsage: %s [options] [arguments]\n\n", flag.CommandLine.Name())

		flag.VisitAll(func(f *flag.Flag) {
			s := ""
			if f.Alias > 0 {
				s = fmt.Sprintf("  -%c, --%s", f.Alias, f.Name)
			} else {
				s = fmt.Sprintf("  --%s", f.Name)
			}
			_, usage := flag.UnquoteUsage(f)
			if len(s) <= 4 { // space, space, '-', 'x'.
				s += "\t"
			} else {
				s += "\n    \t"
			}
			s += strings.ReplaceAll(usage, "\n", "\n    \t")
			fmt.Fprint(flag.CommandLine.Output(), s, "\n")
		})
	}

	flag.BoolVar(&flags.request, "request", 'r', false, "Force HTTP request\n")
	flag.BoolVar(&flags.clean, "clean", 'c', false, "Remove cache and .ovpn files\n")
	flag.StringVar(&flags.option, "option", 'o', "--auth-nocache", "Specify the options for openvpn command\n")
	flag.Version("version", 'v', VERSION, "Output version information and exit\n")
}
