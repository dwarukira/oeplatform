// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type Authuser struct {
	User  *User   `json:"user"`
	Token *string `json:"token"`
}

type Bank struct {
	ID            string     `json:"id"`
	Name          string     `json:"name"`
	HolderName    *string    `json:"holderName"`
	AccountNumber *string    `json:"accountNumber"`
	Location      *string    `json:"location"`
	CreatedAt     time.Time  `json:"createdAt"`
	UpdatedAt     *time.Time `json:"updatedAt"`
}

type BankInput struct {
	Name          string `json:"name"`
	HolderName    string `json:"holderName"`
	AccountNumber string `json:"accountNumber"`
	Location      string `json:"location"`
}

type Permission struct {
	ID          string     `json:"id"`
	Description *string    `json:"description"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   *time.Time `json:"updatedAt"`
	Tag         string     `json:"tag"`
}

type Product struct {
	ID              string            `json:"id"`
	CreatedAt       time.Time         `json:"createdAt"`
	UpdatedAt       time.Time         `json:"updatedAt"`
	PublishedAt     *time.Time        `json:"publishedAt"`
	Seller          *Seller           `json:"seller"`
	Active          bool              `json:"active"`
	Name            string            `json:"name"`
	Slug            string            `json:"slug"`
	Code            *string           `json:"code"`
	Description     string            `json:"description"`
	DescriptionJSON string            `json:"descriptionJson"`
	Extra           *string           `json:"extra"`
	PublicationDate *time.Time        `json:"publicationDate"`
	IsPublished     bool              `json:"isPublished"`
	Variants        []*ProductVariant `json:"variants"`
	PublishedScope  *string           `json:"publishedScope"`
	Color           *string           `json:"color"`
	ColorFamily     *string           `json:"colorFamily"`
	Brand           *string           `json:"brand"`
}

type ProductCreateInput struct {
	PublicationDate *time.Time             `json:"publicationDate"`
	Description     *string                `json:"description"`
	Seller          *string                `json:"seller"`
	IsPublished     *bool                  `json:"isPublished"`
	Name            *string                `json:"name"`
	Code            *string                `json:"code"`
	Extra           *string                `json:"extra"`
	DescriptionJSON *string                `json:"descriptionJson"`
	Active          *bool                  `json:"active"`
	Color           *string                `json:"color"`
	Brand           *string                `json:"brand"`
	ColorFamily     *string                `json:"colorFamily"`
	Variants        []*ProductVariantInput `json:"variants"`
}

type ProductVariant struct {
	ID                string    `json:"id"`
	Sku               string    `json:"sku"`
	Name              string    `json:"name"`
	Taxable           *bool     `json:"taxable"`
	Barcode           *string   `json:"barcode"`
	InventoryQuantity *int      `json:"inventoryQuantity"`
	Weight            *float64  `json:"weight"`
	WeightUnit        *string   `json:"weightUnit"`
	RequiresShipping  *bool     `json:"requiresShipping"`
	Grams             *float64  `json:"grams"`
	CompareAtPrice    *float64  `json:"compareAtPrice"`
	Price             *float64  `json:"price"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
	TrackInventory    *bool     `json:"trackInventory"`
	QuantityAllocated *int      `json:"quantityAllocated"`
	Product           *Product  `json:"product"`
}

type ProductVariantInput struct {
	Sku               string   `json:"sku"`
	Name              string   `json:"name"`
	Taxable           *bool    `json:"taxable"`
	Barcode           *string  `json:"barcode"`
	InventoryQuantity *int     `json:"inventoryQuantity"`
	Weight            *float64 `json:"weight"`
	WeightUnit        *string  `json:"weightUnit"`
	RequiresShipping  *bool    `json:"requiresShipping"`
	Grams             *float64 `json:"grams"`
	CompareAtPrice    *float64 `json:"compareAtPrice"`
	Price             *float64 `json:"price"`
	TrackInventory    *bool    `json:"trackInventory"`
	QuantityAllocated *int     `json:"quantityAllocated"`
}

type Products struct {
	Count *int       `json:"count"`
	List  []*Product `json:"list"`
}

type Role struct {
	ID          string        `json:"id"`
	Name        string        `json:"name"`
	Description *string       `json:"description"`
	CreatedAt   time.Time     `json:"createdAt"`
	UpdatedAt   *time.Time    `json:"updatedAt"`
	Permissions []*Permission `json:"permissions"`
}

type RoleInput struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

type Roles struct {
	Count *int    `json:"count"`
	List  []*Role `json:"list"`
}

type Seller struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	Phone       *string    `json:"phone"`
	Website     *string    `json:"website"`
	User        *User      `json:"user"`
	Bank        *Bank      `json:"bank"`
	DisplayName *string    `json:"displayName"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   *time.Time `json:"updatedAt"`
}

type SellerInput struct {
	Name        string     `json:"name"`
	Phone       string     `json:"phone"`
	Website     *string    `json:"website"`
	DisplayName string     `json:"displayName"`
	Bank        *BankInput `json:"bank"`
}

type Sellers struct {
	Count *int      `json:"count"`
	List  []*Seller `json:"list"`
}

type TokenCreateInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	ID          string         `json:"id"`
	Email       string         `json:"email"`
	AvatarURL   *string        `json:"avatarURL"`
	Name        *string        `json:"name"`
	FirstName   *string        `json:"firstName"`
	LastName    *string        `json:"lastName"`
	NickName    *string        `json:"nickName"`
	Description *string        `json:"description"`
	Location    *string        `json:"location"`
	APIkey      *string        `json:"APIkey"`
	Profiles    []*UserProfile `json:"profiles"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   *time.Time     `json:"updatedAt"`
	LastLogin   *time.Time     `json:"lastLogin"`
	Status      *UserStatus    `json:"status"`
	Roles       []*Role        `json:"roles"`
}

type UserInput struct {
	Email       *string `json:"email"`
	Password    *string `json:"password"`
	AvatarURL   *string `json:"avatarURL"`
	DisplayName *string `json:"displayName"`
	Name        *string `json:"name"`
	FirstName   *string `json:"firstName"`
	LastName    *string `json:"lastName"`
	NickName    *string `json:"nickName"`
	Description *string `json:"description"`
	Location    *string `json:"location"`
}

type UserProfile struct {
	ID          string         `json:"id"`
	Email       string         `json:"email"`
	AvatarURL   *string        `json:"avatarURL"`
	Name        *string        `json:"name"`
	FirstName   *string        `json:"firstName"`
	LastName    *string        `json:"lastName"`
	NickName    *string        `json:"nickName"`
	Description *string        `json:"description"`
	Location    *string        `json:"location"`
	APIkey      *string        `json:"APIkey"`
	Profiles    []*UserProfile `json:"profiles"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   *time.Time     `json:"updatedAt"`
}

type Users struct {
	Count *int    `json:"count"`
	List  []*User `json:"list"`
}

type UserStatus string

const (
	UserStatusActive    UserStatus = "Active"
	UserStatusInActive  UserStatus = "InActive"
	UserStatusSuspended UserStatus = "Suspended"
)

var AllUserStatus = []UserStatus{
	UserStatusActive,
	UserStatusInActive,
	UserStatusSuspended,
}

func (e UserStatus) IsValid() bool {
	switch e {
	case UserStatusActive, UserStatusInActive, UserStatusSuspended:
		return true
	}
	return false
}

func (e UserStatus) String() string {
	return string(e)
}

func (e *UserStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = UserStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid UserStatus", str)
	}
	return nil
}

func (e UserStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
