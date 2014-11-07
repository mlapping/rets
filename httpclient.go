// httpclient
package rets

import (
	"fmt"
	"github.com/mlapping/rets/results"
	"net/http"
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
	params.Add("Format", "STANDARD-XML")
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
