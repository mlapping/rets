// capabilities
package rets

import (
	"errors"
	"github.com/mlapping/rets/results"
	"regexp"
)

type Capabilities struct {
	Host        string
	GetObject   string
	GetMetadata string
	Login       string
	Logout      string
	Search      string
}

func (cap *Capabilities) setFromLogin(reply *results.LoginReply) error {
	rMetadata := regexp.MustCompile(`GetMetadata=([^\s]+)`)
	rObject := regexp.MustCompile(`GetObject=([^\s]+)`)
	rSearch := regexp.MustCompile(`Search=([^\s]+)`)
	rLogout := regexp.MustCompile(`Logout=([^\s]+)`)

	cap.GetObject = rObject.FindStringSubmatch(reply.Response.Text)[1]
	cap.GetMetadata = rMetadata.FindStringSubmatch(reply.Response.Text)[1]
	cap.Search = rSearch.FindStringSubmatch(reply.Response.Text)[1]
	cap.Logout = rLogout.FindStringSubmatch(reply.Response.Text)[1]

	return cap.Validate()
}

func (cap *Capabilities) Validate() error {
	errPrefix := "Server did not return valid capabilities for: "

	if cap.GetObject == "" {
		return errors.New(errPrefix + "GetObject")
	}

	if cap.GetMetadata == "" {
		return errors.New(errPrefix + "GetMetadata")
	}

	if cap.Search == "" {
		return errors.New(errPrefix + "Search")
	}

	if cap.Logout == "" {
		return errors.New(errPrefix + "Logout")
	}

	return nil
}

func (cap *Capabilities) LoginUrl() string {
	return cap.Host + cap.Login
}

func (cap *Capabilities) GetObjectUrl() string {
	return cap.Host + cap.GetObject
}

func (cap *Capabilities) MetadataUrl() string {
	return cap.Host + cap.GetMetadata
}

func (cap *Capabilities) LogoutUrl() string {
	return cap.Host + cap.Logout
}

func (cap *Capabilities) SearchUrl() string {
	return cap.Host + cap.Search
}
