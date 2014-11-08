// compactdecodedutils
package results

import (
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
)

func (reply *SearchReply) CompactDecoded() ([]map[string]string, error) {
	if reply == nil {
		return nil, fmt.Errorf("rets/CompactDecoded: Passed a nil pointer")
	}

	slice, err := hex.DecodeString(reply.Delimiter.Value)
	if err != nil {
		return nil, errors.New("Error reading Compact Decoded delimiter")
	}

	delimiter := string(slice[0])
	keys := strings.Split(reply.Columns, delimiter)

	// iterate through each of the data elements
	propertyData := make([]map[string]string, len(reply.Data))
	for i, dataString := range reply.Data {
		values := strings.Split(dataString, delimiter)

		if len(keys) != len(values) {
			return nil, errors.New("Error decoding Compact Decoded data: column count does not match values count. Is this a compliant RETS server?")
		}

		propertyData[i] = make(map[string]string, len(keys))
		for j, key := range keys {
			if values[j] != "" {
				propertyData[i][key] = values[j]
			}
		}
	}

	return propertyData, nil
}
