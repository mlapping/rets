// compactdecodedutils
package results

import (
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
)

var DecodeData = strings.Split

func (reply *SearchReply) DecodeCompactMultiColumn() ([]map[string]string, error) {
	keys, delimiter, err := reply.DecodeColumns()
	if err != nil {
		return nil, fmt.Errorf("Decoding Server Response failed. Actual RETS reply msg was %v. Error: %v",
			reply.Text, err)
	}

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

func (reply *SearchReply) DecodeColumns() ([]string, string, error) {
	delimiter, err := reply.GetDelimiter()
	if err != nil {
		return nil, "", err
	}
	return DecodeData(reply.Columns, delimiter), delimiter, nil
}

func (reply *SearchReply) GetDelimiter() (string, error) {
	byteSlice, err := hex.DecodeString(reply.Delimiter.Value)
	if err != nil {
		return "", err
	} else if len(byteSlice) == 0 {
		return "", errors.New("Delimiter that the RETS Server returned is not valid")
	}
	return string(byteSlice[0]), nil
}
