package rotation

import (
	"errors"
	"strings"
	"testing"
	"time"
)

func TestRead(t *testing.T) {
	tests := []struct {
		name string
		file string
	}{{
		name: "simple",
		file: `
		# Comment
		#@ at0: 2021-02-12T00:00:00Z before rotation
		#@ after0: 2021-02-12T00:00:00Z some words
		2021-03-01T01:00:00Z | some words
		#@ at1: 2021-03-01T04:00:00Z some words
		#@ after1: 2021-03-01T04:00:00Z more
		2021-03-02T01:00:00Z | more
		#@ after2: 2021-03-08T01:00:00Z last
		#@ after3: 2021-05-08T01:00:00Z last
		#@ at2: 2021-05-08T00:59:59Z more
		2021-05-08T01:00:00Z | last
		#@ at3: 2021-06-01T00:00:00Z last
		#@ after4: 2021-06-01T00:00:00Z last
		`,
	},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Read(strings.NewReader(tt.file))
			if err != nil {
				t.Errorf("Read(%s) error = %v", tt.file, err)
				return
			}
			for k, v := range got.Metadata {
				fields := strings.Fields(v)
				if strings.HasPrefix(k, "at") {
					date, err := time.Parse(time.RFC3339, fields[0])
					if err != nil {
						t.Errorf("Unable to parse %q as a time: %s", fields[0], err)
						continue
					}
					e := got.At(date)
					want := strings.Join(fields[1:], " ")
					if strings.Join(e.Data, " ") != want {
						t.Errorf("Expected %s at %s, got %s", want, date, e.Data)
					}
				}
				if strings.HasPrefix(k, "after") {
					date, err := time.Parse(time.RFC3339, fields[0])
					if err != nil {
						t.Errorf("Unable to parse %q as a time: %s", fields[0], err)
						continue
					}
					e := got.Next(date)
					want := strings.Join(fields[1:], " ")
					if strings.Join(e.Data, " ") != want {
						t.Errorf("Expected %s at %s, got %s", want, date, e.Data)
					}
				}
			}
		})
	}
}

func TestReadErrors(t *testing.T) {
	tests := []struct {
		name string
		file string
		err  error
	}{{
		name: "badtime",
		file: "March 20, 2011 | stuff",
		err:  errors.New(`parsing time "March" as "2006-01-02T15:04:05Z07:00": cannot parse "March" as "2006"`),
	}, {
		name: "outoforder",
		file: `2021-03-11T01:00:00Z | okey
			2021-03-21T01:00:00Z | dokey
			2021-03-31T01:00:00Z | arti
			2021-03-10T01:00:00Z | chokey`,
		err: errors.New(`Dates out of order at 3: 2021-03-31 01:00:00 +0000 UTC >= 2021-03-10 01:00:00 +0000 UTC`),
	}, {
		name: "nopipe",
		file: "2021-03-11T01:00:00Z oops, i did it again",
		err:  errors.New(`Expected "<DATE> | <DATA>", missing "|"`),
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Read(strings.NewReader(tt.file))
			if err == nil {
				t.Errorf("Expected %q to have an error (%s), got no error", tt.file, tt.err)
			}
			if err.Error() != tt.err.Error() {
				t.Errorf("Expected %q from %q, got %q", tt.err, tt.file, err)
			}
		})
	}
}
