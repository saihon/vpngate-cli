package vpngate

import (
	"errors"
	"fmt"
	"os"
	"os/user"
	"strings"

	flag "github.com/saihon/flags"
	"github.com/saihon/pathmap"

	"github.com/saihon/vpngate-cli/fetch"
	"github.com/saihon/vpngate-cli/pager"
)

const (
	QUIT      = pager.KEY_Q
	CONNECT   = pager.KEY_ENTER
	EDIT      = pager.KEY_E
	SELECTION = pager.KEY_B
)

var (
	Version      string
	ErrIgnorable = errors.New("ignorable error")
)

type VPNGate struct {
	rawurl    string
	pathmap   *pathmap.Map
	container *Container
	fetch     *fetch.Fetch
	flags     Flags

	query string
}

// New
func New(rawurl, directory string) (*VPNGate, error) {
	usr, err := user.Current()
	if err != nil {
		return nil, err
	}

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		if err == flag.ErrHelp {
			return nil, ErrIgnorable
		}
		return nil, err
	}

	v := &VPNGate{
		rawurl:    rawurl,
		pathmap:   mapping(strings.Replace(directory, "~", usr.HomeDir, 1)),
		container: NewContainer(),
		flags:     flags,
	}

	if v.flags.version {
		os.Stderr.WriteString(os.Args[0] + ": " + Version + "\n")
		return nil, ErrIgnorable
	}

	if v.flags.clean {
		return nil, Clean(v.pathmap)
	}

	v.fetch, err = fetch.New(v.pathmap.Get(`F@CACHE`).String(), false)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (v *VPNGate) Run() error {
	body, err := v.fetch.Do(v.rawurl, v.flags.request)
	if err != nil {
		return err
	}
	if err := v.container.Parse(body); err != nil {
		return err
	}

LABEL_SELECTION:
	rec, err := v.selection()
	if err != nil {
		return fmt.Errorf("selection: %v", err)
	}
	if rec == nil {
		return nil
	}

LABEL_VIEW:
	n, err := v.view(rec)
	if err != nil {
		return err
	}

	switch n {
	case CONNECT:
		return v.connect(rec)
	case EDIT:
		if err := v.edit(rec); err != nil {
			return err
		}
		goto LABEL_VIEW
	case SELECTION:
		goto LABEL_SELECTION
	// case QUIT:
	default:
		return nil
	}
}
