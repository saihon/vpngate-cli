package record

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	data := [][]string{
		{
			"vpn000000000",
			"111.111.111.111",
			"2222222",
			"33",
			"44444444",
			"United States",
			"US",
			"55",
			"666666666",
			"777777",
			"88888888888888",
			"2weeks",
			"owner",
			"",
			"",
		},
	}

	for i, row := range data {
		r, err := New(row)
		if err != nil {
			t.Errorf("%d: error: %v", i, err)
			return
		}

		for j := 0; j < len(r.record) && j < len(row); j++ {
			expect := row[j]
			actual := r.record[j]
			if actual != expect {
				t.Errorf("%d-%d:\ngot : %s, want: %s\n", i, j, actual, expect)
				return
			}
		}
	}
}

func TestMapping(t *testing.T) {
	data := "IyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIwojIE9wZW5WUE4gMi4wIFNhbXBsZSBDb25maWd1cmF0aW9uIEZpbGUKIyBmb3IgUGFja2V0aVggVlBOIC8gU29mdEV0aGVyIFZQTiBTZXJ2ZXIKIwojICEhISBBVVRPLUdFTkVSQVRFRCBCWSBTT0ZURVRIRVIgVlBOIFNFUlZFUiBNQU5BR0VNRU5UIFRPT0wgISEhCiMKIyAhISEgWU9VIEhBVkUgVE8gUkVWSUVXIElUIEJFRk9SRSBVU0UgQU5EIE1PRElGWSBJVCBBUyBORUNFU1NBUlkgISEhCiMKIyBUaGlzIGNvbmZpZ3VyYXRpb24gZmlsZSBpcyBhdXRvLWdlbmVyYXRlZC4gWW91IG1pZ2h0IHVzZSB0aGlzIGNvbmZpZyBmaWxlCiMgaW4gb3JkZXIgdG8gY29ubmVjdCB0byB0aGUgUGFja2V0aVggVlBOIC8gU29mdEV0aGVyIFZQTiBTZXJ2ZXIuCiMgSG93ZXZlciwgYmVmb3JlIHlvdSB0cnkgaXQsIHlvdSBzaG91bGQgcmV2aWV3IHRoZSBkZXNjcmlwdGlvbnMgb2YgdGhlIGZpbGUKIyB0byBkZXRlcm1pbmUgdGhlIG5lY2Vzc2l0eSB0byBtb2RpZnkgdG8gc3VpdGFibGUgZm9yIHlvdXIgcmVhbCBlbnZpcm9ubWVudC4KIyBJZiBuZWNlc3NhcnksIHlvdSBoYXZlIHRvIG1vZGlmeSBhIGxpdHRsZSBhZGVxdWF0ZWx5IG9uIHRoZSBmaWxlLgojIEZvciBleGFtcGxlLCB0aGUgSVAgYWRkcmVzcyBvciB0aGUgaG9zdG5hbWUgYXMgYSBkZXN0aW5hdGlvbiBWUE4gU2VydmVyCiMgc2hvdWxkIGJlIGNvbmZpcm1lZC4KIwojIE5vdGUgdGhhdCB0byB1c2UgT3BlblZQTiAyLjAsIHlvdSBoYXZlIHRvIHB1dCB0aGUgY2VydGlmaWNhdGlvbiBmaWxlIG9mCiMgdGhlIGRlc3RpbmF0aW9uIFZQTiBTZXJ2ZXIgb24gdGhlIE9wZW5WUE4gQ2xpZW50IGNvbXB1dGVyIHdoZW4geW91IHVzZSB0aGlzCiMgY29uZmlnIGZpbGUuIFBsZWFzZSByZWZlciB0aGUgYmVsb3cgZGVzY3JpcHRpb25zIGNhcmVmdWxseS4KCgojIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjCiMgU3BlY2lmeSB0aGUgdHlwZSBvZiB0aGUgbGF5ZXIgb2YgdGhlIFZQTiBjb25uZWN0aW9uLgojCiMgVG8gY29ubmVjdCB0byB0aGUgVlBOIFNlcnZlciBhcyBhICJSZW1vdGUtQWNjZXNzIFZQTiBDbGllbnQgUEMiLAojICBzcGVjaWZ5ICdkZXYgdHVuJy4gKExheWVyLTMgSVAgUm91dGluZyBNb2RlKQojCiMgVG8gY29ubmVjdCB0byB0aGUgVlBOIFNlcnZlciBhcyBhIGJyaWRnaW5nIGVxdWlwbWVudCBvZiAiU2l0ZS10by1TaXRlIFZQTiIsCiMgIHNwZWNpZnkgJ2RldiB0YXAnLiAoTGF5ZXItMiBFdGhlcm5ldCBCcmlkZ2luZSBNb2RlKQoKZGV2IHR1bgoKCiMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMKIyBTcGVjaWZ5IHRoZSB1bmRlcmx5aW5nIHByb3RvY29sIGJleW9uZCB0aGUgSW50ZXJuZXQuCiMgTm90ZSB0aGF0IHRoaXMgc2V0dGluZyBtdXN0IGJlIGNvcnJlc3BvbmQgd2l0aCB0aGUgbGlzdGVuaW5nIHNldHRpbmcgb24KIyB0aGUgVlBOIFNlcnZlci4KIwojIFNwZWNpZnkgZWl0aGVyICdwcm90byB0Y3AnIG9yICdwcm90byB1ZHAnLgoKcHJvdG8gdGNwCgoKIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIwojIFRoZSBkZXN0aW5hdGlvbiBob3N0bmFtZSAvIElQIGFkZHJlc3MsIGFuZCBwb3J0IG51bWJlciBvZgojIHRoZSB0YXJnZXQgVlBOIFNlcnZlci4KIwojIFlvdSBoYXZlIHRvIHNwZWNpZnkgYXMgJ3JlbW90ZSA8SE9TVE5BTUU+IDxQT1JUPicuIFlvdSBjYW4gYWxzbwojIHNwZWNpZnkgdGhlIElQIGFkZHJlc3MgaW5zdGVhZCBvZiB0aGUgaG9zdG5hbWUuCiMKIyBOb3RlIHRoYXQgdGhlIGF1dG8tZ2VuZXJhdGVkIGJlbG93IGhvc3RuYW1lIGFyZSBhICJhdXRvLWRldGVjdGVkCiMgSVAgYWRkcmVzcyIgb2YgdGhlIFZQTiBTZXJ2ZXIuIFlvdSBoYXZlIHRvIGNvbmZpcm0gdGhlIGNvcnJlY3RuZXNzCiMgYmVmb3JlaGFuZC4KIwojIFdoZW4geW91IHdhbnQgdG8gY29ubmVjdCB0byB0aGUgVlBOIFNlcnZlciBieSB1c2luZyBUQ1AgcHJvdG9jb2wsCiMgdGhlIHBvcnQgbnVtYmVyIG9mIHRoZSBkZXN0aW5hdGlvbiBUQ1AgcG9ydCBzaG91bGQgYmUgc2FtZSBhcyBvbmUgb2YKIyB0aGUgYXZhaWxhYmxlIFRDUCBsaXN0ZW5lcnMgb24gdGhlIFZQTiBTZXJ2ZXIuCiMKIyBXaGVuIHlvdSB1c2UgVURQIHByb3RvY29sLCB0aGUgcG9ydCBudW1iZXIgbXVzdCBzYW1lIGFzIHRoZSBjb25maWd1cmF0aW9uCiMgc2V0dGluZyBvZiAiT3BlblZQTiBTZXJ2ZXIgQ29tcGF0aWJsZSBGdW5jdGlvbiIgb24gdGhlIFZQTiBTZXJ2ZXIuCgpyZW1vdGUgMDAwLjAwMC4wMDAuMDAwIDAwMDAKCgojIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjCiMgVGhlIEhUVFAvSFRUUFMgcHJveHkgc2V0dGluZy4KIwojIE9ubHkgaWYgeW91IGhhdmUgdG8gdXNlIHRoZSBJbnRlcm5ldCB2aWEgYSBwcm94eSwgdW5jb21tZW50IHRoZSBiZWxvdwojIHR3byBsaW5lcyBhbmQgc3BlY2lmeSB0aGUgcHJveHkgYWRkcmVzcyBhbmQgdGhlIHBvcnQgbnVtYmVyLgojIEluIHRoZSBjYXNlIG9mIHVzaW5nIHByb3h5LWF1dGhlbnRpY2F0aW9uLCByZWZlciB0aGUgT3BlblZQTiBtYW51YWwuCgo7aHR0cC1wcm94eS1yZXRyeQo7aHR0cC1wcm94eSBbcHJveHkgc2VydmVyXSBbcHJveHkgcG9ydF0KCgojIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjCiMgVGhlIGVuY3J5cHRpb24gYW5kIGF1dGhlbnRpY2F0aW9uIGFsZ29yaXRobS4KIwojIERlZmF1bHQgc2V0dGluZyBpcyBnb29kLiBNb2RpZnkgaXQgYXMgeW91IHByZWZlci4KIyBXaGVuIHlvdSBzcGVjaWZ5IGFuIHVuc3VwcG9ydGVkIGFsZ29yaXRobSwgdGhlIGVycm9yIHdpbGwgb2NjdXIuCiMKIyBUaGUgc3VwcG9ydGVkIGFsZ29yaXRobXMgYXJlIGFzIGZvbGxvd3M6CiMgIGNpcGhlcjogW05VTEwtQ0lQSEVSXSBOVUxMIEFFUy0xMjgtQ0JDIEFFUy0xOTItQ0JDIEFFUy0yNTYtQ0JDIEJGLUNCQwojICAgICAgICAgIENBU1QtQ0JDIENBU1Q1LUNCQyBERVMtQ0JDIERFUy1FREUtQ0JDIERFUy1FREUzLUNCQyBERVNYLUNCQwojICAgICAgICAgIFJDMi00MC1DQkMgUkMyLTY0LUNCQyBSQzItQ0JDCiMgIGF1dGg6ICAgU0hBIFNIQTEgTUQ1IE1ENCBSTUQxNjAKCmNpcGhlciBBRVMtMTI4LUNCQwphdXRoIFNIQTEKCgojIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjCiMgT3RoZXIgcGFyYW1ldGVycyBuZWNlc3NhcnkgdG8gY29ubmVjdCB0byB0aGUgVlBOIFNlcnZlci4KIwojIEl0IGlzIG5vdCByZWNvbW1lbmRlZCB0byBtb2RpZnkgaXQgdW5sZXNzIHlvdSBoYXZlIGEgcGFydGljdWxhciBuZWVkLgoKcmVzb2x2LXJldHJ5IGluZmluaXRlCm5vYmluZApwZXJzaXN0LWtleQpwZXJzaXN0LXR1bgpjbGllbnQKdmVyYiAzCiNhdXRoLXVzZXItcGFzcwoKCiMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMKIyBUaGUgY2VydGlmaWNhdGUgZmlsZSBvZiB0aGUgZGVzdGluYXRpb24gVlBOIFNlcnZlci4KIwojIFRoZSBDQSBjZXJ0aWZpY2F0ZSBmaWxlIGlzIGVtYmVkZGVkIGluIHRoZSBpbmxpbmUgZm9ybWF0LgojIFlvdSBjYW4gcmVwbGFjZSB0aGlzIENBIGNvbnRlbnRzIGlmIG5lY2Vzc2FyeS4KIyBQbGVhc2Ugbm90ZSB0aGF0IGlmIHRoZSBzZXJ2ZXIgY2VydGlmaWNhdGUgaXMgbm90IGEgc2VsZi1zaWduZWQsIHlvdSBoYXZlIHRvCiMgc3BlY2lmeSB0aGUgc2lnbmVyJ3Mgcm9vdCBjZXJ0aWZpY2F0ZSAoQ0EpIGhlcmUuCgo8Y2E+Ci0tLS0tQkVHSU4gQ0VSVElGSUNBVEUtLS0tLQotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCgo8L2NhPgoKCiMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMKIyBUaGUgY2xpZW50IGNlcnRpZmljYXRlIGZpbGUgKGR1bW15KS4KIwojIEluIHNvbWUgaW1wbGVtZW50YXRpb25zIG9mIE9wZW5WUE4gQ2xpZW50IHNvZnR3YXJlCiMgKGZvciBleGFtcGxlOiBPcGVuVlBOIENsaWVudCBmb3IgaU9TKSwKIyBhIHBhaXIgb2YgY2xpZW50IGNlcnRpZmljYXRlIGFuZCBwcml2YXRlIGtleSBtdXN0IGJlIGluY2x1ZGVkIG9uIHRoZQojIGNvbmZpZ3VyYXRpb24gZmlsZSBkdWUgdG8gdGhlIGxpbWl0YXRpb24gb2YgdGhlIGNsaWVudC4KIyBTbyB0aGlzIHNhbXBsZSBjb25maWd1cmF0aW9uIGZpbGUgaGFzIGEgZHVtbXkgcGFpciBvZiBjbGllbnQgY2VydGlmaWNhdGUKIyBhbmQgcHJpdmF0ZSBrZXkgYXMgZm9sbG93cy4KCjxjZXJ0PgotLS0tLUJFR0lOIENFUlRJRklDQVRFLS0tLS0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQoKPC9jZXJ0PgoKPGtleT4KLS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQotLS0tLUVORCBSU0EgUFJJVkFURSBLRVktLS0tLQoKPC9rZXk+Cgo="
	length := N_OpenVPNGate_ConfigData_Base64 + 1
	r := Record{
		record: make([]string, length, length),
		config: make(map[string][]string, len(ConfigKeys)),
	}
	r.record[N_OpenVPNGate_ConfigData_Base64] = data

	if err := r.mapping(); err != nil {
		t.Errorf("\nerror: %v\n", err)
	}

	expects := map[string][]string{
		"dev":          {"tun"},
		"proto":        {"tcp"},
		"remote":       {"000.000.000.000", "0000"},
		"resolv-retry": {"infinite"},
		"nobind":       nil,
		"persist-key":  nil,
		"persist-tun":  nil,
		"cipher":       {"AES-128-CBC"},
		"auth":         {"SHA1"},
		"verb":         {"3"},
	}

	for _, key := range ConfigKeys {
		actual, _ := r.Config(key)
		expect := expects[key]
		if !reflect.DeepEqual(actual, expect) {
			t.Errorf("key: %s\n  got : %v\n  want: %v\n", key, actual, expect)
			break
		}
	}
}

func TestFileName(t *testing.T) {
	type testdata struct {
		hostname, expect string
		proto, remote    []string
	}

	data := []testdata{
		{
			hostname: "",
			proto:    []string{},
			remote:   []string{},
			expect:   `vpngate_unknown.opengw.net.ovpn`,
		},
		{
			hostname: "vpn000000000000",
			proto:    []string{},
			remote:   []string{},
			expect:   `vpngate_vpn000000000000.opengw.net.ovpn`,
		},
		{
			hostname: "vpn111111111111",
			proto:    []string{"udp"},
			expect:   `vpngate_vpn111111111111.opengw.net_udp.ovpn`,
		},
		{
			hostname: "vpn222222222222",
			proto:    []string{"tcp"},
			remote:   []string{"", "0000"},
			expect:   `vpngate_vpn222222222222.opengw.net_tcp_0000.ovpn`,
		},
	}

	for i, v := range data {
		r := Record{
			record: make([]string, 15, 15),
			config: make(map[string][]string),
		}
		r.record[N_HostName] = v.hostname
		r.config["proto"] = v.proto
		r.config["remote"] = v.remote

		actual := r.FileName()
		if actual != v.expect {
			t.Errorf("\n%d:\n  got : %s, want: %s\n", i, actual, v.expect)
		}
	}
}
