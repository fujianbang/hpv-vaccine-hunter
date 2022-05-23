package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://miaomiao.scmttec.com/seckill/seckill/list.do?offset=0&limit=10&regionCode=5101"

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

// GetSecondKillList 获取秒杀列表
func (c *Client) GetSecondKillList() {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Host", "miaomiao.scmttec.com")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("tk", c.auth.TK)
	req.Header.Set("Cookie", c.auth.Cookie)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var data GetOrganizationResponse
	if err := json.Unmarshal(bytes, &data); err != nil {
		panic(err)
	}
	fmt.Println(data)
}

type GetOrganizationResponse struct {
	List   []Organization `json:"list"`
	Status int            `json:"status"` // 200 ok
	Msg    string         `json:"msg"`    // "102"
}

type Organization struct {
	Id           int           `json:"id"`
	Cname        string        `json:"cname"`
	Addr         string        `json:"addr"`
	SmallPic     string        `json:"SmallPic"`
	BigPic       interface{}   `json:"BigPic"`
	Lat          float64       `json:"lat"`
	Lng          float64       `json:"lng"`
	Tel          string        `json:"tel"`
	Addr2        string        `json:"addr2"`
	Province     int           `json:"province"`
	City         int           `json:"city"`
	County       int           `json:"county"`
	Sort         int           `json:"sort"`
	DistanceShow int           `json:"DistanceShow"`
	PayMent      string        `json:"PayMent"`
	IdcardLimit  bool          `json:"IdcardLimit"`
	Notice       string        `json:"notice"`
	Distance     float64       `json:"distance"`
	Tags         []interface{} `json:"tags"`
}
