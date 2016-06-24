package api

import (
	"github.com/bitly/go-simplejson"
	"github.com/levigross/grequests"
)

const (
	hostCN = "https://api.leancloud.cn"
	hostUS = "https://us-api.leancloud.cn"
)

// API server regions
const (
	RegionInvalid = iota
	RegionCN
	RegionUS
)

const apiVersion = "1.1"

// Client info
type Client struct {
	AppID     string
	MasterKey string
	Region    int
	inited    bool // eg: if router is fetched
}

func fetchRouter() error {
	// TODO: fetch router from server
	return nil
}

func (client *Client) baseURL() string {
	switch client.Region {
	case RegionCN:
		return hostCN + "/" + apiVersion
	case RegionUS:
		return hostUS + "/" + apiVersion
	default:
		panic("invalid region")
	}
}

func (client *Client) options() *grequests.RequestOptions {
	return &grequests.RequestOptions{
		Headers: map[string]string{
			"X-AVOSCloud-Application-Id":         client.AppID,
			"X-AVOSCloud-Master-Key":             client.MasterKey,
			"X-AVOSCloud-Application-Production": "1",
			"Content-Type":                       "application/json",
		},
	}
}

func (client *Client) get(path string, options *grequests.RequestOptions) (*simplejson.Json, error) {
	if options == nil {
		options = client.options()
	}
	resp, err := grequests.Get(client.baseURL()+path, options)
	if err != nil {
		return nil, err
	}
	if !resp.Ok {
		return nil, NewErrorFromBody(resp.String())
	}
	return simplejson.NewFromReader(resp)
}

func (client *Client) post(path string, params map[string]interface{}, options *grequests.RequestOptions) (*simplejson.Json, error) {
	if options == nil {
		options = client.options()
	}
	options.JSON = params
	resp, err := grequests.Post(client.baseURL()+path, options)
	if err != nil {
		return nil, err
	}
	if !resp.Ok {
		return nil, NewErrorFromBody(resp.String())
	}
	return simplejson.NewFromReader(resp)
}

func (client *Client) delete(path string, options *grequests.RequestOptions) (*simplejson.Json, error) {
	if options == nil {
		options = client.options()
	}
	resp, err := grequests.Delete(client.baseURL()+path, options)
	if err != nil {
		return nil, err
	}
	if !resp.Ok {
		return nil, NewErrorFromBody(resp.String())
	}
	return simplejson.NewFromReader(resp)
}