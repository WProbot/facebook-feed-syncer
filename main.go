package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/geminikim/FbToWp/facebook"
	"github.com/geminikim/FbToWp/time"
	"github.com/geminikim/FbToWp/domain"
	"github.com/geminikim/FbToWp/config"
	"fmt"
)

func main() {
	config.Init()

	syncTime := time.GetSyncTime()
	fmt.Println("Start SyncTime : " + syncTime)

	feeds := facebook.GetFeeds(syncTime)

	for _, feed := range feeds {
		syncTime = feed.Created_time
		if(feed.IsSyncFeed()) {
			post := feed.ToPost()
			domain.WritePost(post)
		}
	}

	fmt.Println("End SyncTime : " + syncTime)
	time.SetSyncTime(syncTime)

}