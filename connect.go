package vpngate

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"

	"github.com/saihon/vpngate-cli/record"
)

const (
	OPENVPN = "openvpn"
	BINSH   = "/bin/sh"
)

// connect execute openvpn command and terminate current process
func (v *VPNGate) connect(rec *record.Record) error {
	file, err := v.dump(rec)
	if err != nil {
		return fmt.Errorf("Write .ovpn file: %v", err)
	}

	if _, err := exec.LookPath(OPENVPN); err != nil {
		return fmt.Errorf("Look Path: %v", err)
	}

	command := fmt.Sprintf("sudo %s %s --config %s", OPENVPN, v.flags.option, file)
	return syscall.Exec(BINSH, []string{BINSH, "-c", command}, os.Environ())
}
