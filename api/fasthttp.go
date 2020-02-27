/**
 * Copyright (c) 2018, TransChain.
 *
 * This source code is licensed under the Apache 2.0 license found in the
 * LICENSE file in the root directory of this source tree.
 */

package api

import (
    "net"
    "time"

    "github.com/valyala/fasthttp"

    "github.com/transchain/sdk-go/uri"
)

// FastHttpClient is a fasthttp.FastHttpClient wrapper to dialog with a JSON API.
type FastHttpClient struct {
    fastHttpClient *fasthttp.Client
    apiUrl         string
    headers        map[string]string
}

// FastHttpClient constructor.
func NewFastHttpClient(apiUrl string) *FastHttpClient {
    return &FastHttpClient{
        fastHttpClient: &fasthttp.Client{
            Dial: func(addr string) (net.Conn, error) {
                return fasthttp.DialTimeout(addr, 15*time.Second)
            },
            MaxIdemponentCallAttempts: 1,
        },
        apiUrl:  apiUrl,
        headers: make(map[string]string),
    }
}

// AddHeader adds a persistent header that will be sent in every future doRequest calls.
func (c FastHttpClient) AddHeader(key string, value string) {
    c.headers[key] = value
}

// RemoveHeader removes a persistent header.
func (c FastHttpClient) RemoveHeader(key string) {
    delete(c.headers, key)
}

// Get wraps the doRequest method to do a GET HTTP request.
func (c FastHttpClient) Get(
    route string,
    headers map[string]string,
    queryValues map[string]string,
) (*RawResponse, error) {
    return c.doRequest("GET", route, nil, headers, queryValues)
}

// Post wraps the doRequest method to do a POST HTTP request.
func (c FastHttpClient) Post(
    route string,
    body []byte,
    headers map[string]string,
    queryValues map[string]string,
) (*RawResponse, error) {
    return c.doRequest("POST", route, body, headers, queryValues)
}

// doRequest uses the fasthttp.FastHttpClient to call a distant api and returns a response.
func (c FastHttpClient) doRequest(
    method string,
    route string,
    body []byte,
    headers map[string]string,
    queryValues map[string]string,
) (*RawResponse, error) {

    req := fasthttp.AcquireRequest()
    resp := fasthttp.AcquireResponse()

    req.SetConnectionClose()

    defer func() {
        if req != nil {
            fasthttp.ReleaseRequest(req)
        }
        if resp != nil {
            fasthttp.ReleaseResponse(resp)
        }
    }()

    fullUri, err := uri.BuildUri(c.apiUrl, []string{route}, queryValues)
    if err != nil {
        return nil, err
    }
    req.SetRequestURI(fullUri.String())

    if body != nil {
        req.SetBody(body)
    }

    req.Header.SetMethod(method)

    for key, value := range c.headers {
        req.Header.Set(key, value)
    }

    for key, value := range headers {
        req.Header.Set(key, value)
    }

    err = c.fastHttpClient.Do(req, resp)
    if err != nil {
        return nil, err
    }

    originalBody := resp.Body()
    copiedBody := make([]byte, len(originalBody))
    copy(copiedBody, originalBody)

    return &RawResponse{
        StatusCode: resp.StatusCode(),
        Body:       copiedBody,
    }, nil
}
