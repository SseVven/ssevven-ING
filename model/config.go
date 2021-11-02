package model

import (
	"bufio"
	"encoding/json"
	"os"
)

type Config struct {
	AppName      string `json:"app_name"`
	AppMode      string `json:"app_mode"`
	AppHost      string `json:"app_host"`
	AppPort      string `json:"app_port"`
	HttpProtocol string `json:"http_protocol"`
	Domain       string `json:"domain"`
}

var Cfg *Config = nil

func ParseConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		panic(err) //初始化出错，程序直接退出
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	decoder := json.NewDecoder(reader)
	err = decoder.Decode(&Cfg)
	if err != nil {
		return nil, err
	}
	return Cfg, nil
}
