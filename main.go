package main

import (
	"fmt"
	"log"
	"os"

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

	aa, err := client.GetMemberList()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(aa)
}
