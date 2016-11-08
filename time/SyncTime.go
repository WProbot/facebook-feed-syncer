package time

import (
	"io/ioutil"
	"strings"
	"time"
	"os"
)

func GetSyncTime() (string){
	bs, err := ioutil.ReadFile("syncTime.txt")
	if err != nil {
		return strings.TrimSpace(time.Now().Format("2006-01-02T15:04:05"))
	}
	return strings.TrimSpace(string(bs))
}

func SetSyncTime(syncTime string) {
	file, err := os.Create("syncTime.txt")
	if err != nil {
		return
	}
	defer file.Close()
	file.WriteString(strings.Split(syncTime, "+")[0])
}