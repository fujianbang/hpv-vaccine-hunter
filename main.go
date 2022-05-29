package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

	// 待抢疫苗信息
	vaccineList, err := client.GetSecondKillList()
	if err != nil {
		log.Fatalln(err)
	}

	for i, item := range vaccineList {
		fmt.Printf("[%d] %d: %s，开始时间: %s, 地址: %s\n",
			i+1,
			item.Id, item.Name, item.StartTime, item.Address)
	}

	inputReader := bufio.NewReader(os.Stdin)
	fmt.Printf("请输入待抢的疫苗编号：")
	input, err := inputReader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	targetId, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		log.Printf("疫苗ID不合法 [%s]\n", err)
	}

	if targetId <= 0 || targetId > len(vaccineList) {
		log.Printf("疫苗ID不合法 [%d]\n", targetId)
	}

	target := vaccineList[targetId-1]
	fmt.Printf("目标疫苗信息(%d-%s-%s)\n", target.Id, target.Name, target.StartTime)

	targetTime, err := time.ParseInLocation("2006-01-02 15:04:05", target.StartTime, time.Local)
	if err != nil {
		log.Fatalf("时间解析失败 [%s]", err.Error())
	}
	// 获取提前时间
	earlyStartMs := viper.GetInt("early_start")
	// 时间调整对齐
	targetTime = targetTime.Add(-time.Duration(earlyStartMs) * time.Millisecond)
	remaining := targetTime.Sub(time.Now())
	fmt.Printf("目标时间: %v，倒计时：%v\n", targetTime, remaining)

	if remaining.Nanoseconds() < 0 {
		log.Fatalln("抢票已过期")
	}

	timer := time.NewTimer(remaining)
	select {
	case <-timer.C:
		fmt.Println("-----------------------------------------------------------")
		fmt.Printf("任务开始，执行抢票: %s\n", time.Now())

		Run(10, func() {
			// 抢疫苗
			Subscribe(client)
		})
	}
	time.Sleep(10 * time.Second)
}

func Run(times int, f func()) {
	ticker := time.NewTicker(100 * time.Millisecond) // 100毫秒间隔抢票

	var counter = 0
	for {
		t := <-ticker.C
		log.Println(t)

		go f()

		counter++

		if counter >= times {
			break
		}
	}
}

func Subscribe(c *api.Client) {
	target, memberId, idCard := viper.GetInt("target"), viper.GetString("member_id"), viper.GetString("id_card")

	result, err := c.Subscribe(target, memberId, idCard)
	if err != nil {
		log.Println(err)
	}
	log.Println(result)
}

// ShowMemberInfo 从viper配置环境中读取待接种者信息
func ShowMemberInfo() {
	id := viper.GetString("member_id")
	id_card := viper.GetString("id_card")

	log.Printf("待接种者ID：%s，身份证：%s", id, id_card)
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
	data, err := c.CheckStock(2170)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(data)

	// 核对服务器时间
	t := time.UnixMilli(data.ServerTimestamp)
	log.Println("时间差：", t.Sub(time.Now()))
}
