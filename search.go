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

	for key, value := range addtlParams {
		queryString[key] = value
	}

	// fire off the query
	searchReply := &results.SearchReply{}
	err := sess.getResults("Search", "GET", sess.Capabilities.SearchUrl(), queryString, searchReply)

	if err == nil && searchReply.Code != 0 {
		return nil, fmt.Errorf("rets.%s: Error: %s", "Search", searchReply.Text)
	}

	return searchReply, err

}

// SearchType=Property&Class=A&Query=*&QueryType=DMQL2&Select=LIST_1,LIST_105&Limit=NONE&Format=COMPACT
/*
	Using a special feature in RETS 1.7.2, you can get a key for every record in the database regardless of your search query hard-limit
*/
func (sess *Session) AllKeys(resource, class string) (*results.SearchReply, error) {
	keyField, err := sess.KeyField(resource)
	if err != nil {
		return nil, fmt.Errorf("rets.AllKeys: Resource %s does not have a KeyField. This feature is not available for this resource.", resource)
	}

	queryString := map[string]string{
		"SearchType": resource,
		"Class":      class,
		"QueryType":  "DMQL2",
		"Query":      "*",
		"Limit":      "NONE",
		"Format":     "COMPACT",
		"Select":     keyField,
	}

	// fire off the query
	searchReply := &results.SearchReply{Data: make([]string, 0)}
	err = sess.getResults("Search", "GET", sess.Capabilities.SearchUrl(), queryString, searchReply)

	if err == nil && searchReply.Code != 0 {
		return nil, fmt.Errorf("rets.%s: Error: %s", "Search", searchReply.Text)
	}

	delimiter, err := searchReply.GetDelimiter()
	if err != nil {
		return nil, err
	}

	for i, value := range searchReply.Data {
		searchReply.Data[i] = results.DecodeData(value, delimiter)[1]
	}

	return searchReply, err
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}
