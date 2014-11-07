// search
package rets

import (
	"github.com/mlapping/rets/results"
)

type Count int

const (
	COUNT_OFF                Count = iota
	COUNT_PREPEND_TO_RESULTS Count = iota
	COUNT_IS_ONLY_RESULT     Count = iota
)

type SearchQuery struct {
	DMQL          string
	Limit         int
	Resource      string
	Class         string
	Offset        int
	Count         Count
	StandardNames bool
}

// SearchType=Property&Class=A&QueryType=DMQL2&Query=%28LIST_15=%7COV61GOJ13C0%29&StandardNames=0&Limit=50
/*
	Run a DMQL Search query on a Resource/Class pair
*/
func (sess *Session) Search(query *SearchQuery, addtlParams map[string]string) (interface{}, error) {
	queryString := map[string]string{
		"SearchType":    query.Resource,
		"Class":         query.Class,
		"QueryType":     "DMQL2",
		"Query":         query.DMQL,
		"Count":         string(query.Count),
		"StandardNames": string(boolToInt(query.StandardNames)),
		"Limit":         string(query.Limit),
	}

	// fire off the query
	metaReply := &results.MetadataReply{}
	err := sess.getResults("TypeLookup", "GET", sess.Capabilities.SearchUrl(), queryString, metaReply)

	return metaReply, err

}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}
