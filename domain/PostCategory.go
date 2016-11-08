package domain

type PostCategory struct {
	ObjectId	int64 `gorm:"object_id"`
	TermTaxonomyId	int64 `gorm:"term_taxonomy_id"`
}

func (PostCategory) TableName() string {
	return "wp_term_relationships"
}