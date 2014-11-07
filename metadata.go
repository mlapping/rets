// metadata
package rets

import (
	"fmt"
	"github.com/mlapping/rets/results"
)

// For lookup types that have several values. Essentially, the id is the primary key to an enumerated field
// Query looks something like: Type=METADATA-LOOKUP_TYPE&ID=Property:20130414180426065306000000
func (sess *Session) TypeLookup(typeId string) (*results.MetadataReply, error) {
	queryString := map[string]string{
		"Type": "METADATA-LOOKUP_TYPE",
		"ID":   fmt.Sprintf("Property:%s", typeId),
	}

	// fire off the query
	metaReply := &results.MetadataReply{}
	err := sess.getResults("TypeLookup", "GET", sess.Capabilities.MetadataUrl(), queryString, metaReply)

	return metaReply, err
}

/*
	Get information about the classes nested under the given resource
	Ex: {Property (Resource): [A (Class), B (Class), C (Class)]}
*/
// Type=METADATA-CLASS&ID=Property
func (sess *Session) Classes(resource string) (*results.MetadataReply, error) {
	queryString := map[string]string{
		"Type": "METADATA-CLASS",
		"ID":   resource,
	}

	// fire off the query
	metaReply := &results.MetadataReply{}
	err := sess.getResults("Classes", "GET", sess.Capabilities.MetadataUrl(), queryString, metaReply)

	return metaReply, err
}

// Type=METADATA-RESOURCE&ID=0
func (sess *Session) Resources() (*results.MetadataReply, error) {
	queryString := map[string]string{
		"Type": "METADATA-RESOURCE",
		"ID":   string(0),
	}

	// fire off the query
	metaReply := &results.MetadataReply{}
	err := sess.getResults("Resources", "GET", sess.Capabilities.MetadataUrl(), queryString, metaReply)

	return metaReply, err
}

// Type=METADATA-TABLE&ID=Property:B
func (sess *Session) Fields(resource, class string) (*results.MetadataReply, error) {
	queryString := map[string]string{
		"Type": "METADATA-TABLE",
		"ID":   fmt.Sprintf("%s:%s", resource, class),
	}

	// fire off the query
	metaReply := &results.MetadataReply{}
	err := sess.getResults("Fields", "GET", sess.Capabilities.MetadataUrl(), queryString, metaReply)

	return metaReply, err
}
