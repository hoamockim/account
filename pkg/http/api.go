package http

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"time"
)

type ApiClient struct {
	client *http.Client
}

type ClientInfo struct {
	BaseUrl  string
	UserName string
	PassWord string
	JwtAuth  string
	AuthType int // 1: Jwt, 2: BasicAuth, 0: default
}

func NewClient(tlsConfig *tls.Config, connectTimeout time.Duration, requestTimeout time.Duration) *ApiClient {
	transport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   connectTimeout,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		ResponseHeaderTimeout: requestTimeout,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		TLSClientConfig:       tlsConfig,
	}

	return &ApiClient{client: &http.Client{
		Timeout:   requestTimeout,
		Transport: transport,
	}}
}

func (clInfo *ClientInfo) makeEndpoint(path string) (endpoint *url.URL, err error) {
	btu := []byte(clInfo.BaseUrl)
	if btu[len(btu)-1] == '/' {
		newUrl := btu[:len(btu)-2]
		clInfo.BaseUrl = string(newUrl)
	}
	urlPath := clInfo.BaseUrl + "/" + path

	if endpoint, err = url.Parse(urlPath); err != nil {
		return nil, errors.New(fmt.Sprintf("path service is invalid: %s", path))
	}
	return
}

func (ac *ApiClient) Call(method, requestId string, clientInfo *ClientInfo, path string, reqBody interface{}, resData interface{}) error {
	var err error
	var payload io.Reader
	var endpoint *url.URL
	var resp *http.Response
	if payload, err = toJSON(reqBody); err != nil {
		return errors.New(fmt.Sprintf("request body is invalid: %s", err))
	}

	if endpoint, err = clientInfo.makeEndpoint(path); err != nil {
		return err
	}

	req, _ := http.NewRequest(method, endpoint.String(), payload)
	req.Header.Add("X-Request-ID", requestId)
	req.Header.Add("Content-Type", "application/json")

	switch clientInfo.AuthType {
	case 1:
		req.Header.Add("Authorization", clientInfo.JwtAuth)
		break
	case 2:
		req.SetBasicAuth(clientInfo.UserName, clientInfo.PassWord)
		break
	default:
		break
	}

	if resp, err = ac.client.Do(req); err != nil {
		return errors.New(fmt.Sprintf("call api error: %s, %s, %s", requestId, endpoint, err))
	}

	if err = json.NewDecoder(resp.Body).Decode(resData); err != nil {
		return errors.New(fmt.Sprintf("response data is invallid: %s, %s", requestId, err))
	}
	resp.Body.Close()

	return nil
}

func toJSON(v interface{}) (io.Reader, error) {
	data, err := json.Marshal(v)
	return bytes.NewBuffer(data), err
}
