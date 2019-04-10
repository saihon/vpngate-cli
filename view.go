package vpngate

import (
	"fmt"
	"strings"

	"github.com/nsf/termbox-go"

	"github.com/saihon/vpngate-cli/pager"
	"github.com/saihon/vpngate-cli/record"
)

const (
	coldef = termbox.ColorDefault
)

func (v VPNGate) view(rec *record.Record) (int, error) {
	lines := make([]pager.Line, 0, 24)

	keys := v.container.Keys()
	n := 1
	for i := 0; i < rec.Len()-1; i++ {
		bg := coldef
		if n%2 == 0 {
			bg = termbox.ColorBlack
		}
		lines = append(lines, pager.Line{Bg: bg, Text: fmt.Sprintf("%s: %s", keys[i], rec.Get(i))})
		n++
	}

	for _, key := range record.ConfigKeys {
		v, ok := rec.Config(key)
		if ok {
			bg := coldef
			if n%2 == 0 {
				bg = termbox.ColorBlack
			}
			lines = append(lines, pager.Line{Bg: bg, Text: fmt.Sprintf("%s: %s", key, strings.Join(v, " "))})
			n++
		}
	}

	head := []pager.Line{
		{
			Text: "MENU KEY: enter(connect) | e(edit) | (b)go back | q(quit)",
			Fg:   coldef | termbox.AttrUnderline,
			Bg:   coldef,
		},
	}

	key, err := pager.Start(head, nil, lines)
	if err != nil {
		return QUIT, err
	}

	return key, nil
}
