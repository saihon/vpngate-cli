package record

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

const (
	// index number of record
	N_HostName                      = 0
	N_IP                            = 1
	N_Score                         = 2
	N_Ping                          = 3
	N_Speed                         = 4
	N_CountryLong                   = 5
	N_CountryShort                  = 6
	N_NumVPNGateSessions            = 7
	N_Uptime                        = 8
	N_TotalUsers                    = 9
	N_TotalTraffic                  = 10
	N_LogType                       = 11
	N_Operator                      = 12
	N_Message                       = 13
	N_OpenVPNGate_ConfigData_Base64 = 14
)

var (
	reSpace = regexp.MustCompile(`\s+`)
)

var (
	ConfigKeys = []string{
		"dev",
		"proto",
		"remote",
		"resolv-retry",
		"nobind",
		"persist-key",
		"persist-tun",
		"cipher",
		"auth",
		"verb",
	}
)

type Record struct {
	record []string
	config map[string][]string
}

// Len
func (r Record) Len() int {
	return len(r.record)
}

// Get
func (r Record) Get(index int) string {
	return r.record[index]
}

// Config get the value of given key. See var ConfigKeys for a list of valid keys
func (r Record) Config(key string) ([]string, bool) {
	v, ok := r.config[key]
	return v, ok
}

// Data return the decoded data from base64 to be the contents of the .ovpn file.
func (r Record) Data() []byte {
	data, _ := decodeBase64([]byte(r.record[N_OpenVPNGate_ConfigData_Base64]))
	return data
}

// Record return the items of CSV in slice
func (r Record) Record() []string {
	return r.record
}

// FileName return the file name create based on config data
func (r Record) FileName() string {
	hostname := "unknown"

	if v := r.record[N_HostName]; len(v) > 0 {
		hostname = v
	}

	filename := "vpngate_" + hostname + ".opengw.net"

	if v, ok := r.config["proto"]; ok {
		if len(v) > 0 && len(v[0]) > 0 {
			filename += "_" + v[0]
		}
	}

	if v, ok := r.config["remote"]; ok {
		if len(v) > 1 && len(v[1]) > 0 {
			filename += "_" + v[1]
		}
	}

	return filename + ".ovpn"
}

// New
func New(row []string) (*Record, error) {
	if len(row) < N_OpenVPNGate_ConfigData_Base64+1 {
		return nil, errors.New("insufficient csv row")
	}

	r := &Record{
		record: row,
		config: make(map[string][]string, len(ConfigKeys)),
	}

	for _, k := range ConfigKeys {
		r.config[k] = nil
	}

	if err := r.mapping(); err != nil {
		return nil, fmt.Errorf("mapping: %v", err)
	}

	return r, nil
}

func decodeBase64(src []byte) ([]byte, error) {
	data := make([]byte, base64.StdEncoding.DecodedLen(len(src)))
	n, err := base64.StdEncoding.Decode(data, []byte(src))
	if err != nil {
		return nil, err
	}

	return data[:n], nil
}

func (r *Record) mapping() error {
	data, err := decodeBase64([]byte(r.record[N_OpenVPNGate_ConfigData_Base64]))
	if err != nil {
		return fmt.Errorf("decodeBase64: %v", err)
	}

	s := bufio.NewScanner(bytes.NewReader(data))

	// <ca></ca>
	// <cert></cert>
	// <key></key>
	for i := 0; s.Scan(); i++ {
		t := strings.TrimSpace(s.Text())
		if t == "" || strings.HasPrefix(t, "#") {
			continue
		}

		for _, k := range ConfigKeys {
			if strings.HasPrefix(t, k) {
				a := reSpace.Split(t, -1)
				if len(a) > 1 {
					r.config[k] = a[1:]
				}
			}
		}
	}

	return s.Err()
}
