package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"./websocket"
)

func GetChannel(shop string) []string {
	var ALL_CHANNELS []string
	res, err := http.Get("http://live.66boss.com/api/queryall?shop=" + shop)
	if err != nil {
		return ALL_CHANNELS
	}
	detail, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return ALL_CHANNELS
	}
	var channel interface{}
	json.Unmarshal(detail, &channel)
	data := channel.(map[string]interface{})
	data1 := data["live"]
	data2 := data1.([]interface{})
	for _, real := range data2 {
		ok := real.(map[string]interface{})
		liveid := ok["liveid"].(string)
		ALL_CHANNELS = append(ALL_CHANNELS, liveid)
	}
	return ALL_CHANNELS
}

var origin = "http://live.66boss.com/"
var request = "ws://live.66boss.com:8080/ws?"

func robot(liveid string, name string, userid string) {
	v := url.Values{}
	v.Set("liveid", liveid)
	v.Set("name", name)
	v.Set("userid", userid)
	rl := request + v.Encode()
	sleep := time.Duration(rand.Intn(10))
	time.Sleep(sleep * time.Second)
	ws, err := websocket.Dial(rl, "", origin)
	if err != nil {
		log.Fatal(err)
	}
	times := 0
	quit := rand.Intn(60)
	for ; times < 60; times++ {

		time.Sleep(1000 * time.Millisecond)
		message := []byte(WORDS[rand.Intn(len(WORDS))])
		_, err = ws.Write(message)
		if err != nil {
			log.Fatal(err)
		}

		var msg = make([]byte, 512)
		_, err = ws.Read(msg)
		if err != nil {
			log.Fatal(err)
		}
		sleep := time.Duration(rand.Intn(30))
		time.Sleep(sleep * time.Second)
		fmt.Printf("%d %d\n", times, quit)
		if times == quit {
			break
		}
	}
}

var WORDS = []string{
	"this is what i want to see",
	"awesome",
	"这么丑还直播",
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
	"走了 没意思",
	"真开眼界，带带我",
}

func main() {
	rand.Seed(time.Now().Unix()) // Try changing this number!
	for {
		ALL_CHANNELS := GetChannel("0")
		for _, liveid := range ALL_CHANNELS {
			for i := 0; i < 10; i++ { //10个机器人
				go func(liveid string) {
					fmt.Println(liveid)
					which := rand.Intn(1000)
					fmt.Println(which)
					name := "哈哈" + strconv.Itoa(which)
					ID := "123" + strconv.Itoa(which)
					robot(liveid, name, ID)
				}(liveid)
			}
		}
		ALL_CHANNELS = GetChannel("1")
		for _, liveid := range ALL_CHANNELS {
			for i := 0; i < 10; i++ { //10个机器人
				go func(liveid string) {
					fmt.Println(liveid)
					which := rand.Intn(1000)
					fmt.Println(which)
					name := "哈哈" + strconv.Itoa(which)
					ID := "123" + strconv.Itoa(which)
					robot(liveid, name, ID)
				}(liveid)
			}
		}
		time.Sleep(60 * time.Second)
	}
}
