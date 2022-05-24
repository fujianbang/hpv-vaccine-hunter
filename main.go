package main

import (
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
	"hpv-vaccine-hunter/api"
)

func readConfig() (string, string) {
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

	tk, cookie := viper.GetString("tk"), viper.GetString("cookie")
	log.Printf("tk: %s, cookie: %s\n", tk, cookie)
	return tk, cookie
}

func main() {
	client := api.NewClient(readConfig())

	// aa, err := client.GetSecondKillList()
	// aa, err := client.GetMemberList()
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// log.Println(aa)

	CheckStock(client)
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
