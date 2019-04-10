package vpngate

import (
	"os"

	"github.com/saihon/pathmap"
)

func mapping(directory string) *pathmap.Map {
	m := pathmap.New()
	a, _ := pathmap.Abs(directory)
	root := m.Add(a)
	m.Join(root, `F@CACHE`, `cache.json`)
	m.Join(root, `D@OVPN`, `ovpn`)
	m.Lock()
	if err := root.MkdirIfNotExist(0775); err != nil {
		panic(err)
	}
	return m
}

// Clean clean cache and ovpn files
func Clean(m *pathmap.Map) error {
	if err := m.Get(`F@CACHE`).Remove(); err != nil {
		if !os.IsNotExist(err) {
			return err
		}
	}
	if err := m.Get(`D@OVPN`).RemoveAll(); err != nil {
		if !os.IsNotExist(err) {
			return err
		}
	}
	return nil
}
