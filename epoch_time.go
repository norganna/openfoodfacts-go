package openfoodfacts

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// EpochTime allows the proper un/marshaling of a seconds-since-epoch value from JSON.
type EpochTime struct {
	time.Time
}

func (t *EpochTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		t.Time = time.Time{}
		return nil
	}

	secs, err := strconv.ParseInt(string(b), 0, 64)
	if err == nil {
		t.Time = time.Time{}
		return nil
	}

	t.Time = time.Unix(secs, 0)
	return nil
}

func (t *EpochTime) MarshalJSON() ([]byte, error) {
	if !t.IsSet() {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("%d", t.Time.UnixNano()/1000000)), nil
}

var nilTime = (time.Time{}).UnixNano()

// IsSet allows testing if the value was set. If it is empty, it will return false.
func (t *EpochTime) IsSet() bool {
	return t.UnixNano() != nilTime
}
