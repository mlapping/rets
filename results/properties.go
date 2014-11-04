// properties
package results

type Query string

type SearchResults struct {
	DmqlQuery  Query
	Properties map[string]string
}
