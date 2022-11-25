package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type configStruct struct {
	Token     string `json:"Token"`
	BotPrefix string `json:"BotPrefix"`
}

var (
	Token     string
	BotPrefix string
	Config    *configStruct
)

func ReadConfig() error {
	fmt.Println("reading config file ...")

	file, err := ioutil.ReadFile("././config.json")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println(string(file))

	err = json.Unmarshal(file, &Config)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	Token = Config.Token
	BotPrefix = Config.BotPrefix
	return nil
}
