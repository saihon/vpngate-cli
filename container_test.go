package vpngate

import (
	"testing"
)

func TestTrimPrefix(t *testing.T) {
	expect := "HostName"
	b := []byte(`*vpn_servers\n#` + expect)

	actual := string(trimPrefix(b))
	if actual != expect {
		t.Errorf("\ngot : %s, want: %s\n", actual, expect)
	}
}
