package vpngate

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/saihon/vpngate-cli/record"
)

// edit edit the .ovpn file. reading changes is not implemented
func (v *VPNGate) edit(rec *record.Record) error {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "vi"
	}

	file, err := v.dump(rec)
	if err != nil {
		return fmt.Errorf("write ovpn file: %v", err)
	}

	if _, err := exec.LookPath(editor); err != nil {
		return err
	}

	return exec.Command("/bin/sh", "-c", fmt.Sprintf("%s %s", editor, file)).Run()
}
