package httpReq

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type IRequest interface {
	Get(wechatUrl string) (map[string]interface{}, error)
	PostJson(wechatUrl string, data interface{}) (map[string]interface{}, error)
	PostForm(wechatUrl string, data url.Values) (map[string]interface{}, error)
}
//contentType常量
const (
	JSON = "application/json"
	HTML = "text/html"
	FORM = "application/x-www-form-urlencoded"
)

type Request struct {
	ContentType string
	Timeout time.Duration
}


func (r Request) Get(wechatUrl string) (map[string]interface{}, error) {
	client := &http.Client{Timeout: getTimeout(r.Timeout)}
	res, err := client.Get(wechatUrl)
	fmt.Println(err)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	mp, err := parseResponse(res)
	return mp, err
}

func (r Request) PostJson(wechatUrl string, data interface{}) (map[string]interface{}, error) {
	client := &http.Client{Timeout: getTimeout(r.Timeout)}
	jsonStr, _ := json.Marshal(data)
	contentType := JSON
	fmt.Println(contentType)
	res, err := client.Post(wechatUrl, contentType, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return parseResponse(res)
}

func (r Request) PostForm(wechatUrl string, data url.Values) (map[string]interface{}, error) {
	resp, err := http.PostForm(wechatUrl, data)
	if err != nil {
		fmt.Printf("get request failed, err:[%s]", err.Error())
		return nil, err
	}
	defer resp.Body.Close()
	return parseResponse(resp)
}

func parseResponse(resp *http.Response) (map[string]interface{}, error) {
	var mp map[string]interface{}
	result, _ := ioutil.ReadAll(resp.Body)
	err := json.Unmarshal(result, &mp)
	if err != nil {
		return nil, err
	}
	return mp, nil
}

func getContentType(ct string) string {
	if len(ct) > 0 {
		return ct
	}
	return JSON
}

func getTimeout(timeour time.Duration) time.Duration {
	if timeour > 0 {
		return timeour
	} else {
		return 5*time.Second
	}
}





