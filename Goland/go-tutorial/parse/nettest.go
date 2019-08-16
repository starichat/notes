package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	client := &http.Client{}

	req, err := http.NewRequest("POST", "http://www.maimaiche.com/loginRegister/login.do",
		strings.NewReader("mobile=xxxxxxxxx&isRemberPwd=1"))
	if err != nil {
		log.Println(err)
		return
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(resp.Header.Get("Content-Type")) //application/json;charset=UTF-8

	type Result struct {
		Msg    string
		Status string
		Obj    string
	}

	result := &Result{}
	json.Unmarshal(body, result) //解析json字符串

	if result.Status == "1" {
		fmt.Println(result.Msg)
	} else {
		fmt.Println("login error")
	}
	fmt.Println(result)

	postParam := url.Values{
		"mobile":      {"xxxxxx"},
		"isRemberPwd": {"1"},
	}

	resp, err := http.PostForm("http://www.maimaiche.com/loginRegister/login.do", postParam)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(body))
}
