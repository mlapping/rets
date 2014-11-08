// search
package rets

import (
	"fmt"
	"github.com/mlapping/rets/results"
	"strconv"
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

/*
	Run a DMQL Search query on a Resource/Class pair

	SearchType=Property&Class=A&QueryType=DMQL2&Query=%28LIST_15=%7COV61GOJ13C0%29&StandardNames=0&Limit=50

*/
func (sess *Session) Search(query *SearchQuery, addtlParams map[string]string) (*results.SearchReply, error) {
	queryString := map[string]string{
		"SearchType":    query.Resource,
		"Class":         query.Class,
		"QueryType":     "DMQL2",
		"Query":         query.DMQL,
		"Count":         strconv.Itoa(int(query.Count)),
		"StandardNames": strconv.Itoa(boolToInt(query.StandardNames)),
		"Limit":         strconv.Itoa(query.Limit),
		"Format":        "COMPACT-DECODED",
	}

	// fire off the query
	searchReply := &results.SearchReply{}
	err := sess.getResults("Search", "GET", sess.Capabilities.SearchUrl(), queryString, searchReply)

	if err == nil && searchReply.Code != 0 {
		return nil, fmt.Errorf("rets.%s: Error: %s", "Search", searchReply.Text)
	}

	return searchReply, err

}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}
