package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"go.uber.org/zap"
)

const (
	JsonContentType = "application/json"
)

type BaseClient struct {
	Ctx          context.Context
	RequestId    string
	ContentType  string
	BasicAuth    string
	ApiKey       string
	ClientPath   string
	ResponseData interface{}
}

func NewClient(ctx context.Context, requestId string, contentType string, responseData interface{}) *BaseClient {
	return &BaseClient{
		Ctx:          ctx,
		RequestId:    requestId,
		ContentType:  contentType,
		ResponseData: responseData,
	}
}

func (client *BaseClient) makeEndpoint(ctx context.Context, apiPath string) endpoint.Endpoint {
	fullURL, err := url.Parse(fmt.Sprintf("%v/%v", client.ClientPath, apiPath))

	if err != nil {
		zap.S().Errorf("Cannot parse url %v,client path: %v ,%v",
			client.RequestId, client.ClientPath,
			zap.Error(err))
	}
	return httptransport.NewClient(
		http.MethodPost, fullURL,
		client.encodeRequest,
		client.decodeResponse,
	).Endpoint()
}

func (client *BaseClient) encodeRequest(ctx context.Context, r *http.Request, reqBody interface{}) error {
	r.Header.Add("X-Request-ID", client.RequestId)
	r.Header.Add("Content-Type", client.ContentType)

	if strings.TrimSpace(client.BasicAuth) != "" {
		r.Header.Add("Authorization", fmt.Sprintf("Basic %v", client.BasicAuth))
	} else if strings.TrimSpace(client.ApiKey) != "" {
		r.Header.Add("X-Api-Key", client.ApiKey)
	}

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(reqBody)
	if err != nil {
		return err
	}
	r.Body = ioutil.NopCloser(&buf)

	return nil
}

func (client *BaseClient) decodeResponse(ctx context.Context, r *http.Response) (interface{}, error) {
	if err := json.NewDecoder(r.Body).Decode(client.ResponseData); err != nil {
		return nil, err
	}
	return client.ResponseData, nil
}
