package models

import (
	"oe/pkg/seed"

	"github.com/jinzhu/gorm"
)

// func CreateUser(db *gorm.DB, name string, age int) error {
// 	return db.Create(&User{Name: name, Age: age}).Error
// }

func CreatePermission(db *gorm.DB, tag string, description string) error {
	return db.Create(&Permission{Tag: tag, Description: description}).Error
}

func CreateDefaultRole(db *gorm.DB, name string, description string) error {
	p := &Permission{}
	db = db.Model(p).Where("tag = ?", "Seller.Create").Find(&p)
	return db.Create(&Role{Name: name, Description: description, Permissions: []Permission{*p}}).Error
}

func CreateDefaultSellerRole(db *gorm.DB, name string, description string) error {

	p := []Permission{}
	db = db.Model(p).Where("tag = ?", "Product.Create").Or("tag = ?", "Product.Read").Find(&p)

	return db.Create(&Role{
		Name:        name,
		Description: description,
		Permissions: p,
	}).Error
}

func All() []seed.Seed {
	return []seed.Seed{
		seed.Seed{
			Name: "CreateDefault",
			Run: func(db *gorm.DB) error {
				return CreatePermission(db, "Seller.Create", "Can create an new Seller")
			},
		},

		seed.Seed{
			Name: "CreateProductRead",
			Run: func(db *gorm.DB) error {
				return CreatePermission(db, "Product.Read", "Can read products")
			},
		},

		seed.Seed{
			Name: "CreateProductCreate",
			Run: func(db *gorm.DB) error {
				return CreatePermission(db, "Product.Create", "Can create an new Product")
			},
		},

		seed.Seed{
			Name: "CreateDefaultRole",
			Run: func(db *gorm.DB) error {
				return CreateDefaultRole(db, "User", "Create default user's")
			},
		},

		seed.Seed{
			Name: "CreateDefaultSellerRole",
			Run: func(db *gorm.DB) error {
				return CreateDefaultSellerRole(db, "Seller", "Create default seller")
			},
		},
	}
}
