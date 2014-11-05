// search
package rets

//import (
//	"github.com/mlapping/rets/results"
//)

type SearchQuery struct {
	DMQL     string
	Limit    int
	Resource string
	Class    string
	Offset   int
	Count    bool
}

// SearchType=Property&
// Class=A&
// QueryType=DMQL2&
// Query=%28LIST_15=%7COV61GOJ13C0%29&
// Count=0&
// StandardNames=0&
// Limit=50
func (sess *Session) Search(dmql string) {
	//req, _ := newRequest("GET", sess.BaseAddress+sess.Capabilities.Search, nil)

	//res, _ := sess.HttpClient.Do(req)

}
