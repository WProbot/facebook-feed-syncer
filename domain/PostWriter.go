package domain

import (
	"github.com/jinzhu/gorm"
	"github.com/geminikim/FbToWp/config"
)

func WritePost(post Post) {
	conf := config.Config
	db, err := gorm.Open("mysql", conf.Database.User + ":" + conf.Database.Password +
					"@tcp(" + conf.Database.Ip + ":" + conf.Database.Port + ")/" +
					conf.Database.Schema + "?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	db.Create(&post)
	postCategory := PostCategory{ObjectId:post.Id, TermTaxonomyId : conf.ShareCategoryId}
	db.Create(&postCategory)
	defer db.Close()
}
