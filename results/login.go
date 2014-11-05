// login
package results

type LoginResponse struct {
	Text string `xml:",innerxml"`
}

type LoginReply struct {
	ReplyCode int           `xml:"ReplyCode,attr"`
	ReplyText string        `xml:"ReplyText,attr"`
	Response  LoginResponse `xml:"RETS-RESPONSE"`
}
