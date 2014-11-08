// xmlutils
package results

import (
	"bytes"
	"code.google.com/p/go-charset/charset"
	"encoding/xml"
	"errors"
	//	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

// Used For Debugging
func readerToString(r interface {
	io.Reader
}) string {
	byteArr, _ := ioutil.ReadAll(r)
	b := bytes.NewBuffer(byteArr)
	return b.String()
}

func stringToReader(s string) interface {
	io.Reader
} {
	return strings.NewReader(s)
}

/*
	Converts RETS XML to a struct
*/
func ConvertServerResponse(results io.ReadCloser, object interface{}) error {
	defer results.Close()

	// setup a decoder

	// uncomment to debug
	//str := readerToString(results)
	//fmt.Println(str)
	//xmlReader := xml.NewDecoder(stringToReader(str))
	xmlReader := xml.NewDecoder(results)

	xmlReader.CharsetReader = charset.NewReader

	err := xmlReader.Decode(object)
	if err != nil {
		return errors.New("rets/results.ConvertServerResponse: Error with response encoding: inner exception => " + err.Error())
	}
	return nil
}
