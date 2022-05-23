package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"hpv-vaccine-hunter/model"
)

const (
	seckillUrl   = "/seckill/seckill/list.do?offset=0&limit=10&regionCode=5101"
	getMemberUrl = "/seckill/linkman/findByUserId.do"
)

type Client struct {
	auth Auth // 认证
}

func NewClient(tk, cookie string) *Client {
	return &Client{
		auth: Auth{
			TK:     tk,
			Cookie: cookie,
		},
	}
}

func (c *Client) get(path string) ([]byte, error) {
	url := "https://miaomiao.scmttec.com" + path

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Host", "miaomiao.scmttec.com")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("tk", c.auth.TK)
	req.Header.Set("Cookie", c.auth.Cookie)
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("Accept-Encoding", "gzip,compress,br,deflate")
	req.Header.Set("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 15_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 MicroMessenger/8.0.20(0x18001435) NetType/WIFI Language/zh_CN")
	req.Header.Set("Referer", "https://servicewechat.com/wxff8cad2e9bf18719/27/page-frame.html")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data CommonResponse
	if err := json.Unmarshal(bytes, &data); err != nil {
		return nil, err
	}

	log.Printf("get http response: %v\n", data)

	// 错误处理
	if !data.Ok {
		log.Printf("[error] get http resonse: %v\n", data)
		return nil, nil
	}

	result, err := json.Marshal(data.Data)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetSecondKillList 获取秒杀列表
func (c *Client) GetSecondKillList() ([]model.VaccineItem, error) {
	data, err := c.get(seckillUrl)
	if err != nil {
		return nil, err
	}

	var vaccines []model.VaccineItem
	if err := json.Unmarshal(data, &vaccines); err != nil {
		return nil, err
	}

	return vaccines, nil
}

func (c *Client) GetMemberList() ([]model.Member, error) {
	data, err := c.get(getMemberUrl)
	if err != nil {
		return nil, err
	}

	var members []model.Member
	if err := json.Unmarshal(data, &members); err != nil {
		return nil, err
	}

	return members, nil
}

type CommonResponse struct {
	Data   interface{} `json:"data"`
	Status string      `json:"code"` // 0000 ok
	Ok     bool        `json:"ok"`
	NotOk  bool        `json:"notOk"`
	Msg    string      `json:"msg"` // 错误信息
}
