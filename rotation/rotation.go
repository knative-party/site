package rotation

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"time"
)

type RotationEntry struct {
	Start time.Time
	End   time.Time
	Data  []string
}

type Rotation struct {
	entries  []RotationEntry
	Metadata map[string]string
}

func ReadFile(r io.Reader) (Rotation, error) {
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
			return retval, fmt.Errorf("")
		}
		retval.entries = append(retval.entries, RotationEntry{
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

func (r *Rotation) At(t time.Time) RotationEntry {
	if t.Before(r.entries[0].Start) {
		return RotationEntry{time.Time{}, r.entries[0].Start, []string{"before rotation"}}
	}
	for i, s := range r.entries {
		if i >= len(r.entries) {
			break
		}
		if s.Start.Before(t) && r.entries[i+1].Start.After(t) {
			return s
		}
	}
	entry := r.entries[len(r.entries)]
	return entry
}

func (r *RotationEntry) String() string {
	return fmt.Sprintf("%s-%s: %v", r.Start, r.End, r.Data)
}
