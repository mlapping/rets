// search
package results

type SearchReply struct {
	Code      int       `xml:"ReplyCode,attr"`
	Text      string    `xml:"ReplyText,attr"`
	MaxRows   string    `xml:"MAXROWS"` // make this an int
	Delimiter Delimiter `xml:"DELIMITER"`
	Columns   string    `xml:"COLUMNS"`
	Data      []string  `xml:"DATA"`
}

type Delimiter struct {
	Value string `xml:"value,attr"`
}
