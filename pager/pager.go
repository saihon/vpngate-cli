package pager

import (
	"errors"

	"github.com/mattn/go-runewidth"
	termbox "github.com/nsf/termbox-go"
)

// printLineFill
func printLineFill(x, y, w int, s string, fg, bg termbox.Attribute) {
	cell := termbox.Cell{Ch: ' ', Fg: fg, Bg: bg}
	lineFill(0, y, x, cell)
	for _, c := range s {
		termbox.SetCell(x, y, c, fg, bg)
		x += runewidth.RuneWidth(c)
	}
	lineFill(x, y, w, cell)
}

func lineFill(x, y, w int, cell termbox.Cell) {
	for lx := x; lx < w; lx++ {
		termbox.SetCell(lx, y, cell.Ch, cell.Fg, cell.Bg)
	}
}

type Line struct {
	Text string
	Fg   termbox.Attribute
	Bg   termbox.Attribute
	X    int
}

type Pager struct {
	Head       []Line
	Tail       []Line
	Lines      []Line
	Index      int
	pressedKey int
}

func (p *Pager) lastIndex() int {
	_, h := termbox.Size()
	l := len(p.Lines)
	if l-h < 0 {
		return 0
	}
	return l - h
}

// moveToTop move to begin
func (p *Pager) moveToTop() {
	p.Index = 0
	p.draw()
}

// moveToBottom move to end
func (p *Pager) moveToBottom() {
	p.Index = p.lastIndex() + len(p.Head) + len(p.Tail)
	p.draw()
}

// scrollDown scroll down one line
func (p *Pager) scrollDown() {
	_, h := termbox.Size()
	if p.Index+h-len(p.Head)-len(p.Tail) < len(p.Lines) {
		p.Index++
		p.draw()
	}
}

// scrollUp scroll up one line
func (p *Pager) scrollUp() {
	if p.Index > 0 {
		p.Index--
		p.draw()
	}
}

func (p *Pager) draw() {
	w, h := termbox.Size()

	for y := 0; y < h && y < len(p.Head); y++ {
		v := p.Head[y]
		printLineFill(v.X, y, w, v.Text, v.Fg, v.Bg)
	}

	i := p.Index
	for y := len(p.Head); y < h-len(p.Tail) && i < len(p.Lines); y, i = y+1, i+1 {
		v := p.Lines[i]
		printLineFill(v.X, y, w, v.Text, v.Fg, v.Bg)
	}

	i = 0
	for y := h - len(p.Tail); y < h && i < len(p.Tail); y, i = y+1, i+1 {
		v := p.Tail[i]
		printLineFill(v.X, y, w, v.Text, v.Fg, v.Bg)
	}

	termbox.Flush()
}

const (
	KEY_Q = iota
	KEY_E
	KEY_B
	KEY_ENTER
)

var (
	errBreak = errors.New("break")
)

// eventListener
func (p *Pager) eventListener() error {
	switch ev := termbox.PollEvent(); ev.Type {
	case termbox.EventResize:
		p.draw()
	case termbox.EventMouse:
		switch ev.Key {
		case termbox.MouseWheelDown:
			p.scrollDown()
			p.draw()
		case termbox.MouseWheelUp:
			p.scrollUp()
			p.draw()
		}
	case termbox.EventKey:
		switch ev.Key {
		case termbox.KeyEsc, termbox.KeyCtrlC:
			p.pressedKey = KEY_Q
			return errBreak
		case termbox.KeyArrowDown, termbox.KeyCtrlN:
			p.scrollDown()
			p.draw()
		case termbox.KeyArrowUp, termbox.KeyCtrlP:
			p.scrollUp()
			p.draw()
		case termbox.KeyHome:
			p.moveToTop()
			p.draw()
		case termbox.KeyEnd:
			p.moveToBottom()
			p.draw()
		case termbox.KeyEnter: // connect
			p.pressedKey = KEY_ENTER
			return errBreak
		default:
			switch ev.Ch {
			case 'j':
				p.scrollDown()
				p.draw()
			case 'k':
				p.scrollUp()
				p.draw()
			case 'g':
				p.moveToTop()
				p.draw()
			case 'G':
				p.moveToBottom()
				p.draw()
			case 'q': // quit
				p.pressedKey = KEY_Q
				return errBreak
			case 'e': // edit
				p.pressedKey = KEY_E
				return errBreak
			case 'b': // selection, go back
				p.pressedKey = KEY_B
				return errBreak
			}
		}
	case termbox.EventError:
		return ev.Err
	}
	return nil
}

// Start
func Start(head, tail, lines []Line) (int, error) {
	p := &Pager{
		Head:  head,
		Tail:  tail,
		Lines: lines,
		Index: 0,
	}

	if err := termbox.Init(); err != nil {
		return 0, err
	}
	defer termbox.Close()

	// Run before termbox.Close
	defer termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	// termbox.SetInputMode(termbox.InputEsc)
	termbox.HideCursor()

	p.draw()

	for {
		if err := p.eventListener(); err != nil {
			if err == errBreak {
				return p.pressedKey, nil
			}
			return 0, err
		}
	}
}
