package main

import (
	"flag"
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/exec"
	"service-all/library/common"
	"service-all/library/file"
	"net/url"
)

var (
	confPath string
	repo     string
)

var defaultConf = `{
  "token": "{Hash_Token}"
}`

type config struct {
	Token string
}
var Config = &config{}

func init() {
}

func main() {
	color.Cyan(`
	  *****    *****
	******************
	****   ******   ****
	****   ******   ****
	******************
	 *******   ******
	  *****   *****
	   **** *****
	     *****
	    *****
	   *****
	  ****
	`)
	/**
	 * @Author demon
	 * @Description //TODO: 配置路径
	 * @Date 2020-7-12 16:25:31
	 **/
	flag.StringVar(&confPath, "c", "config.json", "add config file")
	flag.StringVar(&repo, "clone", "", "git clone repo")
	flag.Parse()

	if confPath == "config.json" {
		// TODO: add config.js
		if file.Exists(confPath) {
			// alternatively, you can create a new viper instance.
			runtime_viper := viper.New()
			runtime_viper.SetConfigFile(confPath)

			runtime_viper.SetConfigFile(confPath)
			if err := runtime_viper.ReadInConfig(); err != nil { // Handle errors reading the config file
				log.Panic("Fatal error config file: %s \n", err)
			}

			// unmarshal config
			if err := runtime_viper.Unmarshal(&Config); err != nil {
				log.Panic("json unmarshal error: %s", err)
			}
			if repo == "" || len(repo) == 0{
				fmt.Println("Please enter your need git clone repo:")
				fmt.Scanln(&repo)
			}
			u, err := url.Parse(repo)
			if err != nil {
				panic(err)
			}

			gitpath := fmt.Sprintf("%s://%s@%s%s", u.Scheme, Config.Token, u.Host, u.Path)

			cmd := exec.Command("git", "clone", gitpath)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			if err := cmd.Run(); err != nil {
				panic(err)
			}

		} else {

			fmt.Printf("Please enter your token: ")
			fmt.Scanln(&Config.Token)
			var confContent string
			if Config.Token == "" || len(Config.Token) == 0{
				confContent = common.Replace(map[string]string{
					"{Hash_Token}":    common.RandStringRunes(64),
				}, defaultConf)
			} else {
				confContent = common.Replace(map[string]string{
					"{Hash_Token}":    Config.Token,
				}, defaultConf)
			}
			f, err := file.CreatNestedFile(confPath)
			if err != nil {
				log.Panic("create config file ror %s", err)
			}
			// 写入配置文件
			_, err = f.WriteString(confContent)
			if err != nil {
				log.Panic("create config file ror %s", err)
			}

			f.Close()

			log.Printf("create config file success")

		}
	}
}
