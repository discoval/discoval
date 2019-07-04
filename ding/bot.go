package ding

import (
	"net/http"
	"strings"
)

const WebHook = "https://oapi.dingtalk.com/robot/send?access_token="

/**
发送@一个或多个人的机器人消息
params:
	msg      : string          发送通知的内容
	token	 : string          机器人对应的webhook地址中的token
	at		 : ...string	   需要@的人的手机号
*/
func SendMsg(msg, token string, at ...string) error {
	content := `{
   		"msgtype": "text",
   		"text": {
   		    "content": "` + strings.Replace(msg, "\"", "\\\"", -1) + `"
   		},
   		"at": {
   		   "atMobiles": [
   		       ` + formatMobiles(at) + `
   		   ],
   		   "isAtAll": false
   		}
	}`
	return send(content, token)
}

/**
发送@所有人的机器人消息
params:
	msg		: string 	发送通知的内容
	token	: string 	机器人对应的webhook地址中的token
*/
func SendMsgAll(msg, token string) error {
	content := `{
   		"msgtype": "text",
   		"text": {
   		    "content": "` + strings.Replace(msg, "\"", "\\\"", -1) + `"
   		},
   		"at": {
   		   "atMobiles": [],
   		   "isAtAll": true
   		}
	}`
	return send(content, token)
}

/**
发送机器人消息post请求
params:
	content	: 请求体
	token	: 机器人对应的webhook地址中的token
*/
func send(content, token string) error {
	//创建一个请求
	req, err := http.NewRequest("POST", WebHook+token, strings.NewReader(content))
	if err != nil {
		return err
	}

	client := &http.Client{}
	//设置请求头
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	//发送请求
	resp, err := client.Do(req)
	//关闭请求
	defer resp.Body.Close()

	if err != nil {
		return err
	}
	return nil
}

/**
格式化mobiles信息
*/
func formatMobiles(mobiles interface{}) string {
	switch mobiles.(type) {
	case string:
		return mobiles.(string)
	case []string:
		s := ""
		for k, v := range mobiles.([]string) {
			if k != 0 {
				s += `,`
			}

			s += `"` + v + `"`
		}
		return s
	default:
		return ""
	}
}
