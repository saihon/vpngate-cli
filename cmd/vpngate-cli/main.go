package main

import (
	"fmt"
	"os"

	vpngate "github.com/saihon/vpngate-cli"
)

const (
	LIST_URL  = "https://www.vpngate.net/api/iphone/"
	CACHE_DIR = "~/.cache/vpngate-cli"
)

var (
	version string
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Fprintf(os.Stderr, "Error: Recover: %v\n", err)
			os.Exit(1)
		}
	}()

	os.Exit(_main())
}

func _main() int {
	vpngate.Version = version

	v, err := vpngate.New(LIST_URL, CACHE_DIR)
	if err != nil {
		if err == vpngate.ErrIgnorable {
			return 2
		}
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return 1
	}

	if err := v.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return 1
	}

	return 0
}
