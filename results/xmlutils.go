// xmlutils
package results

import (
	"encoding/xml"
	"io"
)

type RetsReply struct {
	ReplyCode int          `xml:"ReplyCode,attr"`
	ReplyText string       `xml:"ReplyText,attr"`
	Response  RetsResponse `xml:"RETS-RESPONSE"`
}

type RetsResponse struct {
	Text string `xml:">text`
}

func ConvertServerResponse(results interface {
	io.Reader
	io.Closer
}) (*RetsReply, error) {
	defer results.Close()
	xmlDecoder := xml.NewDecoder(results)

	reply := &RetsReply{}
	err := xmlDecoder.Decode(reply)

	if err != nil {
		return nil, err
	}

	return reply, nil
}
