package models

import (
	"github.com/gofrs/uuid"
)

// Seller reps the oe seller
type Seller struct {
	BaseModelSoftDelete // We don't to actually delete the sellers, maybe audit

	// Owner of the shop
	User   User      `json:"-" gorm:"foreignkey:UserID"`
	UserID uuid.UUID `json:"-"`

	Logo string `json:"logo"`
	Name string `json:"name"`

	Address1 string `json:"address1"`
	Address2 string `json:"address2"`
	City     string `json:"city"`
	Country  string `json:"country"`
	State    string `json:"state"`
	Zip      string `json:"zip"`

	DisplayName string `json:"display_name"`
	Phone       string `json:"phone"`
	Bank        *Bank
	BankID      string

	Product []Product
}

// TableName returns the database table name for the Seller model.
// func (Seller) TableName() string {
// 	return tableName("sellers")
// }

// Bank rep a sellers bank account
type Bank struct {
	BaseModelSoftDelete
	Name              *string
	Location          *string
	Seller            *Seller `gorm:"foreignkey:SellerID"`
	SellerID          *string
	AccountHolderName *string
	BankAccount       *string
}

// TableName returns the database table name for the Seller model.
// func (Bank) TableName() string {
// 	return tableName("banks")
// }
