// search
package rets

//import (
//	"encoding/xml"
//)

func (sess *Session) Search(dmql string) {
	req, _ := newRequest("GET", sess.BaseAddress+sess.Capabilities.Search)
	res, _ := sess.HttpClient.Do(req)

	defer res.Body.Close()
	//properties := xml.NewDecoder(res.Body)

}
