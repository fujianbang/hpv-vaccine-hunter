package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/spf13/viper"
	"hpv-vaccine-hunter/model"
	"hpv-vaccine-hunter/tools"
)

const (
	seckillUrl = "/seckill/seckill/list.do?offset=0&limit=10&regionCode=5101"
)

type Client struct {
	auth Auth // 认证
}

func NewClient() *Client {
	tk, cookie := viper.GetString("tk"), viper.GetString("cookie")
	log.Printf("tk: %s, cookie: %s\n", tk, cookie)

	return &Client{
		auth: Auth{
			TK:     tk,
			Cookie: cookie,
		},
	}
}

func (c *Client) get(path string, params map[string]string, header map[string]string) ([]byte, error) {
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

	// 请求头
	if header != nil && len(header) > 0 {
		for k, v := range header {
			req.Header.Set(k, v)
		}
	}

	// 构造参数
	if params != nil {
		q := req.URL.Query()
		for k, v := range params {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

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
	data, err := c.get(seckillUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	var vaccines []model.VaccineItem
	if err := json.Unmarshal(data, &vaccines); err != nil {
		return nil, err
	}

	return vaccines, nil
}

// CheckStock 检查指定秒杀剩余疫苗数
func (c *Client) CheckStock(id int) (*model.CheckStockResult, error) {
	url := fmt.Sprintf("/seckill/seckill/checkstock2.do?id=%d", id)
	data, err := c.get(url, nil, nil)
	if err != nil {
		return nil, err
	}

	var result model.CheckStockResult
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// Subscribe 抢疫苗
func (c *Client) Subscribe(seckillId int, linkmanId string, idCard string) (*model.SubscribeResult, error) {
	log.Printf("秒杀ID: %d, 接种者ID: %s, 身份证: %s\n", seckillId, linkmanId, idCard)

	query := map[string]string{
		"seckillId":    strconv.Itoa(seckillId),
		"vaccineIndex": "1",
		"linkmanId":    linkmanId,
		"idCardNo":     idCard,
	}

	st := time.Now().UnixMilli()
	secret := tools.Md5Hex(fmt.Sprintf("%d%s%d", seckillId, linkmanId, st))
	log.Printf("加密参数结果：%s (st:%d)", secret, st)
	header := map[string]string{
		"ecc-hs": secret,
	}

	data, err := c.get("/seckill/seckill/subscribe.do", query, header)
	if err != nil {
		return nil, err
	}

	var result model.SubscribeResult
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return nil, nil
}

func (c *Client) GetMemberList() ([]model.Member, error) {
	data, err := c.get("/seckill/linkman/findByUserId.do", nil, nil)
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
