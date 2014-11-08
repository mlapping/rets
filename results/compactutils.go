// compactdecodedutils
package results

import (
	"encoding/hex"
	"errors"
	"strings"
)

var DecodeData = strings.Split

func (reply *SearchReply) DecodeCompactMultiColumn() ([]map[string]string, error) {
	keys, delimiter := reply.DecodeColumns()

	// iterate through each of the data elements
	propertyData := make([]map[string]string, len(reply.Data))
	for i, dataString := range reply.Data {
		values := DecodeData(dataString, delimiter)

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

func (reply *SearchReply) DecodeColumns() ([]string, string) {
	delimiter := reply.GetDelimiter()
	return DecodeData(reply.Columns, delimiter), delimiter
}

func (reply *SearchReply) GetDelimiter() string {
	byteSlice, err := hex.DecodeString(reply.Delimiter.Value)
	if err != nil {
		return ""
	}
	return string(byteSlice[0])
}
