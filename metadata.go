// metadata
package rets

import (
	"errors"
	"fmt"
	"github.com/mlapping/rets/results"
)

// For lookup types that have several values. Essentially, the id is the primary key to an enumerated field
// Query looks something like: Type=METADATA-LOOKUP_TYPE&ID=Property:20130414180426065306000000
func (sess *Session) TypeLookup(typeId string) (*results.MetadataLookupType, error) {
	queryString := map[string]string{
		"Type":   "METADATA-LOOKUP_TYPE",
		"ID":     fmt.Sprintf("Property:%s", typeId),
		"Format": "STANDARD-XML",
	}

	// fire off the query
	metaReply := &results.MetadataReply{}
	err := sess.getResults("TypeLookup", "GET", sess.Capabilities.MetadataUrl(), queryString, metaReply)

	if err != nil {
		return nil, err
	}

	return &metaReply.Metadata.LookupType, nil
}

/*
	Get information about the classes nested under the given resource
	Ex: {Property (Resource): [A (Class), B (Class), C (Class)]}
*/
// Type=METADATA-CLASS&ID=Property
func (sess *Session) Classes(resource string) ([]results.MetadataClass, error) {
	queryString := map[string]string{
		"Type":   "METADATA-CLASS",
		"ID":     resource,
		"Format": "STANDARD-XML",
	}

	// fire off the query
	metaReply := &results.MetadataReply{}
	err := sess.getResults("Classes", "GET", sess.Capabilities.MetadataUrl(), queryString, metaReply)

	if err != nil {
		return nil, err
	}

	return metaReply.Metadata.Class.Classes, nil
}

// Type=METADATA-RESOURCE&ID=0
func (sess *Session) Resources() ([]results.MetadataResource, error) {
	queryString := map[string]string{
		"Type":   "METADATA-RESOURCE",
		"ID":     "0",
		"Format": "STANDARD-XML",
	}

	// fire off the query
	metaReply := &results.MetadataReply{}
	err := sess.getResults("Resources", "GET", sess.Capabilities.MetadataUrl(), queryString, metaReply)

	if err != nil {
		return nil, err
	}

	return metaReply.Metadata.Resource.Resources, nil
}

// Type=METADATA-TABLE&ID=Property:B
func (sess *Session) Fields(resource, class string) ([]results.MetadataField, error) {
	queryString := map[string]string{
		"Type":   "METADATA-TABLE",
		"ID":     fmt.Sprintf("%s:%s", resource, class),
		"Format": "STANDARD-XML",
	}

	// fire off the query
	metaReply := &results.MetadataReply{}
	err := sess.getResults("Fields", "GET", sess.Capabilities.MetadataUrl(), queryString, metaReply)

	if err != nil {
		return nil, err
	}

	return metaReply.Metadata.Table.Fields, nil
}

func (sess *Session) KeyField(resourceName string) (string, error) {
	resources, err := sess.Resources()
	if err != nil {
		return "", err
	}

	for _, resource := range resources {
		if resource.Id == resourceName {
			return resource.KeyField, nil
		}
	}

	return "", errors.New("rets.KeyField: Invalid resource name given")
}
