package models

// PendingState is the pending state of an Order
const PendingState = "pending"

// PaidState is the paid state of an Order
const PaidState = "paid"

// ShippingState is the shipping state of an order
const ShippingState = "shipping"

// ShippedState is the shipped state of an Order
const ShippedState = "shipped"

// FailedState is the failed state of an Order
const FailedState = "failed"

// PaymentState are the possible values for the PaymentState field
var PaymentStates = []string{
	PendingState,
	PaidState,
	FailedState,
}

// FulfillmentStates are the possible values for the FulfillmentState field
var FulfillmentStates = []string{
	PendingState,
	ShippingState,
	ShippedState,
}

// Order model
type Order struct {
	BaseModelSoftDelete
	IP                string `json:"ip"`
	User              *User
	UserID            string
	Email             string
	ShippingAddress   Address `gorm:"ForeignKey:ShippingAddressID"`
	ShippingAddressID string
	BillingAddress    Address
	BillingAddressID  string `gorm:"ForeignKey:BillingAddressID"`
}

type OrderItem struct {
	BaseModelSoftDelete
	OrderID  string "json:\"-\""
	Quantity uint64
}
