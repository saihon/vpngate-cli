package fetch

import (
	"context"
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type Fetch struct {
	Headers map[string]string
	Client  *http.Client
	Timeout time.Duration
	Cache   *Cache
}

// New name is path of cache file
func New(name string, disabled bool) (*Fetch, error) {
	var err error

	f := &Fetch{
		Headers: map[string]string{
			"User-Agent": "vpngate-cli",
		},
		Timeout: 120 * time.Second,
	}

	f.Client = &http.Client{
		Transport: &Progress{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
			output: os.Stderr,
		},
	}

	f.Cache, err = newCache(name, disabled)
	return f, err
}

// Do send request
func (f *Fetch) Do(rawurl string, force bool) ([]byte, error) {
	if !force && f.Cache.cached {
		return f.Cache.Data.CSV, nil
	}

	req, err := http.NewRequest(`GET`, rawurl, nil)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), f.Timeout)
	defer cancel()
	req = req.WithContext(ctx)

	for k, v := range f.Headers {
		req.Header.Set(k, v)
	}

	validatable := f.Cache.SetValidationHeader(req)

	resp, err := f.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	switch resp.StatusCode {
	case http.StatusOK:
		return body, f.Cache.Parse(resp, body)
	case http.StatusNotModified:
		if validatable {
			return f.Cache.Data.CSV, nil
		}
		fallthrough
	default:
		return body, nil
	}
}
