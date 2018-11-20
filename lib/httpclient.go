package lib

import (
	"context"
	"net/url"
	"io/ioutil"
	"strings"
	"net/http"
)

func HttpGet(ctx context.Context, path string, params map[string]interface{}, headers map[string]interface{}) (string, error){
	return request(ctx, path, params, headers, "get")
}

func HttpPost(ctx context.Context, path string, params map[string]interface{}, headers map[string]interface{}) (string, error){
	return request(ctx, path, params, headers, "get")
}

func request(ctx context.Context, path string, params map[string]interface{}, headers map[string]interface{}, method string) (string, error){
	formData := url.Values{}
	for key, val := range params {
		formData.Set(key, val.(string))
	}
	body := ioutil.NopCloser(strings.NewReader(formData.Encode()))

	client := &http.Client{}

	request, err := http.NewRequest(method, path, body)
	if err != nil {
		panic("failed to new request")
	}

	for key, val := range headers {
		request.Header.Set(key, val.(string))
	}

	resp, err := client.Do(request)
	if err != nil {
		panic("failed to request source")
	}

	respData, err := ioutil.ReadAll(resp.Body)

	return string(respData), nil
}