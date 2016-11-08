package domain

type Post struct {
	Id	int64 `gorm:"column:ID;AUTO_INCREMENT;primary_key"`
	PostAuthor	int64 `gorm:"column:post_author"`
	PostDate	string `gorm:"column:post_date"`
	PostDateGmt	string `gorm:"column:post_date_gmt"`
	PostContent	string `gorm:"column:post_content;type:LONGTEXT;"`
	PostTitle	string `gorm:"column:post_title;type:TEXT"`
	PostStatus	string `gorm:"column:post_status;type:varchar(20);"`
	CommentStatus	string `gorm:"column:comment_status;type:varchar(20);"`
	PingStatus	string `gorm:"column:ping_status;type:varchar(20);"`
	PostPassword	string `gorm:"column:post_password;type:varchar(20);"`
	PostModified	string `gorm:"column:post_modified"`
	PostModifiedGmt	string `gorm:"column:post_modified_gmt"`
	Guid	string `gorm:"column:guid;type:varchar(255);"`
	PostMimeType	string `gorm:"column:post_mime_type;type:varchar(100);"`
}

func (Post) TableName() string {
	return "wp_posts"
}