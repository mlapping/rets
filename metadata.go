// metadata
package rets

import (
	"fmt"
	"github.com/mlapping/rets/results"
)

// For lookup types that have several values. Essentially, the id is the primary key to an enumerated field
// Query looks something like: Type=METADATA-LOOKUP_TYPE&ID=Property:20130414180426065306000000
func (sess *Session) MetadataTypeLookup(typeId string) (*results.MetadataReply, error) {
	queryString := map[string]string{
		"Type": "METADATA-LOOKUP_TYPE",
		"ID":   fmt.Sprintf("Property:%s", typeId),
	}

	req, _ := sess.newRequest("GET", sess.Capabilities.MetadataUrl(), queryString)
	res, err := sess.HttpClient.Do(req)

	if err != nil {
		return nil, err
	}

	metaReply := &results.MetadataReply{}
	err = results.ConvertServerResponse(res.Body, metaReply)

	return metaReply, err
}
