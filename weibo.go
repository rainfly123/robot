package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const URL = "http://live.66boss.com/weibo/square?login_user=1"

var WORDS = []string{
	"this is what i want to see",
	"awesome",
	"无聊啊 晚上一起吃饭吧",
	"doubtful",
	"顶顶顶",
	"一起看电影吧",
	"你胸大你先说，哈",
	"你那么漂亮，说什么都对",
	"给力，楼主",
	"prefect",
	"我不以为这样，你说的片面",
	"太牛比啦",
	"路过",
	"留个脚印",
	"你好啊 楼主，好久不见",
	"I think so",
	":)---:)",
	"哈哈",
	"全是套路",
	"嘿嘿、老大想你哦",
	"什么啊 这是",
	"这个不错哦 马上转发",
	"约吗?今天晚上",
	"加个微信吧!",
	"整天发，拉黑你!",
	"牛逼啊 服了",
	"我要成为你的粉丝",
	"加我QQ 有事情找你",
	"好玩吗? 还是发红包吧",
	"楼主 快发红包",
	"太绝了",
	"求关注",
	"不看会后悔",
	"楼主你是90后吗",
	"看看肯定有收获",
	"真开眼界，带带我",
}

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
	rand.Seed(time.Now().Unix()) // Try changing this number!
	ReadUser()
	GetWeibo()
	for _, weibo := range ALL_WEIBOS {
		if weibo.Supports < 10 {
			fmt.Println(weibo)
			num := rand.Intn(10)
			for j := 0; j < num; j++ {
				user := ALL_USERS[rand.Intn(len(ALL_USERS))]
				baseurl := "http://live.66boss.com/weibo/support?"
				url := fmt.Sprintf("%sweiboid=%d&login_user=%s", baseurl, weibo.Weiboid, user.ID)
				res, _ := http.Get(url)
				//fmt.Println(url)
				res.Body.Close()
			}
		}
		if weibo.Comments < 3 {
			num := rand.Intn(8)
			for j := 0; j < num; j++ {
				fmt.Println(weibo)
				user := ALL_USERS[rand.Intn(len(ALL_USERS))]
				baseurl := "http://live.66boss.com/weibo/comment?"

				v := url.Values{}
				v.Set("comment", WORDS[rand.Intn(len(WORDS))])
				v.Set("weiboid", strconv.Itoa(weibo.Weiboid))
				v.Set("login_user", user.ID)
				url := baseurl + v.Encode()
				res, _ := http.Get(url)
				//detail, _ := ioutil.ReadAll(res.Body)
				//fmt.Println(string(detail))
				res.Body.Close()
			}
		}
	}
	num := rand.Intn(len(ALL_WEIBOS))
	weibo := ALL_WEIBOS[num]
	user := ALL_USERS[rand.Intn(len(ALL_USERS))]
	baseurl := "http://live.66boss.com/weibo/forward?"

	v := url.Values{}
	v.Set("msg", WORDS[rand.Intn(len(WORDS))])
	v.Set("origin", strconv.Itoa(weibo.Weiboid))
	v.Set("login_user", user.ID)
	url := baseurl + v.Encode()
	res, _ := http.Get(url)
	//detail, _ := ioutil.ReadAll(res.Body)
	//fmt.Println(string(detail))
	res.Body.Close()

}
