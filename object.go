// object
package rets

import (
	"fmt"
	"github.com/mlapping/rets/results"
	"strconv"
)

// GetObject?Resource=Property&Type=Thumbnail&ID=20141106154635467418000000:0&Location=0
type ObjectQuery struct {
	Resource string
	Id       string // this is value associated with the KeyField property which is set at the resource-level
	Type     string
	Index    string // [0-9]+|* , * meaning all
	Location int    // 0 for picture download, 1 for link (if you're using flexmls)
}

//Location:http://cdn.photos.flexmls.com/knx/20141106174321256467000000-t.jpg
//MIME-Version:1.0
//Object-ID:1
//Preferred:1
//RETS-Version:RETS/1.5
//Server:Apache-Coyote/1.1
//Set-Cookie:JSESSIONID=CC297D51FEF893664506CE5A20101C6D; Path=/rets2_1

// GetObject?Resource=Property&Type=Thumbnail&ID=20141106154635467418000000:0&Location=0&Format=STANDARD-XML
func (sess *Session) Object(query *ObjectQuery, addtlParams map[string]string) ([]results.ObjectResult, error) {
	queryString := map[string]string{
		"Resource": query.Resource,
		"Type":     query.Type,
		"ID":       fmt.Sprintf("%s:%s", query.Id, query.Index),
		"Location": strconv.Itoa(query.Location),
	}

	// check for multipart results
	if query.Index == "*" {
		// fire off the query
		multipartObjects, err := sess.getMultipartResults("Search", "GET", sess.Capabilities.GetObjectUrl(), queryString)

		if err != nil {
			return nil, err
		}
		return multipartObjects, nil
	}
	return nil, fmt.Errorf("Not implemented, sorry :(")
}
