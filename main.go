package main

// Created by GaoFeng
// Get the current number of problems
// 2019-06-10
// https://www.qiansw.com/how-to-use-zcate-to-receive-zabbix-alarm-messages.html

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var (
	// Your zabbix info
	zUser   = "read"
	zPasswd = "read"
	zURL    = "http://www.qiansw.com/api_jsonrpc.php"
)

// ZabbixResult ...
type ZabbixResult struct {
	Jsonrpc string `json:"jsonrpc"`
	Result  string `json:"result"`
	Error   struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Data    string `json:"data"`
	} `json:"error"`
	ID int `json:"id"`
}

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	var _user, _passwd, _url string
	flag.StringVar(&_user, "u", "", "Your zabbix user")
	flag.StringVar(&_passwd, "p", "", "Your zabbix password")
	flag.StringVar(&_url, "url", "", "Your zabbix api url")
	flag.Parse()
	if _user != "" && _passwd != "" && _url != "" {
		zUser = _user
		zPasswd = _passwd
		zURL = _url
	}
}

func main() {
	fmt.Print(getCountProblems(getZabbixToken()))
}

func getCountProblems(token string) string {
	resp, err := http.Post(zURL,
		"application/json",
		strings.NewReader(fmt.Sprintf(`{
			"jsonrpc": "2.0",
			"method": "trigger.get",
			"params": {
				"filter": {
					"value": 1
				},
				"countOutput": 1,
				"only_true": "1",
				"monitored": "1",
				"active": "1"
			},
			"auth": "%v",
			"id": 1
		}`, token)))
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	if resp.StatusCode != 200 {
		log.Fatalln(string(body))
	}

	var countResult ZabbixResult

	err = json.Unmarshal(body, &countResult)

	if err != nil {
		log.Fatalln(err)
	}

	// log.Println(countResult)

	if countResult.Error.Message != "" {
		log.Fatal(countResult.Error)
	}

	return countResult.Result
}

// getZabbixToken Get zabbix Token
func getZabbixToken() string {
	resp, err := http.Post(zURL,
		"application/json",
		strings.NewReader(fmt.Sprintf(`{
			"jsonrpc": "2.0",
			"method": "user.login",
			"params": {
				"user": "%v",
				"password": "%v"
			},
			"id": 1
		}`, zUser, zPasswd)))
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	if resp.StatusCode != 200 {
		log.Fatal(string(body))
	}

	var tokenResult ZabbixResult

	err = json.Unmarshal(body, &tokenResult)
	if err != nil {
		log.Println(string(body))
		log.Fatalln(err)
	}

	if tokenResult.Error.Message != "" {
		log.Fatalln(tokenResult.Error)
	}

	return tokenResult.Result
}
