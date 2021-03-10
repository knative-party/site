// Package rotation defines tools for managing a simple "oncall rotation" file.
// The file format is designed to be easy for both computers and humans to read,
// and is line-oriented, with `#` comments.
//
// There are three types of lines:
// - Empty lines or simple comments, of the form `# ....`
// - Metadata lines, indicated by `#@ key: values`
// - Rotation start lines, indicated as `RFC3339Date | who is oncall`
//
// Rotation start lines must be ordered in the file from oldest to newest; it is
// possible to keep as much rotation history (or future rotation information) as
// desired.
package rotation

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

// Rotation defines a rotation -- a series of events with associated data which
// fill a calendar range (one entry ends when the next one starts, entries are
// in calendar order), and some metadata about the Rotation.
type Rotation struct {
	entries  []Entry
	Metadata map[string]string
}

// Entry describes a single entry in a rotation -- a calendar period
// with a defined start and end time, and some string data.
type Entry struct {
	Start time.Time
	End   time.Time
	Data  []string
}

// FromFile reads a rotation from the specified filename.
func FromFile(name string) (Rotation, error) {
	f, err := os.Open(name)
	defer f.Close()
	if err != nil {
		return Rotation{}, err
	}
	return Read(f)
}

// FromURL reads a rotation fetched by a GET of the specified URL.
func FromURL(url string) (Rotation, error) {
	resp, err := http.Get(url)
	if err != nil {
		return Rotation{}, err
	}
	return Read(resp.Body)
}

// Read reads a rotation from any reader; `FromFile` and `FromURL` are
// helpers for common rotation sources.
func Read(r io.Reader) (Rotation, error) {
	retval := Rotation{
		Metadata: make(map[string]string),
	}
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "#") {
			if strings.HasPrefix(line, "#@") {
				// Key-value metadataline
				kv := strings.SplitN(line[2:], ":", 2)
				if len(kv) == 1 {
					kv = append(kv, "")
				}
				retval.Metadata[strings.TrimSpace(kv[0])] = strings.TrimSpace(kv[1])
			}
			continue
		}
		fields := strings.Fields(line)
		if len(fields) == 0 {
			continue
		}
		start, err := time.Parse(time.RFC3339, fields[0])
		if err != nil {
			return retval, err
		}
		if fields[1] != "|" {
			return retval, fmt.Errorf(`Expected "<DATE> | <DATA>", missing "|"`)
		}
		retval.entries = append(retval.entries, Entry{
			Start: start,
			Data:  fields[2:],
		})
	}
	// Ensure that the Rotation is sorted
	prev := time.Time{}
	for i, entry := range retval.entries {
		if !prev.Before(entry.Start) {
			return retval, fmt.Errorf("Dates out of order at %d: %s >= %s", i, prev, entry.Start)
		}
		if i != 0 { // Close the intervals
			retval.entries[i-1].End = entry.Start
		}
		prev = entry.Start
	}
	retval.entries[len(retval.entries)-1].End = retval.entries[len((retval.entries))-1].Start.Add(365 * 24 * time.Hour)
	return retval, nil
}

// At returns the entry from the rotation which encompases the current time
// (i.e. Start < t < End).
func (r *Rotation) At(t time.Time) Entry {
	for i := range r.entries {
		s := r.entries[len(r.entries)-i-1]
		if s.Start.Before(t) {
			return s
		}
	}
	//	if t.Before(r.entries[0].Start) {
	//}
	//for i, s := range r.entries {
	// 	if i >= len(r.entries) {
	// 		break
	// 	}
	// 	if s.Start.Before(t) && r.entries[i+1].Start.After(t) {
	// 		return s
	// 	}
	// }
	return Entry{time.Time{}, r.entries[0].Start, []string{"before rotation"}}
	//	entry := r.entries[len(r.entries)]
	//	return entry
}

// Next determines the entry which begins *after* the current time.
func (r *Rotation) Next(t time.Time) Entry {
	if len(r.entries) == 0 {
		return Entry{Data: []string{"no entries"}}
	}
	for _, s := range r.entries {
		if s.Start.After(t) {
			return s
		}
	}
	return r.entries[len(r.entries)-1]
}

// String implements the Stringer interface.
func (r *Entry) String() string {
	return fmt.Sprintf("%s-%s: %v", r.Start, r.End, r.Data)
}
