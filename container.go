package vpngate

import (
	"bytes"
	"encoding/csv"
	"io"

	"github.com/saihon/vpngate-cli/record"
)

type Container struct {
	keys    []string
	records []*record.Record
}

// NewContainer
func NewContainer() *Container {
	return &Container{}
}

// Append
func (c *Container) Append(row []string) error {
	rec, err := record.New(row)
	if err != nil {
		return err
	}

	c.records = append(c.records, rec)
	return nil
}

// Each
func (c Container) Each(f func(int, *record.Record) bool) {
	for i, r := range c.records {
		if f(i, r) {
			break
		}
	}
}

// Get
func (c Container) Get(index int) *record.Record {
	return c.records[index]
}

// Len records length
func (c *Container) Len() int {
	return len(c.records)
}

// Keys
func (c Container) Keys() []string {
	return c.keys
}

var (
	prefix = []byte("*vpn_servers")
)

func trimPrefix(b []byte) []byte {
	// if begin of b is *vpn_servers, not parse to CSV correctly.
	// so if it is there, skip it
	if bytes.HasPrefix(b, prefix) {
		for i := len(prefix); i < 24 && i < len(b); i++ {
			if b[i] == '#' {
				return b[i+1:]
			}
		}
		return b[len(prefix):]
	}
	return b
}

// Parse
func (c *Container) Parse(b []byte) error {
	r := bytes.NewReader(trimPrefix(b))
	csvr := csv.NewReader(r)
	csvr.LazyQuotes = true
	// csvr.FieldsPerRecord = -1
	// csvr.TrimLeadingSpace = true

	// Use Read because ReadAll can not slice when a syntax error occurs
	for i := 0; ; i++ {
		row, err := csvr.Read() // read one line
		if err == io.EOF {
			break
		}
		if err != nil {
			continue
		}

		if i == 0 {
			c.keys = row
			continue
		}

		rec, err := record.New(row)
		if err != nil {
			continue
		}

		c.records = append(c.records, rec)
	}

	return nil
}
