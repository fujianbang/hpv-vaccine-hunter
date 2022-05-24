package main

import (
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
	"hpv-vaccine-hunter/api"
)

// loadConfig 读取配置文件
func loadConfig() {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

func main() {
	loadConfig()

	ShowMemberInfo()

	client := api.NewClient()

	// SecondKillList(client)
	// MemberList(client)
	// CheckStock(client)
	// Subscribe(client)

	// Run(func() {
	// MemberList(client)

	// 抢票
	Subscribe(client)
	// })
}

func Run(f func()) {
	ticker := time.NewTicker(100 * time.Millisecond) // 100毫秒间隔抢票

	for {
		t := <-ticker.C
		log.Println(t)

		go f()
	}
}

func Subscribe(c *api.Client) {
	target, memberId, idCard := viper.GetInt("target"), viper.GetString("member_id"), viper.GetString("id_card")

	result, err := c.Subscribe(target, memberId, idCard)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(result)
}

// ShowMemberInfo 从viper配置环境中读取待接种者信息
func ShowMemberInfo() {
	id := viper.GetString("member_id")
	id_card := viper.GetString("id_card")

	log.Printf("待接种者ID：%s，姓名：%s", id, id_card)
}

func SecondKillList(c *api.Client) {
	data, err := c.GetSecondKillList()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(data)
}

func MemberList(c *api.Client) {
	memberList, err := c.GetMemberList()
	if err != nil {
		log.Fatalln(err)
	}
	for _, member := range memberList {
		log.Println(member)
	}
}

func CheckStock(c *api.Client) {
	data, err := c.CheckStock(2151)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(data)

	// 核对服务器时间
	t := time.UnixMilli(data.ServerTimestamp)
	log.Println("时间差：", t.Sub(time.Now()))
}
