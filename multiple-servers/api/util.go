package api

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// The special type interface{} allows us to take _any_ value, not just one of a specific type.
// This means we can re-use this function for both a single image _and_ a slice of multiple images.
func MarshalWithIndent(data interface{}, indent string) ([]byte, error) {
	// Convert images to a byte-array for writing back in a response
	var b []byte
	var marshalErr error
	// Allow up to 10 characters of indent
	if i, err := strconv.Atoi(indent); err == nil && i > 0 && i <= 10 {
		b, marshalErr = json.MarshalIndent(data, "", strings.Repeat(" ", i))
	} else {
		b, marshalErr = json.Marshal(data)
	}
	if marshalErr != nil {
		return nil, fmt.Errorf("could not marshal data: [%w]", marshalErr)
	}
	return b, nil
}
