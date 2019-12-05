package models

import (
	"encoding/json"
	"time"
)

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
	VATNumber         string
	Items             []*OrderItem

	Taxes    uint64
	Currency string
	SubTotal uint64 `json:"subtotal"`
	Discount uint64 `json:"discount"`
	NetTotal uint64 `json:"net_total"`

	Total uint64 `json:"total"`

	PaymentState     string `json:"payment_state"`
	FulfillmentState string `json:"fulfillment_state"`
	State            string `json:"state"`

	MetaData    map[string]interface{} `sql:"-" json:"meta"`
	RawMetaData string                 `json:"-" sql:"type:text"`
	CouponCode  string                 `json:"coupon_code,omitempty"`

	Coupon    *Coupon `json:"coupon,omitempty" sql:"-"`
	RawCoupon string  `json:"-" sql:"type:text"`
}

// FixedAmount represents an amount and currency pair
type FixedAmount struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

// Coupon represents a discount redeemable with a code.
type Coupon struct {
	Code string `json:"code"`

	StartDate *time.Time `json:"start_date,omitempty"`
	EndDate   *time.Time `json:"end_date,omitempty"`

	Percentage  uint64         `json:"percentage,omitempty"`
	FixedAmount []*FixedAmount `json:"fixed,omitempty"`

	ProductTypes []string               `json:"product_types,omitempty"`
	Products     []string               `json:"products,omitempty"`
	Claims       map[string]interface{} `json:"claims,omitempty"`
}

// AfterFind database callback.
func (o *Order) AfterFind() error {
	if o.RawMetaData != "" {
		err := json.Unmarshal([]byte(o.RawMetaData), &o.MetaData)
		if err != nil {
			return err

		}

	}
	if o.RawCoupon != "" {
		o.Coupon = &Coupon{}
		err := json.Unmarshal([]byte(o.RawCoupon), &o.Coupon)
		if err != nil {
			return err

		}

	}

	return nil

}

// BeforeSave database callback.
func (o *Order) BeforeSave() error {
	if o.MetaData != nil {
		data, err := json.Marshal(o.MetaData)
		if err != nil {
			return err

		}
		o.RawMetaData = string(data)

	}
	if o.Coupon != nil {
		data, err := json.Marshal(o.Coupon)
		if err != nil {
			return err

		}
		o.RawCoupon = string(data)

	}

	return nil

}

func (o *ORM) CreateNewPeddingOrder() {

}

type OrderItem struct {
	BaseModelSoftDelete
	OrderID  string "json:\"-\""
	Quantity uint64

	ProductID       string
	Product         Product
	ProductVarint   ProductVariant
	ProductVarintID string
	Description     string `json:"description" sql:"type:text"`

	MetaData    map[string]interface{} `sql:"-" json:"meta"`
	RawMetaData string                 `json:"-" sql:"type:text"`
}
