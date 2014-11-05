// metadata
package results

import (
	"time"
)

type MetadataReply struct {
	Code     int      `xml:"ReplyCode,attr"`
	Text     string   `xml:"ReplyText,attr"`
	Metadata Metadata `xml:"METADATA"`
}

type Metadata struct {
	Table      MetadataTable           `xml:"METADATA-TABLE"`
	Resource   MetadataResourceListing `xml:"METADATA-RESOURCE"`
	Class      MetadataClassListing    `xml:"METADATA-CLASS"`
	LookupType MetadataLookupType      `xml:"METADATA-LOOKUP_TYPE"`
}

// Type=METADATA-LOOKUP_TYPE&ID=Property:20130414180426065306000000
type MetadataLookupType struct {
	Name     string                `xml:"Lookup,attr"`
	Resource string                `xml:",attr"`
	System   string                `xml:",attr`
	Values   []MetadataLookupValue `xml:"Lookup"`
}

type MetadataLookupValue struct {
	LongValue  string
	ShortValue string
	Value      string
}

type MetadataClassListing struct {
	Date     time.Time       `xml:",attr"`
	System   string          `xml:",attr"`
	Resource string          `xml:"Resource,attr"`
	Classes  []MetadataClass `xml:"Class"`
}

// Type=METADATA-CLASS&ID=Property
type MetadataClass struct {
	ClassName     string
	StandardName  string
	VisibleName   string
	Description   string
	TableVersion  string
	TableDate     string
	UpdateVersion string
	UpdateDate    string
}

type MetadataResourceListing struct {
	Date      time.Time          `xml:",attr"`
	System    string             `xml:",attr"`
	Resources []MetadataResource `xml:"Resource"`
}

type MetadataTable struct {
	Resource string          `xml:",attr"`
	Class    string          `xml:",attr"`
	Date     time.Time       `xml:",attr"`
	System   string          `xml:",attr"`
	Fields   []MetadataField `xml:"Field"`
}

// Type=METADATA-RESOURCE&ID=0
type MetadataResource struct {
	Id                          string
	StandardName                string
	VisibleName                 string
	Description                 string
	KeyField                    string
	ClassCount                  string
	ClassVersion                string
	ClassDate                   string
	ObjectVersion               string
	ObjectDate                  string
	SearchHelpVersion           string
	SearchHelpDate              string
	EditMaskVersion             string
	EditMaskDate                string
	LookupVersion               string
	LookupDate                  string
	UpdateHelpVersion           string
	UpdateHelpDate              string
	ValidationExpressionVersion string
	ValidationExpressionDate    string
	ValidationLookupVersion     string
	ValidationLookupDate        string
	ValidationExternalVersion   string
	ValidationExternalDate      string
}

// Type=METADATA-TABLE&ID=Property:A
type MetadataField struct {
	SystemName     string
	StandardName   string
	LongName       string
	DBName         string
	ShortName      string
	MaximumLength  int
	DataTyp        string
	Precision      int
	Searchable     bool
	Interpretatiom string
	Alignment      string
	UseSeparator   string
	EditMaskID     string
	LookupName     string
	MaxSelect      string
	Units          string
	Index          string
	Minimum        string
	Maximum        string
	Default        string
	Required       bool
	SearchHelpID   string
	Unique         bool
}
