package main

import (
	"crypto/md5"
	"fmt"
	"github.com/nutsdo/go-next/cmd"
	"github.com/urfave/cli"
	"log"
	"os"
	"sort"
	"strings"
)


func main(){
	app:=cli.NewApp()
	app.Name="go-next"
	app.Usage="Go short video service"

	app.Commands = []cli.Command{
		{
			Name:    "start",
			Aliases: []string{"s"},
			Usage:   "start server...",
			Action:  func(c *cli.Context) error {
				cmd.RunServer()
				fmt.Println("started server...")
				return nil
			},
		},
		{
			Name:    "migrate",
			Aliases: []string{"m"},
			Usage:   "migrate data",
			Action:  func(c *cli.Context) error {
				fmt.Println("completed task: ", c.Args().First())
				cmd.Migrate()
				return nil
			},
		},
	}


	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}


func test(){
	var headerMaps map[string]string=map[string]string{
		//"Vername":"5.6.1",
		"Nonce":"mt9istfT",
		"Apiver":"4",
		"Timestamp":"1561082829813",
		//"Dev":"31DDB46F-4949-4C40-A0FE-553451503E4F",
		"Func":"0x0007",
		//"Cookie":"",
		//"Vercode":"937fdfb57dcaa3c826b36279ac67e8de",
		"Ver":"28",
		//"Funcver":"3",
		//"Model":"phone[+]iPhone 6s Plus[+]12.2",
		"Sign":"2.7",
	}
	str:=GenerateSignMd5(headerMaps)
	fmt.Println(str)
}


func GenerateSignMd5(params map[string]string) string {

	//params["app_key"] = tauth.Appkey

	//map转换为数组
	var keys []string
	for k, v := range params {
		if v != "" {
			keys = append(keys, k)
		}
	}
	var str string
	//按ASCII排序
	sort.Strings(keys)
	for _, k := range keys {
		str += k + params[k]
	}
	//str+="shengqianyoudao"
	signByte := []byte(str)
	fmt.Println(str)
	//生成签名
	sign := fmt.Sprintf("%x", md5.Sum(signByte))
	return strings.ToLower(sign)
}
