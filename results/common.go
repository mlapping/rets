// common
package results

type Setter interface {
	Set(key, value string)
}

type RetsReply struct {
	Code int    `xml:"ReplyCode,attr"`
	Text string `xml:"ReplyText,attr"`
}
