// object
package results

import (
	"io"
	"strconv"
)

//Location:http://cdn.photos.flexmls.com/knx/20141106174321256467000000-t.jpg
//MIME-Version:1.0
//Object-ID:1
//Preferred:1
//RETS-Version:RETS/1.5
//Server:Apache-Coyote/1.1
//Set-Cookie:JSESSIONID=CC297D51FEF893664506CE5A20101C6D; Path=/rets2_1
type ObjectResult struct {
	ObjectId    int
	Preferred   bool
	Location    string
	Description string
	MimeInfo    MIME
	io.ReadCloser
}

func (or *ObjectResult) Set(key, value string) {
	switch key {
	case "Object-ID":
		val, err := strconv.Atoi(value)
		if err != nil {
			or.ObjectId = val
		}
	case "Preferred":
		if value == "1" {
			or.Preferred = true
		} else {
			or.Preferred = false
		}
	case "Location":
		or.Location = value
	case "Content-Description":
		or.Description = value
	case "Content-Type":
		or.MimeInfo.Type = value
	case "MIME-Version":
		or.MimeInfo.Version = value
	}
}

type MIME struct {
	Type    string
	Version string
}
