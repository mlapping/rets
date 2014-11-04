// connection
package rets

import (
	"errors"
	"github.com/edmore/goca/auth"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

type Session struct {
	DialInfo     *DialInfo
	HttpClient   *http.Client
	BaseAddress  string
	LoggedIn     bool
	Capabilities Capabilities
}

type DialInfo struct {
	LoginUrl string
	UserName string
	Password string
}

func DialWithInfo(info *DialInfo) (*Session, error) {
	err := info.validateDialInfo()
	if err != nil {
		return nil, err
	}

	// create a new session
	session := &Session{}
	session.DialInfo = info

	// start setting up the capabilities early
	loginUrl, err := url.Parse(session.DialInfo.LoginUrl)
	if err != nil {
		return nil, err
	}

	session.Capabilities = Capabilities{
		Host: loginUrl.Host,
	}

	// create an http connection with a cookiejar
	cookieJar, _ := cookiejar.New(nil)
	session.HttpClient = &http.Client{
		Jar: cookieJar,
	}

	err = session.tryToLogin()
	if err != nil {
		return nil, err
	}

	// Login was successful

	// attempt to login
	return nil, nil
}

func (sess *Session) tryToLogin() error {

	// The first GET will fail, but set our cookie jar
	req, err := newRequest("GET", sess.DialInfo.LoginUrl)
	if err != nil {
		return err
	}

	// The first GET will fail, but set our cookie jar
	res, _ := sess.HttpClient.Do(req)

	// The second GET will not fail
	req, err = newRequest("GET", sess.DialInfo.LoginUrl)
	// set our digest auth
	req = auth.SetDigestAuth(req, sess.DialInfo.UserName, sess.DialInfo.Password, res, 1)
	res, err = sess.HttpClient.Do(req)

	if res.StatusCode != http.StatusOK {
		return errors.New("login failed. username and/or password is incorrect")
	}

	// setup our capabilities
	sess.Capabilities.setFromLogin(res.Body)

	return nil
}

func newRequest(method string, url string) (*http.Request, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	// set some request parameters related to RETS
	req.Header.Add("User-Agent", "golang rets client v1.0")
	req.Header.Add("RETS-VERSION", "RETS/1.7.2")

	return req, nil
}

func (dial *DialInfo) validateDialInfo() error {
	if dial == nil {
		return errors.New("passed a nil DialInfo object")
	}

	if len(dial.Password) == 0 {
		return errors.New("passed a blank password")
	}

	if len(dial.UserName) == 0 {
		return errors.New("passed a blank username")
	}

	return nil
}
