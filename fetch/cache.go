package fetch

import (
	"encoding/json"
	"net/http"
	"os"
)

type Data struct {
	LastModified string `json:"last-modified"`
	ETag         string `json:"etag"`
	CSV          []byte `json:"csv"`
}

type Cache struct {
	Data     Data
	name     string // cache file path
	cached   bool
	disabled bool
}

func newCache(name string, disabled bool) (*Cache, error) {
	c := &Cache{
		name:     name,
		disabled: disabled,
	}

	if disabled {
		return c, nil
	}

	return c, c.read()
}

func (c *Cache) read() error {
	fp, err := os.Open(c.name)
	if err != nil {
		c.cached = false
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	defer fp.Close()

	dec := json.NewDecoder(fp)
	if err := dec.Decode(&c.Data); err != nil {
		c.cached = false
		return err
	}

	c.cached = true
	return nil
}

func (c *Cache) write() error {
	fp, err := os.OpenFile(c.name, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0664)
	if err != nil {
		return err
	}
	defer fp.Close()

	enc := json.NewEncoder(fp)
	return enc.Encode(c.Data)
}

// Parse
func (c *Cache) Parse(resp *http.Response, body []byte) error {
	if c.disabled {
		return nil
	}

	c.Data.LastModified = resp.Header.Get("Last-Modified")
	if len(c.Data.LastModified) == 0 {
		c.Data.LastModified = resp.Header.Get("Date")
	}

	// can't get with "resp.Header.Get"
	if etag, ok := resp.Header["ETag"]; ok && len(etag) > 0 {
		c.Data.ETag = etag[0]
	}

	c.Data.CSV = body

	return c.write()
}

// SetValidationHeader sets If-MKodified-Since and If-Not-Match to request header
func (c *Cache) SetValidationHeader(req *http.Request) bool {
	ok := false
	if c.disabled || !c.cached {
		return ok
	}

	if len(c.Data.LastModified) > 0 {
		req.Header.Set("If-Modified-Since", c.Data.LastModified)
		ok = true
	}
	if len(c.Data.ETag) > 0 {
		req.Header.Set("If-Not-Match", c.Data.ETag)
		ok = true
	}
	return ok
}
