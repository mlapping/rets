// xmlutils
package results

import (
	"encoding/xml"
	"io"
)

func ConvertServerResponse(results interface {
	io.Reader
	io.Closer
}, object interface{}) error {
	defer results.Close()
	xmlDecoder := xml.NewDecoder(results)

	return xmlDecoder.Decode(object)
}
