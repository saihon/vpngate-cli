package fetch

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type Progress struct {
	Transport http.RoundTripper
	output    io.Writer
	donec     chan string
}

func (p *Progress) transport() http.RoundTripper {
	if p.Transport == nil {
		return http.DefaultTransport
	}
	return p.Transport
}

func (p *Progress) RoundTrip(req *http.Request) (*http.Response, error) {
	go p.Print(req.URL.String())
	resp, err := p.transport().RoundTrip(req)

	statuscode := "!?"
	if resp != nil {
		statuscode = resp.Status
	}
	p.donec <- statuscode

	return resp, err
}

// Print
func (p *Progress) Print(rawurl string) {
	i := 0
	a := []string{"-", "\\", "|", "/"}
	p.donec = make(chan string)
	defer close(p.donec)

	for {
		select {
		case v := <-p.donec:
			fmt.Fprintf(p.output, "\n    [HTTP Status ]: %s\033[0m\n", v)
			return
		default:
			fmt.Fprintf(p.output, "\033[2K\r %s\033[0m  [HTTP Request]: %s", a[i], rawurl)
			if i+1 == len(a) {
				i = 0
			} else {
				i++
			}
			time.Sleep(100 * time.Millisecond)
		}
	}
}
