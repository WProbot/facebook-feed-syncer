package facebook

import (
	"github.com/geminikim/FbToWp/domain"
	"github.com/geminikim/FbToWp/config"
	"strings"
	"fmt"
)

type Feed struct {
	Message      string
	Description  string
	Link         string
	Name         string
	Picture      string
	Created_time string
}

type UserMessage struct {
	Tag	string
	Title	string
	Content	string
}

func (feed Feed) IsSyncFeed() (bool) {
	return strings.HasPrefix(feed.Message, config.Config.SyncFeedPrefix)
}

func (feed Feed) ToPost() (domain.Post) {
	fmt.Println("Synced Feed : " + feed.Name)

	userMessage := feed.toUserMessage()

	return domain.Post{
		PostAuthor:config.Config.WordpressUserId,
		PostTitle: feed.getPostTitle(userMessage),
		PostContent: feed.getPostContent(userMessage),
		PostDate:feed.Created_time,
		PostDateGmt:feed.Created_time,
		PostStatus:"publish",
		CommentStatus:"open",
		PingStatus:"closed",
		PostPassword:"",
		Guid:"",
		PostModified:feed.Created_time,
		PostMimeType:"",
		PostModifiedGmt:feed.Created_time,
	}
}

func (feed Feed) getPostTitle(userMessage UserMessage) (string) {
	if isNotEmpty(userMessage.Title) {
		return userMessage.Title
	} else if isNotEmpty(feed.Name) {
		return feed.Name
	} else {
		return feed.Description[:10] + " ..."
	}
}

func (feed Feed) getPostContent(userMessage UserMessage) (string) {
	var postContent = ""
	if isNotEmpty(userMessage.Content) {
		postContent += userMessage.Content
	} else {
		postContent += feed.Description
	}
	if isNotEmpty(userMessage.Tag) {
		postContent += "<br/>TAG : " + strings.Replace(strings.TrimSpace(userMessage.Tag), ",", " / ", -1)
	}
	if isNotEmpty(feed.Created_time) {
		postContent += "<br/>CREATED TIME : " + feed.Created_time
	}
	if isNotEmpty(feed.Link) {
		postContent += "<br/>ARTICLE LINK : <a href=\"" + feed.Link + "\">" + feed.Link + "<a/>"
	}
	return postContent
}

func (feed Feed) getMessageByIndex(index int) (string) {
	message := strings.Split(feed.Message, config.Config.SyncFeedPrefix);

	if cap(message) < index {
		return ""
	}
	return message[index]
}

func (feed Feed) toUserMessage() (UserMessage) {
	return UserMessage{Tag: feed.getMessageByIndex(1), Title: feed.getMessageByIndex(2), Content: feed.getMessageByIndex(3)}
}

func isNotEmpty(str string) (bool) {
	return len(strings.TrimSpace(str)) != 0
}
