package models

type Category struct {
	BaseModelSeq
	Name          string     `gorm:"INDEX, NOT NULL"`
	Parent        *Category  `json:"parent_category"`
	ParentGroupID uint       `json:"parent_group_id"`
	Children      []Category `json:"children_categories"`
}
