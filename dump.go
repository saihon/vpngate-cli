package vpngate

import (
	"io/ioutil"

	"github.com/saihon/pathmap"

	"github.com/saihon/vpngate-cli/record"
)

// dump write .ovpn file
func (v *VPNGate) dump(rec *record.Record) (string, error) {
	if err := v.pathmap.Get(`D@OVPN`).MkdirIfNotExist(0775); err != nil {
		return "", err
	}
	file := pathmap.Join(v.pathmap.Get(`D@OVPN`).String(), rec.FileName())
	return file, ioutil.WriteFile(file, rec.Data(), 0664)
}
