// httpclient
package rets

import (
	"fmt"
	"github.com/mlapping/rets/results"
	"io"
	"mime"
	"mime/multipart"
	"net/http"
	"strings"
)

func (sess *Session) newRequest(method string, url string, query map[string]string) (*http.Request, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	// set some request parameters related to RETS
	req.Header.Add("User-Agent", "golang rets client v1.0")
	req.Header.Add("RETS-VERSION", "RETS/1.7.2")

	if sess.Authorization != "" {
		req.Header.Add("Authorization", sess.Authorization)
	}
	// parse the query params
	params := req.URL.Query()
	for key, value := range query {
		params.Add(key, value)
	}
	req.URL.RawQuery = params.Encode()

	return req, nil
}

func (sess *Session) getResults(callingFunction, httpMethod, url string, query map[string]string, object interface{}) error {
	req, _ := sess.newRequest(httpMethod, url, query)
	res, err := sess.HttpClient.Do(req)

	if err != nil {
		return fmt.Errorf("rets.%s: Error with http request. Inner exception => %v", callingFunction, err)
	}

	err = results.ConvertServerResponse(res.Body, object)
	if err != nil {
		return fmt.Errorf("rets.%s: Error converting xml to %T. Inner exception => %v", callingFunction, object, err)
	}

	return nil
}

func (sess *Session) getMultipartResults(callingFunction, httpMethod, url string, query map[string]string) ([]results.ObjectResult, error) {
	req, _ := sess.newRequest(httpMethod, url, query)
	res, err := sess.HttpClient.Do(req)

	if err != nil {
		return nil, fmt.Errorf("rets.%s: Error with http request. Inner exception => %v", callingFunction, err)
	}

	// check that the RETS server returned a multi-part header like we're expecting
	mediaType, params, err := mime.ParseMediaType(res.Header.Get("Content-Type"))
	if err != nil {
		return nil, fmt.Errorf("rets.%s: Error parsing the http headers content-type. Inner exception => %v", callingFunction, err)
	}

	multipartResults := make([]results.ObjectResult, 0)

	// Read each part
	if strings.HasPrefix(mediaType, "multipart/") {
		mr := multipart.NewReader(res.Body, params["boundary"])
		for {
			p, err := mr.NextPart()
			if err == io.EOF {
				break
			}

			if err != nil {
				return nil, fmt.Errorf("rets.%s: Error reading multi-part response. Inner exception => %v", callingFunction, err)
			}

			obj := results.ObjectResult{}
			for key, value := range p.Header {
				obj.Set(key, value[0])
			}
			multipartResults = append(multipartResults, obj)
		}
	} else if mediaType == "text/xml" {
		reply := &results.RetsReply{}
		err := results.ConvertServerResponse(res.Body, reply)

		if err == nil && reply.Code != 0 {
			return nil, fmt.Errorf("rets.%s: Error from RETS SERVER: %s", callingFunction, reply.Text)
		}
		return nil, fmt.Errorf("rets.%s: Error parsing response from RETS server: %v", callingFunction, err)
	} else {
		return nil, fmt.Errorf("rets.%s: Http header does not indicate a multipart response from the server. RETS server did not send a response", callingFunction)
	}

	return multipartResults, nil
}
