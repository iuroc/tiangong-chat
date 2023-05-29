package route

import (
	"encoding/json"
	"net/http"

	"apee.top/tiangong-chat/serve/util"
)

func LoginRoute(w http.ResponseWriter, r *http.Request, httpClient *http.Client) {
	r.ParseForm()
	phone := r.PostForm.Get("phone")
	password := r.PostForm.Get("password")
	w.Header().Set("Content-Type", "application/json")
	if phone == "" || password == "" {
		w.Write(util.MakeErr("请输入手机号和密码"))
		return
	}
	url := "https://neice.tiangong.cn/api/v1/user/login"
	type Value struct {
		Phone    string `json:"phone"`
		Password string `json:"passwd"`
	}
	type PostData struct {
		Data Value `json:"data"`
	}
	postData, _ := json.Marshal(PostData{
		Data: Value{
			Phone:    phone,
			Password: password,
		},
	})
	result, _ := util.Ajax(util.AjaxOption{
		Url:    url,
		Method: "POST",
		Data:   string(postData),
	}, httpClient)
	type Response struct {
		Code int               `json:"code"`
		Data map[string]string `json:"resp_data"`
	}
	var response Response
	json.Unmarshal([]byte(result), &response)
	token := response.Data["token"]
	code := response.Code
	if code == 200 {
		w.Write(util.MakeSuc("登录成功", token))
	} else {
		w.Write(util.MakeErr("登录失败"))
	}
}
