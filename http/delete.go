package http

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"time"
)

func Delete(url string, contentType string, headers map[string]string, insecureSkipVerify bool) (content string) {
	req, err := http.NewRequest("DELETE", url, nil)
	req.Header.Add("content-type", contentType)
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	if err != nil {
		panic(err)
	}

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