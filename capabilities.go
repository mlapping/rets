// capabilities
package rets

import (
	"fmt"
	"github.com/mlapping/rets/results"
	"io"
)

type Capabilities struct {
	Host        string
	GetObject   string
	GetMetadata string
	Login       string
	Logout      string
	Search      string
}

type RetsReturnType struct {
}

type RetsCapabilities struct {
}

func (cap *Capabilities) setFromLogin(responseBody interface {
	io.Reader
	io.Closer
}) {
	reply := results.ConvertServerResponse(responseBody)

	fmt.Println(err)
	fmt.Println(representation)
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
