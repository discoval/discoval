package http

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

func Put(url string, data interface{}, contentType string, headers map[string]string, insecureSkipVerify bool) (content string) {
	jsonStr, _ := json.Marshal(data)
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonStr))
	req.Header.Add("content-type", contentType)
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()

	//跳过证书验证
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: insecureSkipVerify},
	}

	client := &http.Client{
		Timeout: 5 * time.Second,
		Transport: tr,
	}
	resp, error := client.Do(req)
	if error != nil {
		panic(error)
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	content = string(result)
	return
}