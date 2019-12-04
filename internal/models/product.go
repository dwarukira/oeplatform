package models

// Product reps the oe product
type Product struct {
	BaseModelSoftDelete // We don't to actually delete the sellers, maybe audit
	Name                string
	Slug                string
	Seller              Seller
	SellerID            string
	Description         string
	Extra               string
	Code                string
	IsPublished         bool
	PublishedScope      string
	Active              bool
	Variants            []ProductVariant
}

// TableName returns the database table name for the ProductVariant model.
// func (Product) TableName() string {
// 	return tableName("products")
// }

// ProductVariant reps the oe productVariant
type ProductVariant struct {
	BaseModelSoftDelete        // We don't to actually delete the product, maybe audit
	Name                string `gorm:"not null;index"`
	Sku                 string
	Taxable             bool
	Barcode             string
	InventoryQuantity   int
	Weight              float64
	WeightUnit          string
	RequiresShipping    bool
	Grams               float64
	CompareAtPrice      float64 `gorm:"money"`
	Price               float64 `gorm:"money"`
	TrackInventory      bool
	QuantityAllocated   int
	Product             Product
	ProductID           string  `gorm:"not null;index"`
	Images              []Image `json:"images"`
}

type Image struct {
	BaseModelSoftDelete
	Name             string
	Source           string
	ProductVariant   ProductVariant
	ProductVariantID string `gorm:"not null; index"`
}

// TableName returns the database table name for the ProductVariant model.
// func (ProductVariant) TableName() string {
// 	return tableName("product_variants")
// }

func (o *ORM) GetProduct(id string) (*Product, error) {
	db := o.DB.New()
	up := &Product{}
	if err := db.Where("id = ?", id).First(up).Error; err != nil {
		return nil, err
	}

	return up, nil
}

func (o *ORM) GetProductVariant(id string) (*ProductVariant, error) {
	db := o.DB.New()
	up := &ProductVariant{}

	if err := db.Where("id = ?", id).First(up).Error; err != nil {
		return nil, err
	}

	return up, nil
}
