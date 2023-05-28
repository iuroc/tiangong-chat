package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/index.html")
	})
	http.HandleFunc("/api/login", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		phone := r.PostForm.Get("phone")
		password := r.PostForm.Get("password")
		if phone == "" || password == "" {
			w.Write([]byte("请求参数不能为空"))
			return
		}
		url := "https://neice.tiangong.cn/api/v1/user/login"
		type Value struct {
			Phone    string `json:"phone"`
			Password string `json:"passwd"`
		}
		type Data struct {
			Data Value `json:"data"`
		}
		jsonData := Data{
			Data: Value{
				Phone:    phone,
				Password: password,
			},
		}
		jsonStr, _ := json.Marshal(jsonData)
		payload := strings.NewReader(string(jsonStr))
		req, _ := http.NewRequest("POST", url, payload)
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return
		}
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))
	})
	http.HandleFunc("/api/test", func(w http.ResponseWriter, r *http.Request) {
		httpClient := &http.Client{}
		result, _ := Ajax(AjaxOption{
			Url:    "https://neice.tiangong.cn/api/v1/user/login",
			Method: "POST",
			Data:   "{\"data\":{\"phone\":\"18870142713\",\"passwd\":\"MyApee@2002\"}}",
		}, httpClient)
		fmt.Println(result)
	})
	fmt.Println("http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

// Ajax 请求函数选项
type AjaxOption struct {
	// 请求地址
	Url string
	// 请求方式
	Method string
	// POST 数据
	Data string
	// 请求头
	Header map[string]string
}

// Ajax 请求函数
func Ajax(ajaxOption AjaxOption, client *http.Client) (string, error) {
	// 获取请求方法并将其转换为大写，默认为 GET
	method := strings.ToUpper(ajaxOption.Method)
	if method == "" {
		method = "GET"
	}

	url := ajaxOption.Url
	data := ajaxOption.Data
	header := ajaxOption.Header

	// 创建包含请求数据的字符串读取器
	payload := strings.NewReader(data)

	// 创建 HTTP 请求
	request, err := http.NewRequest(method, url, payload)
	if err != nil {
		return "", err
	}

	// 设置请求头
	for key, value := range header {
		request.Header.Set(key, value)
	}

	// 发送 HTTP 请求并获取响应
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	// 读取响应体
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
