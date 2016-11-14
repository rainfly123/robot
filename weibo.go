package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
)

const URL = "http://live.66boss.com/weibo/square?login_user=1"

type Weibo struct {
	Weiboid  int
	Supports int
	Comments int
}

var ALL_WEIBOS []Weibo

func GetWeibo() {
	res, err := http.Get(URL)
	if err != nil {
		return
	}
	detail, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return
	}
	var weibo interface{}
	json.Unmarshal(detail, &weibo)
	data := weibo.(map[string]interface{})
	data1 := data["data"]
	data2 := data1.([]interface{})
	for _, real := range data2 {
		ok := real.(map[string]interface{})
		weiboid := int(ok["weiboid"].(float64))
		comments := int(ok["comments"].(float64))
		supports := int(ok["supports"].(float64))
		ALL_WEIBOS = append(ALL_WEIBOS, Weibo{Weiboid: weiboid, Comments: comments, Supports: supports})
	}
}
func main() {
	rand.Seed(42)
	ReadUser()
	GetWeibo()
	for _, weibo := range ALL_WEIBOS {
		if weibo.Supports > 10 {
			continue
		}
		user := ALL_USERS[rand.Intn(len(ALL_USERS))]
		baseurl := "http://live.66boss.com/weibo/support?"
		url := fmt.Sprintf("%sweiboid=%d&login_user=%s", baseurl, weibo.Weiboid, user.ID)
		res, _ := http.Get(url)
		fmt.Println(url)
		res.Body.Close()
	}
}
