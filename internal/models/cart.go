package models

// User Account  Carts
type Cart struct {
	BaseModelSoftDelete
	UserID string
	User   User
	Items  []*CartItem
}

type CartItem struct {
	BaseModelSeq
	ProductID string
	Product   ProductVariant
	Quantity  uint32
}
