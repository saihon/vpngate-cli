package vpngate

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/lestrrat-go/pdebug"
	"github.com/peco/peco"

	"github.com/saihon/vpngate-cli/record"
)

const (
	// Index number delimiter
	Delim = ": "
)

// "HostName":                      0,
// "IP":                            1,
// "Score":                         2,
// "Ping":                          3,
// "Speed":                         4,
// "CountryLong":                   5,
// "CountryShort":                  6,
// "NumVPNGateSessions":            7,
// "Uptime":                        8,
// "TotalUsers":                    9,
// "TotalTraffic":                  10,
// "LogType":                       11,
// "Operator":                      12,

func writeToBuffer(con *Container, buff *bytes.Buffer) {
	for i := 0; i < con.Len(); i++ {
		rec := con.Get(i)
		proto := ""
		p, ok := rec.Config("proto")
		if ok {
			proto = p[0]
		}

		buff.WriteString(fmt.Sprintf("%4d%s[ %-17s %s %s - %-28s ]\n",
			i, Delim,
			rec.Get(record.N_HostName),
			proto,
			rec.Get(record.N_CountryShort),
			rec.Get(record.N_CountryLong)))
	}
}

// selection by peco
func (v *VPNGate) selection() (*record.Record, error) {
	if pdebug.Enabled {
		pdebug.DefaultCtx.Writer = os.Stderr
	}
	if envvar := os.Getenv("GOMAXPROCS"); envvar == "" {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cli := peco.New()

	var out, in bytes.Buffer
	cli.Stdout = &out

	var argv []string
	if len(v.query) > 0 {
		argv = []string{"--query", v.query}
	}
	cli.Argv = argv

	writeToBuffer(v.container, &in)
	cli.Stdin = &in

	// I do not know how this is good
	if err := cli.Run(ctx); err != nil {
		if err.Error() == "collect results" {
			cli.PrintResults()
			v.query = cli.Query().String()
		}
	}

	if out.Len() == 0 {
		return nil, nil
	}

	// Cut out the first line of the selected string,
	// then split into index number and other
	index, err := strconv.Atoi(
		strings.Split(strings.Split(strings.TrimSpace(out.String()), "\n")[0], Delim)[0])

	return v.container.Get(index), err
}
