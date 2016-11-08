package facebook

import (
	"io/ioutil"
	"net/http"
	"encoding/json"
	"github.com/geminikim/FbToWp/config"
)

func GetFeeds(since string) ([]Feed){
	conf := config.Config
	response := get("https://graph.facebook.com/v2.3/" + conf.Facebook.UserId + "/feed?" + getAccessToken() +
				"&limit=50&since=" + since)
	return getFeeds(response)
}

func getFeeds(response []byte) ([]Feed){
	var feeds Feeds
	json.Unmarshal(response, &feeds)
	return feeds.Feeds
}

func getAccessToken() (string){
	conf := config.Config
	return string(get("https://graph.facebook.com/oauth/access_token?client_id=" + conf.Facebook.AppId +
				"&client_secret=" + conf.Facebook.AppSecret +
				"&grant_type=client_credentials"))
}

func get(url string) ([]byte){
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body
}
