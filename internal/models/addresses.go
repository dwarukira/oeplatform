package models

type Address struct {
	BaseModelSoftDelete
	Street   string
	User     User   `gorm:"association_autoupdate:false;association_autocreate:false"`
	UserID   string `gorm:"not null;index"`
	Number   string
	Country  string
	City     string
	PostCode string
}
