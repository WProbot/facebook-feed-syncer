package config

import (
	"os"
	"encoding/json"
	"fmt"
)

type SyncConfig struct {
	WordpressUserId int64
	SyncFeedPrefix string
	ShareCategoryId int64
	Facebook FacebookConfig
	Database DatabaseConfig
}
type FacebookConfig struct {
	UserId string
	AppId string
	AppSecret string
}
type DatabaseConfig struct {
	User string
	Password string
	Ip string
	Port string
	Schema string
}

var Config SyncConfig

func Init() {
	file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)
	initConfig := SyncConfig{}
	err := decoder.Decode(&initConfig)
	if err != nil {
		fmt.Println("error:", err)
	}
	Config = initConfig;
}