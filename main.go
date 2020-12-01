package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/login", login)
	http.HandleFunc("/request_github", requestGithub)
	fmt.Println("url: http://localhost:8080 ")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}

func requestGithub(writer http.ResponseWriter, request *http.Request) {
	http.Get("https://api.github.com/users/jiangkang")
	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodGet, "https://api.github.com/users/jiangkang", nil)
	req.Header.Add("Accept", "application/vnd.github.v3+json")
	response, _ := client.Do(req)
	body, _ := ioutil.ReadAll(response.Body)
	writer.Write(body)
}

func login(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		t, _ := template.ParseFiles("html/login.html")
		t.Execute(writer, nil)
	} else if request.Method == http.MethodPost {
		request.ParseForm()
		username := request.Form.Get("username")
		password := request.Form.Get("password")
		fmt.Fprintf(writer, "登陆成功：username: %s \n, password: %s", username, password)
	}
}

// 默认处理
func home(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		t, _ := template.ParseFiles("html/index.html")
		t.Execute(writer, nil)
	}
}
