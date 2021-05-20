/**
* @program: utils
*
* @description:
*
* @author: lemo
*
* @create: 2021-05-20 15:41
**/

package utils

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	url2 "net/url"
	"os"
	"strconv"
	"strings"

	"github.com/json-iterator/go"
)

func getRequest(method string, url string, info *httpInfo) (*http.Request, context.CancelFunc, error) {
	var contentType = strings.ToLower(getContentType(info))
	switch contentType {
	case applicationFormUrlencoded:
		return getPostFormUrlencoded(method, url, info)
	case applicationJson:
		return getPostJson(method, url, info)
	case multipartFormData:
		return getPostFormData(method, url, info)
	default:
		return doUrl(method, url, info)
	}
}

func getPostFormData(method string, url string, info *httpInfo) (*http.Request, context.CancelFunc, error) {
	if info.body == nil {
		info.body = []map[string]interface{}{}
	}

	body, ok := info.body.([]map[string]interface{})
	if !ok {
		return nil, nil, errors.New("application/x-www-form-urlencoded body must be map[string]interface")
	}

	var buf = new(bytes.Buffer)
	part := multipart.NewWriter(buf)

	for i := 0; i < len(body); i++ {
		for key, value := range body[i] {
			switch value.(type) {
			case string:
				if err := part.WriteField(key, value.(string)); err != nil {
					return nil, nil, err
				}
			case int:
				if err := part.WriteField(key, strconv.Itoa(value.(int))); err != nil {
					return nil, nil, err
				}
			case float64:
				if err := part.WriteField(key, strconv.FormatFloat(value.(float64), 'f', -1, 64)); err != nil {
					return nil, nil, err
				}
			case *os.File:
				ff, err := part.CreateFormFile(key, value.(*os.File).Name())
				if err != nil {
					return nil, nil, err
				}
				_, err = io.Copy(ff, value.(*os.File))
				if err != nil {
					return nil, nil, err
				}
			default:
				if err := part.WriteField(key, fmt.Sprintf("%v", value)); err != nil {
					return nil, nil, err
				}
			}
		}
	}

	if err := part.Close(); err != nil {
		return nil, nil, err
	}

	var ctx, cancel = context.WithCancel(context.Background())
	request, err := http.NewRequestWithContext(ctx, method, url, buf)
	if err != nil {
		cancel()
		return nil, nil, err
	}
	info.SetHeader(contentType, part.FormDataContentType())
	return request, cancel, err
}

func getPostJson(method string, url string, info *httpInfo) (*http.Request, context.CancelFunc, error) {
	body, ok := info.body.([]interface{})
	if !ok {
		return nil, nil, errors.New("application/json body must be interface")
	}

	var jsonBody []byte

	for i := 0; i < len(body); i++ {
		b, err := jsoniter.Marshal(info.body)
		if err != nil {
			return nil, nil, err
		}
		jsonBody = append(jsonBody, b...)
	}

	var ctx, cancel = context.WithCancel(context.Background())
	request, err := http.NewRequestWithContext(ctx, method, url, bytes.NewReader(jsonBody))
	if err != nil {
		cancel()
		return nil, nil, err
	}
	return request, cancel, err
}

func getPostFormUrlencoded(method string, url string, info *httpInfo) (*http.Request, context.CancelFunc, error) {
	if info.body == nil {
		info.body = []map[string]interface{}{}
	}

	body, ok := info.body.([]map[string]interface{})
	if !ok {
		return nil, nil, errors.New("application/x-www-form-urlencoded body must be map[string]interface")
	}

	var buff bytes.Buffer
	for i := 0; i < len(body); i++ {
		for key, value := range body[i] {
			switch value.(type) {
			case string:
				buff.WriteString(key + "=" + value.(string) + "&")
			case int:
				buff.WriteString(key + "=" + strconv.Itoa(value.(int)) + "&")
			case float64:
				buff.WriteString(key + "=" + strconv.FormatFloat(value.(float64), 'f', -1, 64) + "&")
			default:
				buff.WriteString(key + "=" + fmt.Sprintf("%v", value) + "&")
			}
		}
	}

	var b = buff.Bytes()
	if len(b) != 0 {
		b = b[:len(b)-1]
	}

	var ctx, cancel = context.WithCancel(context.Background())
	request, err := http.NewRequestWithContext(ctx, method, url, bytes.NewReader(b))
	if err != nil {
		cancel()
		return nil, nil, err
	}
	return request, cancel, err
}

func doUrl(method string, url string, info *httpInfo) (*http.Request, context.CancelFunc, error) {
	Url, err := url2.Parse(url)
	if err != nil {
		return nil, nil, err
	}

	if info.body == nil {
		info.body = []map[string]interface{}{}
	}

	body, ok := info.body.([]map[string]interface{})
	if !ok {
		return nil, nil, err
	}

	var params = url2.Values{}

	for i := 0; i < len(body); i++ {
		for key, value := range body[i] {
			switch value.(type) {
			case string:
				params.Set(key, value.(string))
			case int:
				params.Set(key, strconv.Itoa(value.(int)))
			case float64:
				params.Set(key, strconv.FormatFloat(value.(float64), 'f', -1, 64))
			default:
				params.Set(key, fmt.Sprintf("%v", value))
			}
		}
	}

	var pStr = params.Encode()

	if pStr != "" {
		Url.RawQuery = Url.RawQuery + "&" + pStr
	}

	var ctx, cancel = context.WithCancel(context.Background())
	request, err := http.NewRequestWithContext(ctx, method, Url.String(), nil)
	if err != nil {
		cancel()
		return nil, nil, err
	}
	return request, cancel, err
}

func getContentType(info *httpInfo) string {
	for i := 0; i < len(info.headerKey); i++ {
		if info.headerKey[i] == contentType {
			return info.headerValue[i]
		}
	}
	return ""
}

func send(info *httpInfo, request *http.Request, cancel context.CancelFunc) *Request {

	hMux.Lock()
	defer hMux.Unlock()
	defer cancel()

	if request == nil {
		return &Request{err: errors.New("invalid request")}
	}

	for i := 0; i < len(info.headerKey); i++ {
		request.Header.Add(info.headerKey[i], info.headerValue[i])
	}

	for i := 0; i < len(info.cookies); i++ {
		request.AddCookie(info.cookies[i])
	}

	if info.userName != "" || info.passWord != "" {
		request.SetBasicAuth(info.userName, info.passWord)
	}

	if info.clientTimeout != 0 {
		defaultClient.Timeout = info.clientTimeout
	}

	if info.dialerKeepAlive != 0 {
		defaultDialer.KeepAlive = info.dialerKeepAlive
	}

	if info.proxy != nil {
		defaultTransport.Proxy = info.proxy
	}

	if info.progress != nil {
		defaultTransport.DisableCompression = true
	}

	defer func() {
		defaultClient.Timeout = clientTimeout
		defaultDialer.KeepAlive = dialerKeepAlive
		defaultTransport.Proxy = http.ProxyFromEnvironment
		defaultTransport.DisableCompression = false
	}()

	response, err := defaultClient.Do(request)
	if err != nil {
		return &Request{err: err}
	}
	defer func() { _ = response.Body.Close() }()

	var dataBytes []byte

	if info.progress != nil {
		var total, _ = strconv.ParseInt(response.Header.Get(contentLength), 10, 64)
		var writer = &writeProgress{
			total:      total,
			onProgress: info.progress.progress,
			rate:       info.progress.rate,
		}

		dataBytes, err = ioutil.ReadAll(io.TeeReader(response.Body, writer))
		if err != nil {
			return &Request{err: err}
		}
	} else {
		dataBytes, err = ioutil.ReadAll(response.Body)
		if err != nil {
			return &Request{err: err}
		}
	}

	return &Request{code: response.StatusCode, data: dataBytes, requestHeader: response.Request.Header, responseHeader: response.Header}
}
