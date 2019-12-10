package models

//type Category struct {
//BaseModelSeq
//	Name          string    `gorm:"INDEX, NOT NULL"`
//	Parent        *Category `json:"parent_category"`
//	Description   *string
//	ParentGroupID uint       `json:"parent_group_id"`
//	Children      []Category `json:"children_categories"`
// }

type Category struct {
	BaseModelSoftDelete
	Name        string `gorm:"INDEX, NOT NULL"`
	Description *string
	Left        int32
	Right       int32
	Depth       int32
	Sub         []*Category
	ParentID    *string
	Parent      *Category
}

func (o *ORM) GetCategoryParent(parent string) (*Category, error) {
	db := o.DB.New()
	up := &Category{}
	db = db.Set("gorm:auto_preload", true)
	if err := db.Where("id = ?", parent).First(up).Error; err != nil {
		return nil, err
	}

	return up, nil
}

func (o *ORM) FindCategory(category string) (*Category, error) {
	db := o.DB.New()
	up := &Category{}
	db = db.Set("gorm:auto_preload", true)
	if err := db.Where("id = ?", category).Find(up).Error; err != nil {
		return nil, err
	}
	return up, nil
}

func (o *ORM) FindRootCategories() ([]*Category, error) {
	db := o.DB.New()
	up := []*Category{}
	db = db.Set("gorm:auto_preload", true)
	if err := db.Where("parent_id IS NULL").Find(&up).Error; err != nil {
		return nil, err
	}

	return up, nil

}

func (o *ORM) FindSubCategories(category string) ([]*Category, error) {
	db := o.DB.New()
	up := []*Category{}
	db = db.Set("gorm:auto_preload", true)
	if err := db.Where("parent_id = ?", category).Find(&up).Error; err != nil {
		return nil, err
	}

	return up, nil
}
