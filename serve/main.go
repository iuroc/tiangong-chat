package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fs := http.FileServer(http.Dir("../static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../static/index.html")
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
